// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateClusterInstallProgressHandlerFunc turns a function with the right signature into a update cluster install progress handler
type UpdateClusterInstallProgressHandlerFunc func(UpdateClusterInstallProgressParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateClusterInstallProgressHandlerFunc) Handle(params UpdateClusterInstallProgressParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateClusterInstallProgressHandler interface for that can handle valid update cluster install progress params
type UpdateClusterInstallProgressHandler interface {
	Handle(UpdateClusterInstallProgressParams, interface{}) middleware.Responder
}

// NewUpdateClusterInstallProgress creates a new http.Handler for the update cluster install progress operation
func NewUpdateClusterInstallProgress(ctx *middleware.Context, handler UpdateClusterInstallProgressHandler) *UpdateClusterInstallProgress {
	return &UpdateClusterInstallProgress{Context: ctx, Handler: handler}
}

/*UpdateClusterInstallProgress swagger:route PUT /clusters/{cluster_id}/progress installer updateClusterInstallProgress

Update cluster installation progress.

*/
type UpdateClusterInstallProgress struct {
	Context *middleware.Context
	Handler UpdateClusterInstallProgressHandler
}

func (o *UpdateClusterInstallProgress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateClusterInstallProgressParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
