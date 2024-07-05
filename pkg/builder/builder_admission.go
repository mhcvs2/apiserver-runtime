package builder

import (
	"io"

	"k8s.io/apiserver/pkg/admission"

	"sigs.k8s.io/apiserver-runtime/internal/sample-apiserver/pkg/cmd/server"
)

// DisableAdmissionControllers disables delegated authentication and authorization
func (a *Server) DisableAdmissionControllers() *Server {
	server.ServerOptionsFns = append(server.ServerOptionsFns, func(o *ServerOptions) *ServerOptions {
		o.RecommendedOptions.Admission = nil
		return o
	})
	return a
}

func (a *Server) WithAdmissionPlugin(name string, plugin admission.Interface) *Server {
	server.ServerOptionsFns = append(server.ServerOptionsFns, func(o *ServerOptions) *ServerOptions {
		if o.RecommendedOptions.Admission != nil {
			o.RecommendedOptions.Admission.Plugins.Register(name, func(config io.Reader) (admission.Interface, error) {
				return plugin, nil
			})
			o.RecommendedOptions.Admission.RecommendedPluginOrder = append(o.RecommendedOptions.Admission.RecommendedPluginOrder, name)
		}
		return o
	})
	return a
}
