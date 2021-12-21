package interceptor

import (
	"log"
)

// Registry is a collector for interceptors.
type Registry struct {
	factories []Factory
}

// Add adds a new Interceptor to the registry.
func (r *Registry) Add(f Factory) {
	r.factories = append(r.factories, f)
}

// Build constructs a single Interceptor from a InterceptorRegistry
func (r *Registry) Build(id string) (Interceptor, error) {

	log.Printf("#-> Registry::Build()\n")

	if len(r.factories) == 0 {
		return &NoOp{}, nil
	}

	interceptors := []Interceptor{}
	for _, f := range r.factories {
		i, err := f.NewInterceptor(id)
		if err != nil {
			return nil, err
		}

		interceptors = append(interceptors, i)
	}
	log.Printf("<-# Registry::Build()\n")
	return NewChain(interceptors), nil
}
