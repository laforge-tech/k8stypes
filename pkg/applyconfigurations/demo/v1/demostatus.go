//-----------------------------------------------------------------------------
// Demo API
//-----------------------------------------------------------------------------

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DemoStatusApplyConfiguration represents an declarative configuration of the DemoStatus type for use
// with apply.
type DemoStatusApplyConfiguration struct {
	Phase *string `json:"phase,omitempty"`
}

// DemoStatusApplyConfiguration constructs an declarative configuration of the DemoStatus type for use with
// apply.
func DemoStatus() *DemoStatusApplyConfiguration {
	return &DemoStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *DemoStatusApplyConfiguration) WithPhase(value string) *DemoStatusApplyConfiguration {
	b.Phase = &value
	return b
}
