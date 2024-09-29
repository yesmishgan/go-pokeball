package swaggway

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type (
	swaggerDesc interface {
		Swagger() []byte
	}
)

// NewMux creates http mux for grpc-gateway
func NewMux(sd swaggerDesc, cr chi.Router, opts ...runtime.ServeMuxOption) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(append([]runtime.ServeMuxOption{
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			// grpc-gateway writes http headers with 'grpcgateway-' prefix by default,
			// (eg "grpcgateway-accept":"application/json", "grpcgateway-user-agent":"curl/7.64.1", etc)
			// this option used for BC with clay. Clay does not modify headers

			// grpc-gateway already pass through `Authorization` (not `authorization`) header in runtime.AnnotateIncomingContext
			if s == "Authorization" {
				// If return true header will be doubled
				return s, false
			}

			return s, true
		})}, opts...)...)

	if err := mountHandlersFromSwagger(sd, cr, mux); err != nil {
		return nil, fmt.Errorf("mountHandlersFromSwagger: %w", err)
	}

	return mux, nil
}

func mountHandlersFromSwagger(desc swaggerDesc, chiMux chi.Router, h *runtime.ServeMux) error {
	doc := &spec.Swagger{}
	if err := doc.UnmarshalJSON(desc.Swagger()); err != nil {
		return fmt.Errorf("doc.UnmarshalJSON: %w", err)
	}

	swag := analysis.New(doc)
	for k, v := range swag.AllPaths() {
		if v.Get != nil {
			chiMux.Method(http.MethodGet, k, h)
		}
		if v.Post != nil {
			chiMux.Method(http.MethodPost, k, h)
		}
		if v.Put != nil {
			chiMux.Method(http.MethodPut, k, h)
		}
		if v.Delete != nil {
			chiMux.Method(http.MethodDelete, k, h)
		}
		if v.Patch != nil {
			chiMux.Method(http.MethodPatch, k, h)
		}
		if v.Options != nil {
			chiMux.Method(http.MethodOptions, k, h)
		}
		if v.Head != nil {
			chiMux.Method(http.MethodHead, k, h)
		}
	}
	return nil
}
