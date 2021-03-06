// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/optikon/api/api/v0/server/restapi/operations/clusters"
	"github.com/optikon/api/api/v0/server/restapi/operations/releases"
)

// NewOptikonAPI creates a new Optikon instance
func NewOptikonAPI(spec *loads.Document) *OptikonAPI {
	return &OptikonAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		ClustersAddClusterHandler: clusters.AddClusterHandlerFunc(func(params clusters.AddClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersAddCluster has not yet been implemented")
		}),
		ReleasesAddReleasesHandler: releases.AddReleasesHandlerFunc(func(params releases.AddReleasesParams) middleware.Responder {
			return middleware.NotImplemented("operation ReleasesAddReleases has not yet been implemented")
		}),
		ClustersDeleteClusterHandler: clusters.DeleteClusterHandlerFunc(func(params clusters.DeleteClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersDeleteCluster has not yet been implemented")
		}),
		ReleasesDeleteReleaseHandler: releases.DeleteReleaseHandlerFunc(func(params releases.DeleteReleaseParams) middleware.Responder {
			return middleware.NotImplemented("operation ReleasesDeleteRelease has not yet been implemented")
		}),
		ClustersGetClusterByIDHandler: clusters.GetClusterByIDHandlerFunc(func(params clusters.GetClusterByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersGetClusterByID has not yet been implemented")
		}),
		ClustersGetClustersHandler: clusters.GetClustersHandlerFunc(func(params clusters.GetClustersParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersGetClusters has not yet been implemented")
		}),
		ReleasesGetReleaseByIDHandler: releases.GetReleaseByIDHandlerFunc(func(params releases.GetReleaseByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ReleasesGetReleaseByID has not yet been implemented")
		}),
		ReleasesGetReleasesHandler: releases.GetReleasesHandlerFunc(func(params releases.GetReleasesParams) middleware.Responder {
			return middleware.NotImplemented("operation ReleasesGetReleases has not yet been implemented")
		}),
		ClustersUpdateClusterHandler: clusters.UpdateClusterHandlerFunc(func(params clusters.UpdateClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersUpdateCluster has not yet been implemented")
		}),
		ReleasesUpdateReleaseHandler: releases.UpdateReleaseHandlerFunc(func(params releases.UpdateReleaseParams) middleware.Responder {
			return middleware.NotImplemented("operation ReleasesUpdateRelease has not yet been implemented")
		}),
	}
}

/*OptikonAPI Optikon is an edge computing API for Kubernetes. It exposes Helm and the cluster registry for a comprehensive UI */
type OptikonAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ClustersAddClusterHandler sets the operation handler for the add cluster operation
	ClustersAddClusterHandler clusters.AddClusterHandler
	// ReleasesAddReleasesHandler sets the operation handler for the add releases operation
	ReleasesAddReleasesHandler releases.AddReleasesHandler
	// ClustersDeleteClusterHandler sets the operation handler for the delete cluster operation
	ClustersDeleteClusterHandler clusters.DeleteClusterHandler
	// ReleasesDeleteReleaseHandler sets the operation handler for the delete release operation
	ReleasesDeleteReleaseHandler releases.DeleteReleaseHandler
	// ClustersGetClusterByIDHandler sets the operation handler for the get cluster by Id operation
	ClustersGetClusterByIDHandler clusters.GetClusterByIDHandler
	// ClustersGetClustersHandler sets the operation handler for the get clusters operation
	ClustersGetClustersHandler clusters.GetClustersHandler
	// ReleasesGetReleaseByIDHandler sets the operation handler for the get release by Id operation
	ReleasesGetReleaseByIDHandler releases.GetReleaseByIDHandler
	// ReleasesGetReleasesHandler sets the operation handler for the get releases operation
	ReleasesGetReleasesHandler releases.GetReleasesHandler
	// ClustersUpdateClusterHandler sets the operation handler for the update cluster operation
	ClustersUpdateClusterHandler clusters.UpdateClusterHandler
	// ReleasesUpdateReleaseHandler sets the operation handler for the update release operation
	ReleasesUpdateReleaseHandler releases.UpdateReleaseHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *OptikonAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *OptikonAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *OptikonAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *OptikonAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *OptikonAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *OptikonAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *OptikonAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the OptikonAPI
func (o *OptikonAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ClustersAddClusterHandler == nil {
		unregistered = append(unregistered, "clusters.AddClusterHandler")
	}

	if o.ReleasesAddReleasesHandler == nil {
		unregistered = append(unregistered, "releases.AddReleasesHandler")
	}

	if o.ClustersDeleteClusterHandler == nil {
		unregistered = append(unregistered, "clusters.DeleteClusterHandler")
	}

	if o.ReleasesDeleteReleaseHandler == nil {
		unregistered = append(unregistered, "releases.DeleteReleaseHandler")
	}

	if o.ClustersGetClusterByIDHandler == nil {
		unregistered = append(unregistered, "clusters.GetClusterByIDHandler")
	}

	if o.ClustersGetClustersHandler == nil {
		unregistered = append(unregistered, "clusters.GetClustersHandler")
	}

	if o.ReleasesGetReleaseByIDHandler == nil {
		unregistered = append(unregistered, "releases.GetReleaseByIDHandler")
	}

	if o.ReleasesGetReleasesHandler == nil {
		unregistered = append(unregistered, "releases.GetReleasesHandler")
	}

	if o.ClustersUpdateClusterHandler == nil {
		unregistered = append(unregistered, "clusters.UpdateClusterHandler")
	}

	if o.ReleasesUpdateReleaseHandler == nil {
		unregistered = append(unregistered, "releases.UpdateReleaseHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *OptikonAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *OptikonAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *OptikonAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *OptikonAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *OptikonAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *OptikonAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the optikon API
func (o *OptikonAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *OptikonAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters"] = clusters.NewAddCluster(o.context, o.ClustersAddClusterHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/releases"] = releases.NewAddReleases(o.context, o.ReleasesAddReleasesHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/clusters/{clusterId}"] = clusters.NewDeleteCluster(o.context, o.ClustersDeleteClusterHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/releases/{releaseId}"] = releases.NewDeleteRelease(o.context, o.ReleasesDeleteReleaseHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{clusterId}"] = clusters.NewGetClusterByID(o.context, o.ClustersGetClusterByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters"] = clusters.NewGetClusters(o.context, o.ClustersGetClustersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/releases/{releaseId}"] = releases.NewGetReleaseByID(o.context, o.ReleasesGetReleaseByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/releases"] = releases.NewGetReleases(o.context, o.ReleasesGetReleasesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/clusters/{clusterId}"] = clusters.NewUpdateCluster(o.context, o.ClustersUpdateClusterHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/releases/{releaseId}"] = releases.NewUpdateRelease(o.context, o.ReleasesUpdateReleaseHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *OptikonAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *OptikonAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
