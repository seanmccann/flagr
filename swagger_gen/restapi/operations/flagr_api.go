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

	"github.com/checkr/flagr/swagger_gen/restapi/operations/constraint"
	"github.com/checkr/flagr/swagger_gen/restapi/operations/distribution"
	"github.com/checkr/flagr/swagger_gen/restapi/operations/evaluation"
	"github.com/checkr/flagr/swagger_gen/restapi/operations/flag"
	"github.com/checkr/flagr/swagger_gen/restapi/operations/segment"
	"github.com/checkr/flagr/swagger_gen/restapi/operations/variant"
)

// NewFlagrAPI creates a new Flagr instance
func NewFlagrAPI(spec *loads.Document) *FlagrAPI {
	return &FlagrAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ConstraintCreateConstraintHandler: constraint.CreateConstraintHandlerFunc(func(params constraint.CreateConstraintParams) middleware.Responder {
			return middleware.NotImplemented("operation ConstraintCreateConstraint has not yet been implemented")
		}),
		FlagCreateFlagHandler: flag.CreateFlagHandlerFunc(func(params flag.CreateFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation FlagCreateFlag has not yet been implemented")
		}),
		SegmentCreateSegmentHandler: segment.CreateSegmentHandlerFunc(func(params segment.CreateSegmentParams) middleware.Responder {
			return middleware.NotImplemented("operation SegmentCreateSegment has not yet been implemented")
		}),
		VariantCreateVariantHandler: variant.CreateVariantHandlerFunc(func(params variant.CreateVariantParams) middleware.Responder {
			return middleware.NotImplemented("operation VariantCreateVariant has not yet been implemented")
		}),
		ConstraintDeleteConstraintHandler: constraint.DeleteConstraintHandlerFunc(func(params constraint.DeleteConstraintParams) middleware.Responder {
			return middleware.NotImplemented("operation ConstraintDeleteConstraint has not yet been implemented")
		}),
		FlagDeleteFlagHandler: flag.DeleteFlagHandlerFunc(func(params flag.DeleteFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation FlagDeleteFlag has not yet been implemented")
		}),
		SegmentDeleteSegmentHandler: segment.DeleteSegmentHandlerFunc(func(params segment.DeleteSegmentParams) middleware.Responder {
			return middleware.NotImplemented("operation SegmentDeleteSegment has not yet been implemented")
		}),
		VariantDeleteVariantHandler: variant.DeleteVariantHandlerFunc(func(params variant.DeleteVariantParams) middleware.Responder {
			return middleware.NotImplemented("operation VariantDeleteVariant has not yet been implemented")
		}),
		ConstraintFindConstraintsHandler: constraint.FindConstraintsHandlerFunc(func(params constraint.FindConstraintsParams) middleware.Responder {
			return middleware.NotImplemented("operation ConstraintFindConstraints has not yet been implemented")
		}),
		DistributionFindDistributionsHandler: distribution.FindDistributionsHandlerFunc(func(params distribution.FindDistributionsParams) middleware.Responder {
			return middleware.NotImplemented("operation DistributionFindDistributions has not yet been implemented")
		}),
		FlagFindFlagsHandler: flag.FindFlagsHandlerFunc(func(params flag.FindFlagsParams) middleware.Responder {
			return middleware.NotImplemented("operation FlagFindFlags has not yet been implemented")
		}),
		SegmentFindSegmentsHandler: segment.FindSegmentsHandlerFunc(func(params segment.FindSegmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation SegmentFindSegments has not yet been implemented")
		}),
		VariantFindVariantsHandler: variant.FindVariantsHandlerFunc(func(params variant.FindVariantsParams) middleware.Responder {
			return middleware.NotImplemented("operation VariantFindVariants has not yet been implemented")
		}),
		FlagGetFlagHandler: flag.GetFlagHandlerFunc(func(params flag.GetFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation FlagGetFlag has not yet been implemented")
		}),
		EvaluationPostEvaluationHandler: evaluation.PostEvaluationHandlerFunc(func(params evaluation.PostEvaluationParams) middleware.Responder {
			return middleware.NotImplemented("operation EvaluationPostEvaluation has not yet been implemented")
		}),
		EvaluationPostEvaluationBatchHandler: evaluation.PostEvaluationBatchHandlerFunc(func(params evaluation.PostEvaluationBatchParams) middleware.Responder {
			return middleware.NotImplemented("operation EvaluationPostEvaluationBatch has not yet been implemented")
		}),
		ConstraintPutConstraintHandler: constraint.PutConstraintHandlerFunc(func(params constraint.PutConstraintParams) middleware.Responder {
			return middleware.NotImplemented("operation ConstraintPutConstraint has not yet been implemented")
		}),
		DistributionPutDistributionsHandler: distribution.PutDistributionsHandlerFunc(func(params distribution.PutDistributionsParams) middleware.Responder {
			return middleware.NotImplemented("operation DistributionPutDistributions has not yet been implemented")
		}),
		FlagPutFlagHandler: flag.PutFlagHandlerFunc(func(params flag.PutFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation FlagPutFlag has not yet been implemented")
		}),
		SegmentPutSegmentHandler: segment.PutSegmentHandlerFunc(func(params segment.PutSegmentParams) middleware.Responder {
			return middleware.NotImplemented("operation SegmentPutSegment has not yet been implemented")
		}),
		SegmentPutSegmentsReorderHandler: segment.PutSegmentsReorderHandlerFunc(func(params segment.PutSegmentsReorderParams) middleware.Responder {
			return middleware.NotImplemented("operation SegmentPutSegmentsReorder has not yet been implemented")
		}),
		VariantPutVariantHandler: variant.PutVariantHandlerFunc(func(params variant.PutVariantParams) middleware.Responder {
			return middleware.NotImplemented("operation VariantPutVariant has not yet been implemented")
		}),
		FlagSetFlagEnabledHandler: flag.SetFlagEnabledHandlerFunc(func(params flag.SetFlagEnabledParams) middleware.Responder {
			return middleware.NotImplemented("operation FlagSetFlagEnabled has not yet been implemented")
		}),
	}
}

/*FlagrAPI Flagr is a feature flagging, A/B testing and dynamic configuration microservice */
type FlagrAPI struct {
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

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ConstraintCreateConstraintHandler sets the operation handler for the create constraint operation
	ConstraintCreateConstraintHandler constraint.CreateConstraintHandler
	// FlagCreateFlagHandler sets the operation handler for the create flag operation
	FlagCreateFlagHandler flag.CreateFlagHandler
	// SegmentCreateSegmentHandler sets the operation handler for the create segment operation
	SegmentCreateSegmentHandler segment.CreateSegmentHandler
	// VariantCreateVariantHandler sets the operation handler for the create variant operation
	VariantCreateVariantHandler variant.CreateVariantHandler
	// ConstraintDeleteConstraintHandler sets the operation handler for the delete constraint operation
	ConstraintDeleteConstraintHandler constraint.DeleteConstraintHandler
	// FlagDeleteFlagHandler sets the operation handler for the delete flag operation
	FlagDeleteFlagHandler flag.DeleteFlagHandler
	// SegmentDeleteSegmentHandler sets the operation handler for the delete segment operation
	SegmentDeleteSegmentHandler segment.DeleteSegmentHandler
	// VariantDeleteVariantHandler sets the operation handler for the delete variant operation
	VariantDeleteVariantHandler variant.DeleteVariantHandler
	// ConstraintFindConstraintsHandler sets the operation handler for the find constraints operation
	ConstraintFindConstraintsHandler constraint.FindConstraintsHandler
	// DistributionFindDistributionsHandler sets the operation handler for the find distributions operation
	DistributionFindDistributionsHandler distribution.FindDistributionsHandler
	// FlagFindFlagsHandler sets the operation handler for the find flags operation
	FlagFindFlagsHandler flag.FindFlagsHandler
	// SegmentFindSegmentsHandler sets the operation handler for the find segments operation
	SegmentFindSegmentsHandler segment.FindSegmentsHandler
	// VariantFindVariantsHandler sets the operation handler for the find variants operation
	VariantFindVariantsHandler variant.FindVariantsHandler
	// FlagGetFlagHandler sets the operation handler for the get flag operation
	FlagGetFlagHandler flag.GetFlagHandler
	// EvaluationPostEvaluationHandler sets the operation handler for the post evaluation operation
	EvaluationPostEvaluationHandler evaluation.PostEvaluationHandler
	// EvaluationPostEvaluationBatchHandler sets the operation handler for the post evaluation batch operation
	EvaluationPostEvaluationBatchHandler evaluation.PostEvaluationBatchHandler
	// ConstraintPutConstraintHandler sets the operation handler for the put constraint operation
	ConstraintPutConstraintHandler constraint.PutConstraintHandler
	// DistributionPutDistributionsHandler sets the operation handler for the put distributions operation
	DistributionPutDistributionsHandler distribution.PutDistributionsHandler
	// FlagPutFlagHandler sets the operation handler for the put flag operation
	FlagPutFlagHandler flag.PutFlagHandler
	// SegmentPutSegmentHandler sets the operation handler for the put segment operation
	SegmentPutSegmentHandler segment.PutSegmentHandler
	// SegmentPutSegmentsReorderHandler sets the operation handler for the put segments reorder operation
	SegmentPutSegmentsReorderHandler segment.PutSegmentsReorderHandler
	// VariantPutVariantHandler sets the operation handler for the put variant operation
	VariantPutVariantHandler variant.PutVariantHandler
	// FlagSetFlagEnabledHandler sets the operation handler for the set flag enabled operation
	FlagSetFlagEnabledHandler flag.SetFlagEnabledHandler

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
func (o *FlagrAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FlagrAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FlagrAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FlagrAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FlagrAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FlagrAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FlagrAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FlagrAPI
func (o *FlagrAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ConstraintCreateConstraintHandler == nil {
		unregistered = append(unregistered, "constraint.CreateConstraintHandler")
	}

	if o.FlagCreateFlagHandler == nil {
		unregistered = append(unregistered, "flag.CreateFlagHandler")
	}

	if o.SegmentCreateSegmentHandler == nil {
		unregistered = append(unregistered, "segment.CreateSegmentHandler")
	}

	if o.VariantCreateVariantHandler == nil {
		unregistered = append(unregistered, "variant.CreateVariantHandler")
	}

	if o.ConstraintDeleteConstraintHandler == nil {
		unregistered = append(unregistered, "constraint.DeleteConstraintHandler")
	}

	if o.FlagDeleteFlagHandler == nil {
		unregistered = append(unregistered, "flag.DeleteFlagHandler")
	}

	if o.SegmentDeleteSegmentHandler == nil {
		unregistered = append(unregistered, "segment.DeleteSegmentHandler")
	}

	if o.VariantDeleteVariantHandler == nil {
		unregistered = append(unregistered, "variant.DeleteVariantHandler")
	}

	if o.ConstraintFindConstraintsHandler == nil {
		unregistered = append(unregistered, "constraint.FindConstraintsHandler")
	}

	if o.DistributionFindDistributionsHandler == nil {
		unregistered = append(unregistered, "distribution.FindDistributionsHandler")
	}

	if o.FlagFindFlagsHandler == nil {
		unregistered = append(unregistered, "flag.FindFlagsHandler")
	}

	if o.SegmentFindSegmentsHandler == nil {
		unregistered = append(unregistered, "segment.FindSegmentsHandler")
	}

	if o.VariantFindVariantsHandler == nil {
		unregistered = append(unregistered, "variant.FindVariantsHandler")
	}

	if o.FlagGetFlagHandler == nil {
		unregistered = append(unregistered, "flag.GetFlagHandler")
	}

	if o.EvaluationPostEvaluationHandler == nil {
		unregistered = append(unregistered, "evaluation.PostEvaluationHandler")
	}

	if o.EvaluationPostEvaluationBatchHandler == nil {
		unregistered = append(unregistered, "evaluation.PostEvaluationBatchHandler")
	}

	if o.ConstraintPutConstraintHandler == nil {
		unregistered = append(unregistered, "constraint.PutConstraintHandler")
	}

	if o.DistributionPutDistributionsHandler == nil {
		unregistered = append(unregistered, "distribution.PutDistributionsHandler")
	}

	if o.FlagPutFlagHandler == nil {
		unregistered = append(unregistered, "flag.PutFlagHandler")
	}

	if o.SegmentPutSegmentHandler == nil {
		unregistered = append(unregistered, "segment.PutSegmentHandler")
	}

	if o.SegmentPutSegmentsReorderHandler == nil {
		unregistered = append(unregistered, "segment.PutSegmentsReorderHandler")
	}

	if o.VariantPutVariantHandler == nil {
		unregistered = append(unregistered, "variant.PutVariantHandler")
	}

	if o.FlagSetFlagEnabledHandler == nil {
		unregistered = append(unregistered, "flag.SetFlagEnabledHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FlagrAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FlagrAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *FlagrAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *FlagrAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *FlagrAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

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
func (o *FlagrAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the flagr API
func (o *FlagrAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FlagrAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/segments/{segmentID}/constraints"] = constraint.NewCreateConstraint(o.context, o.ConstraintCreateConstraintHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags"] = flag.NewCreateFlag(o.context, o.FlagCreateFlagHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/segments"] = segment.NewCreateSegment(o.context, o.SegmentCreateSegmentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/variants"] = variant.NewCreateVariant(o.context, o.VariantCreateVariantHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/segments/{segmentID}/constraints/{constraintID}"] = constraint.NewDeleteConstraint(o.context, o.ConstraintDeleteConstraintHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}"] = flag.NewDeleteFlag(o.context, o.FlagDeleteFlagHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/segments/{segmentID}"] = segment.NewDeleteSegment(o.context, o.SegmentDeleteSegmentHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/variants/{variantID}"] = variant.NewDeleteVariant(o.context, o.VariantDeleteVariantHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/segments/{segmentID}/constraints"] = constraint.NewFindConstraints(o.context, o.ConstraintFindConstraintsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/segments/{segmentID}/distributions"] = distribution.NewFindDistributions(o.context, o.DistributionFindDistributionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags"] = flag.NewFindFlags(o.context, o.FlagFindFlagsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/segments"] = segment.NewFindSegments(o.context, o.SegmentFindSegmentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/variants"] = variant.NewFindVariants(o.context, o.VariantFindVariantsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}"] = flag.NewGetFlag(o.context, o.FlagGetFlagHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/evaluation"] = evaluation.NewPostEvaluation(o.context, o.EvaluationPostEvaluationHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/evaluation/batch"] = evaluation.NewPostEvaluationBatch(o.context, o.EvaluationPostEvaluationBatchHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/{segmentID}/constraints/{constraintID}"] = constraint.NewPutConstraint(o.context, o.ConstraintPutConstraintHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/{segmentID}/distributions"] = distribution.NewPutDistributions(o.context, o.DistributionPutDistributionsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}"] = flag.NewPutFlag(o.context, o.FlagPutFlagHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/{segmentID}"] = segment.NewPutSegment(o.context, o.SegmentPutSegmentHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/reorder"] = segment.NewPutSegmentsReorder(o.context, o.SegmentPutSegmentsReorderHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/variants/{variantID}"] = variant.NewPutVariant(o.context, o.VariantPutVariantHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/enabled"] = flag.NewSetFlagEnabled(o.context, o.FlagSetFlagEnabledHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FlagrAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *FlagrAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
