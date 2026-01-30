# Multi-Agent Game and Time-Incentive Function Under Educational Bureaucracy: A Formal Study Based on Repeated Games and Time-Series Simulation

## Abstract

Under the institutional setting of educational bureaucracy, we model the school organization as a multi-agent game: teachers, students, a psychologist, and school leadership act as rational agents whose strategies are driven by quantified performance metrics—average student score and undergraduate enrollment rate. We introduce a time-series meta-programming model with time as the first dimension, define quantitative factors (family background, intelligence, emotional intelligence, PUA exposure/resistance, legal-moral risk), establish role–strategy–consequence mappings, and specify the leadership time-incentive function \(I(t) = 0.6 \cdot \bar{s}(t) + 0.4 \cdot r(t)\). Through a repeated-game extension (history- and horizon-dependent strategies), we implement a time-series simulation in Go. Results show that teachers choose “lie-and-evade” when legal risk is high, and PUA or reduce-exam-count when performance is low; students switch among study-hard, avoid, and dropout depending on net PUA pressure and whether teachers defected in the previous round; leadership applies downward pressure (kick-the-cat effect) when \(I(t)\) falls below a threshold. The simulation outputs time-series logs and incentive samples, providing a computable reference for policy risk and incentive design.

**Keywords:** Educational bureaucracy; multi-agent game; repeated game; time-incentive function; kick-the-cat effect; time-series simulation

---

## 1. Introduction

### 1.1 Background and Motivation

In educational bureaucracy, school performance is often quantified by observable metrics such as average student score and undergraduate enrollment rate. This incentive structure turns human decisions into computable variables and thereby influences teachers’ and students’ strategy choices, potentially inducing behaviors such as PUA-driven attrition, reduction of exam-taker counts, dropout, and cyberbullying—forming a typical **multi-agent, multi-period, incomplete-information** game. To analyze policy risks and equilibrium outcomes of different decisions and to design a reasonable time-incentive function for school leadership, a formal game-theoretic model and an executable simulation are needed.

### 1.2 Related Concepts

- **Educational bureaucracy:** Students are accountable to teachers, teachers to principals, principals to the school; performance is measured by average student score and undergraduate enrollment rate.
- **Kick-the-cat effect:** Higher-tier agents displace frustration downward along the hierarchy; in the model this appears as leadership pressure when performance is poor and teacher pressure on students.

### 1.3 Contributions and Structure

Main contributions of this paper: (1) formalizing the educational bureaucracy setting as a multi-role, multi-strategy extended/repeated game; (2) proposing a “time as first dimension” time-series meta-programming convention that unifies the time semantics of objects and functions; (3) giving an explicit time-incentive function for leadership and a method for simulation sampling; (4) implementing a full time-series simulation in Go with outputs for strategy choices, consequences, and incentive curves. Structure: §2 formal model, §3 game structure (roles, strategies, payoffs), §4 equilibrium and strategy logic, §5 simulation design and implementation, §6 results and discussion, §7 conclusion.

---

## 2. Formal Model

### 2.1 Time-Series Meta-Programming Convention

- **First principle:** Time is the first dimension. Time-series objects have time as the first member; time-series functions have time as the first parameter.
- **Time-series objects:** `Factor`, `Agent`, `SimState` all have `Birth` or `Current` as the first member.
- **Time-series functions:** \(t \mapsto \text{Incentive}(t,\ldots)\), \(\text{ChooseStrategy}(t,\ldots)\), \(\text{ApplyStrategy}(t,\ldots)\) all take time \(t\) as the first parameter.
- **Time-series logs:** Events are recorded as “time + content” (\(\text{LogTS}\)).

### 2.2 Quantitative Factors

Each agent carries a factor vector \(\mathbf{f}\) with components in \([0,1]\) (or standardized):

| Factor | Symbol | Meaning |
|--------|--------|---------|
| Family background | \(\text{FamilyBackground}\) | 0 = very poor, 1 = very rich; affects resources and enrollment paths |
| Intelligence | \(\text{IQ}\) | Affects academic performance and strategy comprehension |
| Emotional intelligence | \(\text{EQ}\) | Affects stress resistance and emotion propagation in the kick-the-cat chain |
| PUA exposure | \(\text{PUAExposure}\) | Intensity of teacher PUA strategy on this individual |
| PUA resistance | \(\text{PUAResistance}\) | Individual resistance to PUA |
| Legal-moral risk | \(\text{LegalMoralRisk}\) | Risk of legal/moral accountability for the individual or behavior |

**Net PUA pressure:** \(\pi = \text{PUAExposure} \times (1 - \text{PUAResistance})\), used to drive dropout/avoid strategies.

### 2.3 Performance and Incentive Formulas

- **Formula 1:** Average student score \(\bar{s} = \frac{\sum_i s_i}{N}\) (\(N\) = in-school student count).
- **Formula 2:** Undergraduate enrollment rate \(r = \frac{\text{enrolled}}{\text{exam-takers}} \times 100\%\).

**Time-incentive function** (leadership-perceived performance):

\[
I(t) = 0.6 \cdot \bar{s}(t) + 0.4 \cdot r(t)
\]

where \(\bar{s}(t)\) and \(r(t)\) are computed from in-school count, exam-taker count, and enrollment count at time \(t\).

---

## 3. Game Structure: Roles, Strategies, and Payoffs

### 3.1 Roles (Players) and Types

| Role | Incentive basis | Type / traits |
|------|-----------------|----------------|
| Teacher Y | Formula 1 (avg score) | Can reduce denominator (PUA attrition), evade scrutiny (lie-and-evade) |
| Teacher F | Formula 1 + 2 | Can reduce exam-taker count, lie-and-evade |
| Student Judas | Aligns with Teacher F | Affluent middle-class, can use cyberbullying to reduce numbers |
| Student Black Mamba | Wealth | Does not sit exam; strategy set empty |
| Student P | Low IQ, high PUA exposure | Dropout as passive resistance |
| Student Y | Athlete bonus | High-IQ athlete; athlete bonus requires school leader approval |
| Student C13 | High IQ, poor | Stable choice: study hard |
| Psychologist | System balance | Decompress and soothe |
| School leader | \(I(t)\) | Allocates resources (e.g. who gets bonus, who drops out); low \(I(t)\) → pressure down; high → design incentives |

### 3.2 Strategy Space

- **Teachers:** PUA pressure attrition, reduce exam count, lie-and-evade, normal teaching.
- **Students:** Dropout, cyberbullying, athlete bonus, study hard, avoid.
- **Psychologist:** Decompress/soothe, or no action.
- **Leader:** Pressure down (kick-the-cat), design incentive function.

### 3.3 Consequence (Payoff) Structure

Strategy execution yields consequences \(\omega\): \(\Delta\text{Stress}\), \(\Delta\text{Score}\), \(\text{LegalRisk}\), \(\text{Dropout}\), \(\text{LeaveExam}\), etc. For example: PUA pressure → target student stress up, actor legal risk up; dropout → agent leaves school and exam pool; decompress → target student stress down. Payoffs enter leadership utility indirectly via \(\bar{s}(t)\), \(r(t)\), and \(I(t)\), and affect next-period state and other agents’ strategies through \(\text{Stress}\), \(\text{InSchool}\), \(\text{InExamPool}\).

---

## 4. Equilibrium and Strategy Logic (Repeated-Game Extension)

### 4.1 Information and History

- **Repeated game:** Multi-period steps; each period first updates context using current state, **steps remaining**, and **last-round teacher defection** (PUA / reduce exam count), then all agents choose strategies simultaneously, then consequences are applied uniformly.
- **History dependence:** If teachers defected last round and steps remaining \(> 1\), teachers may choose “normal teach” to restore reputation; students may choose “cyberbullying” or “avoid” when teachers defected.

### 4.2 Strategy-Selection Logic (Equilibrium Correspondence)

- **Teacher Y:** If \(\text{LegalMoralRisk} > 0.6\) → lie-and-evade; if last-round defection and steps remaining \(> 1\) → normal teach; if \(\bar{s} < 0.5\) and \(N > 5\) → PUA; else normal teach.
- **Teacher F:** If \(\text{LegalMoralRisk} > 0.5\) → lie-and-evade; if last-round defection and steps remaining \(> 1\) → normal teach; if exam count \(> 3\) and enrollment rate \(< 0.6\) → reduce exam count; else normal teach.
- **Student Judas:** If \(N > 4\), \(\text{LegalMoralRisk} < 0.5\), and (last period or last-round teacher defection) → cyberbullying; else study hard.
- **Student P:** If net PUA pressure \(\pi > 0.5\) and \(\text{Stress} > 0.6\) → dropout; else avoid.
- **Generic student:** High \(\pi\) and high stress → dropout; last-round teacher defection and steps remaining \(> 2\) → avoid; else choose study hard / avoid by IQ.
- **Psychologist:** If average stress \(> 0.4\) → decompress.
- **Leader:** Allocates resources (e.g. approves athlete bonus when choosing “design incentive”). If \(I(t) < 0.5\) → pressure down; else design incentive. Student Y’s athlete bonus takes effect only when leader chooses design incentive in that step.

These rules match the implementation \(\text{ChooseStrategy}(t, a, \text{ctx})\) and can be interpreted as **behavioral strategy** equilibrium correspondences under discrete states and thresholds.

---

## 5. Simulation Design and Implementation

### 5.1 Step and Phases

- **Step size:** Daily (or configurable); number of steps from time interval \([t_0, t_1]\) and step size.
- **Per step:** (1) Update \(\text{SimContext}\) (student count, exam count, enrollment rate, average stress) from current roster; (2) call \(\text{ChooseStrategy}(t, a, \text{ctx})\) for each active agent; (3) call \(\text{ApplyStrategy}(t, s, a, \text{ctx}, \text{rng})\) and apply consequences to actor or random target; (4) append \(\text{LogTS}\); (5) sample \(I(t)\) to get \((t, I(t))\) points.

### 5.2 Implementation Notes

- **Code layout:** `model.go` (time-series objects), `roles.go` (role/strategy enums, Agent construction), `incentive.go` (\(I(t)\)), `strategy.go\) (\(\text{ChooseStrategy}\), \(\text{ApplyStrategy}\), consequence quantification), `sim.go\) (\(\text{SimContext}\), \(\text{Run}\), \(\text{LogTS}\)), main entry builds agents and calls \(\text{Run}\).
- **Outputs:** Time-series logs (time + content), incentive samples (time → performance), final statistics (in-school count, exam count, enrollment count, average score, enrollment rate, \(I(t)\)), and student grouping by dominant strategy with best-strategy recommendations.

---

## 6. Results and Discussion

### 6.1 Incentive Curve and Performance Evolution

The sampled sequence of \(I(t)\) over \(t\) gives the “time–performance” curve. Performance is driven by both average score and enrollment rate; PUA attrition or reducing exam count may raise \(\bar{s}\) or \(r\) in the short run but trigger dropout, exam exit, and legal risk, and in the long run may reduce roster and exam-taker counts—analysis depends on parameters and random seed.

### 6.2 Strategy Grouping and Best Strategies

Simulation groups students by dominant strategy (study hard, avoid, dropout, athlete bonus, cyberbullying). Typical finding: groups whose dominant strategy is “study hard” or “avoid” have better average score and retention; dropout group has zero retention and is triggered mainly under high PUA exposure and high stress; consistent with “cooperate–defect–retaliate” patterns in repeated games.

### 6.3 Policy Implications and Limitations

- **Policy:** Incentive weights (0.6/0.4) and thresholds (e.g. \(I(t) < 0.5\) triggers pressure down) directly affect equilibrium behavior; raising legal-moral risk or lowering PUA exposure can curb attrition and cyberbullying; psychologist decompression helps lower system average stress.
- **Limitations:** Model is a simplified multi-agent, discrete-strategy, deterministic-threshold setting; no Bayesian treatment of incomplete information; calibration and robustness with real data left for future work.

---

## 7. Conclusion

We have built a multi-agent, multi-strategy, time-driven game under educational bureaucracy and applied a “time as first dimension” time-series meta-programming convention throughout object and function design. Via a repeated-game extension (history and horizon dependence) and a Go-based time-series simulation, we obtain strategy-selection logic for teachers, students, psychologist, and leadership and the evolution of the incentive function \(I(t)\). Simulation results support the relative advantage of “study hard / avoid” at the student level and equilibrium behavior of leadership and teachers under performance and legal-risk constraints. The work provides a formal and implementational basis for policy risk analysis and incentive design that is computable and reproducible.

---

## References

1. Kick-the-cat effect. Wikipedia. https://en.wikipedia.org/wiki/Displaced_aggression  
2. Project documentation and implementation: `function/local/n/china/shantou/y/y.md`, `model.go`, `roles.go`, `incentive.go`, `strategy.go`, `sim.go`, main entry `y.go`.
