package heat

type ASGProperty struct {
	Flavor string `yaml:"flavor"`
	Image  string `yaml:"image"`
}

type AutoScalingGroup struct {
	Name       string      `yaml:"name"`
	MinSize    int         `yaml:"min-size"`
	MaxSize    int         `yaml:"max-size"`
	Properties ASGProperty `yaml:"properties"`
}

type ScalingPolicy struct {
	Name                 string `yaml:"name"`
	AutoScalingGroupName string `yaml:"auto-scaling-group-name"`
	Cooldown             int    `yaml:"cooldown"`
	ScalingAdjustment    int    `yaml:"scaling-adjustment"`
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

type Alarm struct {
	Name               string  `yaml:"name"`
	MeterName          string  `yaml:"meter-name"` // string?
	Threshold          float32 `yaml:"threshold"`
	AlarmActions       string  `yaml:"alarm-actions"`
	ComparisonOperator string  `yaml:"comparison-operator"`
}

type Template struct {
	AutoScalingGroups []AutoScalingGroup `yaml:"auto-scaling-groups"`
	ScalingPolicies   []ScalingPolicy    `yaml:"scaling-policies"`
	Alarms            []Alarm            `yaml:"alarms"`
}
