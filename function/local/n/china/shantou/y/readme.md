# Repeated-Game Theoretic Fixes in an Incentive-Based Educational Simulation

## Abstract

We identify and fix several bugs in a multi-agent, time-series simulation of an administrative education system (teachers, students, psychologist, school leader) where performance is driven by an incentive function over average score and college enrollment rate. The original implementation treated each period as a one-shot game, leading to (1) order-dependent outcomes within a period, (2) no history-dependent strategies, and (3) incorrect “best strategy” reporting when no student group had enough members. We apply repeated-game theory—horizon and history dependence, simultaneous move within a period, and reputation—to correct these issues and document the changes in this note.

**Keywords:** repeated games, incentive design, educational administration, multi-agent simulation, time-series model.

---

## 1. Introduction

The simulation in this package models a school cohort over a time window: teachers (Y and F), named and generic students, a psychologist, and a school leader. Each day, agents choose strategies (e.g., normal teaching, PUA, reduce exam count, study hard, avoid, dropout, network violence). Outcomes update scores, stress, exam pool, and dropout state. The leader’s “performance” (incentive) is a function of average score and college enrollment rate. The original code had three main problems:

1. **Within-period order dependency:** Strategies were chosen and applied in a single loop over agents. The first agent’s applied consequence could change the state before the next agent chose, so the outcome depended on agent order—inconsistent with a simultaneous-move interpretation of the stage game.
2. **One-shot stage game:** Strategies depended only on the current aggregate state. There was no use of past actions (e.g., “did the teacher defect last period?”) or of the remaining horizon (e.g., end-game effects), which are central in repeated games.
3. **Reporting bug:** The “best strategy by score / by stay rate” logic only updated when a group had at least three members. When every group had fewer than three, the reported best score and best stay rate stayed at −1, producing invalid summary text.

We fix (1) by making each period **two-phase**: first all in-school agents choose strategies given the same context (and history), then all consequences are applied. We fix (2) by adding **repeated-game state**: remaining steps, last-round “teacher defection” (PUA or reduce exam count), and last strategy per agent; and by making strategy choice **history- and horizon-dependent**. We fix (3) by **fallback** when no group has at least three members: we still choose the best strategy by average score and by stay rate over all non-empty groups so the reported values are always valid.

---

## 2. Repeated-Game Framework

### 2.1 Stage game and repetition

- **Stage game:** Each day is one period. Agents are the players; each in-school agent chooses an action (strategy) from a finite set. Payoffs are implicit in the consequences (score, stress, exam pool, dropout) and in the incentive function for the leader.
- **Repeated game:** The stage game is played for a finite number of periods (from simulation start to end). We assume agents can condition on (i) **remaining horizon** (number of periods left) and (ii) **last period’s “defection”** by teachers (use of PUA or reduce-exam-count).

### 2.2 Cooperation and defection (operational)

- **Teachers:** “Cooperate” = normal teaching; “Defect” = PUA (teacher Y) or reduce exam count (teacher F).
- **Students:** “Cooperate” = study hard or avoid; “Defect” = e.g. network violence (Judas), dropout (P / generic). We do not formalize payoffs here; we only use the idea that **reputation** and **horizon** matter: after teacher defection, students may “punish” (e.g., avoid, or Judas may escalate); teachers may “restore reputation” by cooperating the next period.

### 2.3 Implementation of repeated-game state

- **SimContext** is extended with:
  - `StepsRemaining`: number of periods left (including current). Used so strategies can depend on horizon (e.g., end-game).
  - `LastRoundTeacherDefection`: true if in the previous period at least one teacher used PUA or reduce exam count.
- **Agent** is extended with:
  - `LastStrategy`: the strategy chosen in the previous period (for future use; currently strategy rules use only `LastRoundTeacherDefection` and `StepsRemaining`).
- **Run loop:** Before each period we set `ctx.StepsRemaining = steps - i` and `ctx.LastRoundTeacherDefection = lastRoundTeacherDefection`. After applying all consequences we set `lastRoundTeacherDefection` from this period’s teacher actions.

---

## 3. Two-Phase Step (Simultaneous Move Within Period)

Previously, in one loop over agents we did: choose strategy → apply consequence (possibly changing other agents). Thus later agents could see state changes from earlier agents in the same period.

**Change:** Each period is split into two phases.

1. **Phase 1 – Choice:** For each in-school agent, compute strategy from `ChooseStrategy(now, agent, &ctx)` and store it in a per-agent array. No state is updated in this phase; everyone sees the same `ctx` (and same history).
2. **Phase 2 – Apply:** For each agent, apply the stored strategy (update scores, stress, exam pool, dropout, events) and set `agent.LastStrategy`. If any teacher used PUA or reduce exam count, set `lastRoundTeacherDefection = true` for the next period.

This removes within-period order dependency and aligns the stage game with a simultaneous-move interpretation.

---

## 4. History- and Horizon-Dependent Strategy Rules

### 4.1 Teachers

- **Teacher Y (average-score incentive):**  
  - If legal/moral risk is high → lie/evade.  
  - **Repeated game:** If `LastRoundTeacherDefection` and `StepsRemaining > 1` → normal teach (restore reputation).  
  - Else if average score is low and student count is large → PUA.

- **Teacher F (score + enrollment rate):**  
  - If legal/moral risk is high → lie/evade.  
  - **Repeated game:** If `LastRoundTeacherDefection` and `StepsRemaining > 1` → normal teach.  
  - Else if exam count is high and enrollment rate is low → reduce exam count.

So after “defection,” teachers cooperate for at least one period when there is future play; they can defect again when conditions re-trigger.

### 4.2 Students

- **Judas:** If student count is high and own legal risk is low: if `StepsRemaining <= 1` (end game) or `LastRoundTeacherDefection` (retaliate) → network violence; else study hard.
- **Generic student:** If net PUA and stress are high → dropout. **Repeated game:** If `LastRoundTeacherDefection` and `StepsRemaining > 2` → avoid (reduce exposure). Else by IQ: high IQ → study hard, else avoid.

Other student roles (P, Y, C13) and psychologist/leader are unchanged except that they now operate in the two-phase, same-context setup.

---

## 5. Reporting Bug Fix (Best Strategy When No Group Has ≥3 Members)

The summary “best strategy by average score” and “best strategy by stay rate” only updated when a group had at least three members. So when all strategy groups had size &lt; 3, `bestScore` and `bestStayRate` remained −1 and the printed “best” values were wrong.

**Fix:** After the main loop over strategy groups:

- If `bestScore < 0`, run a fallback loop over all non-empty groups and set best-by-score to the group with the highest average score.
- If `bestStayRate < 0`, run a fallback loop and set best-by-stay-rate to the group with the highest stay rate.

So we always report a valid best strategy and numeric value, with a preference for groups of size ≥ 3 when available.

---

## 6. Summary of Code and Behavioral Changes

| Item | Change |
|------|--------|
| **SimContext** | Added `StepsRemaining`, `LastRoundTeacherDefection`. |
| **Agent** | Added `LastStrategy`. |
| **Run** | Two-phase period: (1) choose all strategies, (2) apply all and update last-round defection. |
| **Teacher Y / F** | After defection and if steps remain, cooperate one period; otherwise keep original condition-based defection. |
| **Judas** | Use network violence in end game or after teacher defection (when conditions hold). |
| **Generic student** | After teacher defection and enough steps left, prefer avoid. |
| **printStudentBestStrategy** | Fallback to best over all non-empty groups when no group has count ≥ 3. |

---

## 7. Conclusion

Applying repeated-game ideas (simultaneous move within the stage game, horizon, and one-period “defection” history) fixes order dependency, adds history- and horizon-dependent behavior, and corrects the best-strategy reporting bug. The simulation remains a discrete-time, incentive-driven multi-agent model; the new parameters and two-phase step are minimal extensions that align behavior with standard repeated-game reasoning and more robust output.

---

## References

1. Fudenberg, D., & Tirole, J. (1991). *Game Theory*. MIT Press.  
2. Mailath, G. J., & Samuelson, L. (2006). *Repeated Games and Reputations: Long-Run Relationships*. Oxford University Press.  
3. Package design doc: `y.md` (time-series model, roles, strategies, incentive function).
