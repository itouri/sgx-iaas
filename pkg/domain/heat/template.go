package heat

import (
	"github.com/google/uuid"
)

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
	AlarmID              uuid.UUID
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

// TODO bsonとyamlは同じなんだけど２回かかなきゃダメ？
type Alarm struct {
	ID                 uuid.UUID `bson:"alarm_id"`
	Name               string    `bson:"name" yaml:"name"`
	MeterName          string    `bson:"meter_name" yaml:"meter-name"` // string?
	Threshold          float32   `bson:"threshold" yaml:"threshold"`
	AlarmAction        string    `bson:"alarm_action" yaml:"alarm-action"`
	ComparisonOperator string    `bson:"comparison_operator" yaml:"comparison-operator"`
}

type Template struct {
	AutoScalingGroups []AutoScalingGroup `bson:"auto_scaling_groups" yaml:"auto-scaling-groups"`
	ScalingPolicies   []ScalingPolicy    `bson:"scaling_policies" yaml:"scaling-policies"`
	Alarms            []Alarm            `bson:"alarms" yaml:"alarms"`
}
