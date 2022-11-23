package chiutil

import "go-archive/pkg/chi"

type Resource interface {
	Register(r chi.Router)
}

func SetupResource(mux *chi.Mux, pattern string, resources ...Resource) {
	mux.Route(pattern, func(r chi.Router) {
		for _, resource := range resources {
			r.Group(resource.Register)
		}
	})
}
