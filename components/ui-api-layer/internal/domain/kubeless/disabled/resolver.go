// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import context "context"
import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"

// Resolver is an autogenerated failing mock type for the Resolver type
type Resolver struct {
	err error
}

// NewResolver creates a new Resolver type instance
func NewResolver(err error) *Resolver {
	return &Resolver{err: err}
}

// FunctionsQuery provides a failing mock function with given fields: ctx, namespace, first, offset
func (_m *Resolver) FunctionsQuery(ctx context.Context, namespace string, first *int, offset *int) ([]gqlschema.Function, error) {
	var r0 []gqlschema.Function
	var r1 error
	r1 = _m.err

	return r0, r1
}
