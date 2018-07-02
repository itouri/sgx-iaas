package heat

type ASGPropaty struct {
	Flavor string
	Image  string
}

type AutoScalingGroup struct {
	MinSize int
	MaxSize int
	Propaty ASGPropaty
}

type ScalingPolicy struct {
	AutoScalingGroup
	Cooldown          int
	ScalingAdjustment int
}

type EnumComparisonOperator int

const (
	Ge EnumComparisonOperator = iota + 1 // =>
	Le                                   // =<
	Gt                                   // >
	Lt                                   // <
	Eq                                   // =
	Ne                                   // !=
)

const ()

type Alarm struct {
	MeterName string // string?
	// Statistic string
	Threshold          float32
	ComparisonOperator EnumComparisonOperator
}

type Template struct {
	AutoScalingGroups []AutoScalingGroup
	ScalingPoliceis   []ScalingPolicy
	Alarms            []Alarm
}
