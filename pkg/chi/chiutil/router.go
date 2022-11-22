package chiutil

import "go-archive/pkg/chi"

type Resource interface {
	Register(r chi.Router)
}

func SetupResource(mux *chi.Mux, pattern string, resources ...Resource) {
	for _, resource := range resources {
		mux.Route(pattern, resource.Register)
	}
}
