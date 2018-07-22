package engine

import "github.com/itouri/sgx-iaas/pkg/domain/heat"

var stacks []heat.Template

func init() {
	stacks = []heat.Template{}
}

func RegisterStack(stack *heat.Template) {
	stacks = append(stacks, *stack)
}
