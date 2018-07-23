package api

import (
	"testing"

	"github.com/itouri/sgx-iaas/pkg/domain/heat"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

var testYaml1 = `
auto-scaling-groups:
- name: tmp_asg
  min-size: 1
  max-size: 3
  properties:
    flavor: tiny
    image: arch
scaling-policies:
- name: tmp_sp
  auto-scaling-group-name: tmp_asg
  cooldown: 1
  scaling-adjustment: 1
alarms:
- name: cpu_vm_scaleup_policy
  meter-name: cpu
  threshold: 1
  comparison-operator: gt
  alarm-action: tmp_sp
- name: mem_vm_scaleup_policy
  meter-name: mem
  threshold: 2
  comparison-operator: gt
  alarm-action: tmp_sp
`

func TestParseYaml(t *testing.T) {
	assert := assert.New(t)

	testSrct := &heat.Template{
		AutoScalingGroups: []heat.AutoScalingGroup{
			heat.AutoScalingGroup{
				Name:    "tmp_asg",
				MinSize: 1,
				MaxSize: 3,
				Properties: heat.ASGProperty{
					Flavor: "tiny",
					Image:  "arch",
				},
			},
		},
		ScalingPolicies: []heat.ScalingPolicy{
			heat.ScalingPolicy{
				Name:                 "tmp_sp",
				AutoScalingGroupName: "tmp_asg",
				Cooldown:             1,
				ScalingAdjustment:    1,
			},
		},
		Alarms: []heat.Alarm{
			heat.Alarm{
				Name:               "cpu_vm_scaleup_policy",
				MeterName:          "cpu",
				Threshold:          1,
				ComparisonOperator: "gt",
				AlarmActions:       "tmp_sp",
			},
			heat.Alarm{
				Name:               "mem_vm_scaleup_policy",
				MeterName:          "mem",
				Threshold:          2,
				ComparisonOperator: "gt",
				AlarmActions:       "tmp_sp",
			},
		},
	}

	template := &heat.Template{}
	err := yaml.Unmarshal([]byte(testYaml1), template)
	assert.NoError(err)

	assert.Equal(testSrct, template)
}
