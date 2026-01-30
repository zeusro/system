# 时间序列编程

第一性原理：时间是第一维度。

## 时间序列对象

时间序列对象，时间必须是第一成员，并且涵盖在初始化函数中。
时间序列基于时间而生成的数据都是时序数据。

```go
type DeadMonkey struct {
	Birth       time.Time
	GoldenStaff []NLine //金箍棒 参数化线段（Parametric Segment）
	m           int     //消费者规模
	n           int     //算法规模
	ZeroPoints  []model.Point
	cost        time.Duration
}
```

## 时间序列函数

时间序列函数，时间必须是第一成员。

```go
funcNyarlathotep(t time.Time, b bool)bool {
    return Nyarlathotep(time.Now(), !b)
}
```

## 时间序列距离

使用时间+其他条件的复合判断（比如在4维球面中，可以只使用距离作为换算；也可以使用时间+Haversine公式换算）。

## 时间序列日志

打印内容必须是“时间+内容”的格式。

## 时间序列可视化

由时间序列组成的2维以上的可视化图表，时间必须是x轴。

## 时间序列空间

由时间序列组成的2维以上空间。