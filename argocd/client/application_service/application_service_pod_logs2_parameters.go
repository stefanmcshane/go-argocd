// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewApplicationServicePodLogs2Params creates a new ApplicationServicePodLogs2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServicePodLogs2Params() *ApplicationServicePodLogs2Params {
	return &ApplicationServicePodLogs2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServicePodLogs2ParamsWithTimeout creates a new ApplicationServicePodLogs2Params object
// with the ability to set a timeout on a request.
func NewApplicationServicePodLogs2ParamsWithTimeout(timeout time.Duration) *ApplicationServicePodLogs2Params {
	return &ApplicationServicePodLogs2Params{
		timeout: timeout,
	}
}

// NewApplicationServicePodLogs2ParamsWithContext creates a new ApplicationServicePodLogs2Params object
// with the ability to set a context for a request.
func NewApplicationServicePodLogs2ParamsWithContext(ctx context.Context) *ApplicationServicePodLogs2Params {
	return &ApplicationServicePodLogs2Params{
		Context: ctx,
	}
}

// NewApplicationServicePodLogs2ParamsWithHTTPClient creates a new ApplicationServicePodLogs2Params object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServicePodLogs2ParamsWithHTTPClient(client *http.Client) *ApplicationServicePodLogs2Params {
	return &ApplicationServicePodLogs2Params{
		HTTPClient: client,
	}
}

/* ApplicationServicePodLogs2Params contains all the parameters to send to the API endpoint
   for the application service pod logs2 operation.

   Typically these are written to a http.Request.
*/
type ApplicationServicePodLogs2Params struct {

	// AppNamespace.
	AppNamespace *string

	// Container.
	Container *string

	// Filter.
	Filter *string

	// Follow.
	Follow *bool

	// Group.
	Group *string

	// Kind.
	Kind *string

	// Name.
	Name string

	// Namespace.
	Namespace *string

	// PodName.
	PodName *string

	// Previous.
	Previous *bool

	// ResourceName.
	ResourceName *string

	// SinceSeconds.
	//
	// Format: int64
	SinceSeconds *string

	/* SinceTimeNanos.

	     Non-negative fractions of a second at nanosecond resolution. Negative
	second values with fractions must still have non-negative nanos values
	that count forward in time. Must be from 0 to 999,999,999
	inclusive. This field may be limited in precision depending on context.

	     Format: int32
	*/
	SinceTimeNanos *int32

	/* SinceTimeSeconds.

	     Represents seconds of UTC time since Unix epoch
	1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	9999-12-31T23:59:59Z inclusive.

	     Format: int64
	*/
	SinceTimeSeconds *string

	// TailLines.
	//
	// Format: int64
	TailLines *string

	// UntilTime.
	UntilTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service pod logs2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServicePodLogs2Params) WithDefaults() *ApplicationServicePodLogs2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service pod logs2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServicePodLogs2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithTimeout(timeout time.Duration) *ApplicationServicePodLogs2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithContext(ctx context.Context) *ApplicationServicePodLogs2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithHTTPClient(client *http.Client) *ApplicationServicePodLogs2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppNamespace adds the appNamespace to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithAppNamespace(appNamespace *string) *ApplicationServicePodLogs2Params {
	o.SetAppNamespace(appNamespace)
	return o
}

// SetAppNamespace adds the appNamespace to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetAppNamespace(appNamespace *string) {
	o.AppNamespace = appNamespace
}

// WithContainer adds the container to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithContainer(container *string) *ApplicationServicePodLogs2Params {
	o.SetContainer(container)
	return o
}

// SetContainer adds the container to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetContainer(container *string) {
	o.Container = container
}

// WithFilter adds the filter to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithFilter(filter *string) *ApplicationServicePodLogs2Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithFollow adds the follow to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithFollow(follow *bool) *ApplicationServicePodLogs2Params {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithGroup adds the group to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithGroup(group *string) *ApplicationServicePodLogs2Params {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetGroup(group *string) {
	o.Group = group
}

// WithKind adds the kind to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithKind(kind *string) *ApplicationServicePodLogs2Params {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetKind(kind *string) {
	o.Kind = kind
}

// WithName adds the name to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithName(name string) *ApplicationServicePodLogs2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithNamespace(namespace *string) *ApplicationServicePodLogs2Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithPodName adds the podName to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithPodName(podName *string) *ApplicationServicePodLogs2Params {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetPodName(podName *string) {
	o.PodName = podName
}

// WithPrevious adds the previous to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithPrevious(previous *bool) *ApplicationServicePodLogs2Params {
	o.SetPrevious(previous)
	return o
}

// SetPrevious adds the previous to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetPrevious(previous *bool) {
	o.Previous = previous
}

// WithResourceName adds the resourceName to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithResourceName(resourceName *string) *ApplicationServicePodLogs2Params {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetResourceName(resourceName *string) {
	o.ResourceName = resourceName
}

// WithSinceSeconds adds the sinceSeconds to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithSinceSeconds(sinceSeconds *string) *ApplicationServicePodLogs2Params {
	o.SetSinceSeconds(sinceSeconds)
	return o
}

// SetSinceSeconds adds the sinceSeconds to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetSinceSeconds(sinceSeconds *string) {
	o.SinceSeconds = sinceSeconds
}

// WithSinceTimeNanos adds the sinceTimeNanos to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithSinceTimeNanos(sinceTimeNanos *int32) *ApplicationServicePodLogs2Params {
	o.SetSinceTimeNanos(sinceTimeNanos)
	return o
}

// SetSinceTimeNanos adds the sinceTimeNanos to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetSinceTimeNanos(sinceTimeNanos *int32) {
	o.SinceTimeNanos = sinceTimeNanos
}

// WithSinceTimeSeconds adds the sinceTimeSeconds to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithSinceTimeSeconds(sinceTimeSeconds *string) *ApplicationServicePodLogs2Params {
	o.SetSinceTimeSeconds(sinceTimeSeconds)
	return o
}

// SetSinceTimeSeconds adds the sinceTimeSeconds to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetSinceTimeSeconds(sinceTimeSeconds *string) {
	o.SinceTimeSeconds = sinceTimeSeconds
}

// WithTailLines adds the tailLines to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithTailLines(tailLines *string) *ApplicationServicePodLogs2Params {
	o.SetTailLines(tailLines)
	return o
}

// SetTailLines adds the tailLines to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetTailLines(tailLines *string) {
	o.TailLines = tailLines
}

// WithUntilTime adds the untilTime to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) WithUntilTime(untilTime *string) *ApplicationServicePodLogs2Params {
	o.SetUntilTime(untilTime)
	return o
}

// SetUntilTime adds the untilTime to the application service pod logs2 params
func (o *ApplicationServicePodLogs2Params) SetUntilTime(untilTime *string) {
	o.UntilTime = untilTime
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServicePodLogs2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppNamespace != nil {

		// query param appNamespace
		var qrAppNamespace string

		if o.AppNamespace != nil {
			qrAppNamespace = *o.AppNamespace
		}
		qAppNamespace := qrAppNamespace
		if qAppNamespace != "" {

			if err := r.SetQueryParam("appNamespace", qAppNamespace); err != nil {
				return err
			}
		}
	}

	if o.Container != nil {

		// query param container
		var qrContainer string

		if o.Container != nil {
			qrContainer = *o.Container
		}
		qContainer := qrContainer
		if qContainer != "" {

			if err := r.SetQueryParam("container", qContainer); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Follow != nil {

		// query param follow
		var qrFollow bool

		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {

			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}
	}

	if o.PodName != nil {

		// query param podName
		var qrPodName string

		if o.PodName != nil {
			qrPodName = *o.PodName
		}
		qPodName := qrPodName
		if qPodName != "" {

			if err := r.SetQueryParam("podName", qPodName); err != nil {
				return err
			}
		}
	}

	if o.Previous != nil {

		// query param previous
		var qrPrevious bool

		if o.Previous != nil {
			qrPrevious = *o.Previous
		}
		qPrevious := swag.FormatBool(qrPrevious)
		if qPrevious != "" {

			if err := r.SetQueryParam("previous", qPrevious); err != nil {
				return err
			}
		}
	}

	if o.ResourceName != nil {

		// query param resourceName
		var qrResourceName string

		if o.ResourceName != nil {
			qrResourceName = *o.ResourceName
		}
		qResourceName := qrResourceName
		if qResourceName != "" {

			if err := r.SetQueryParam("resourceName", qResourceName); err != nil {
				return err
			}
		}
	}

	if o.SinceSeconds != nil {

		// query param sinceSeconds
		var qrSinceSeconds string

		if o.SinceSeconds != nil {
			qrSinceSeconds = *o.SinceSeconds
		}
		qSinceSeconds := qrSinceSeconds
		if qSinceSeconds != "" {

			if err := r.SetQueryParam("sinceSeconds", qSinceSeconds); err != nil {
				return err
			}
		}
	}

	if o.SinceTimeNanos != nil {

		// query param sinceTime.nanos
		var qrSinceTimeNanos int32

		if o.SinceTimeNanos != nil {
			qrSinceTimeNanos = *o.SinceTimeNanos
		}
		qSinceTimeNanos := swag.FormatInt32(qrSinceTimeNanos)
		if qSinceTimeNanos != "" {

			if err := r.SetQueryParam("sinceTime.nanos", qSinceTimeNanos); err != nil {
				return err
			}
		}
	}

	if o.SinceTimeSeconds != nil {

		// query param sinceTime.seconds
		var qrSinceTimeSeconds string

		if o.SinceTimeSeconds != nil {
			qrSinceTimeSeconds = *o.SinceTimeSeconds
		}
		qSinceTimeSeconds := qrSinceTimeSeconds
		if qSinceTimeSeconds != "" {

			if err := r.SetQueryParam("sinceTime.seconds", qSinceTimeSeconds); err != nil {
				return err
			}
		}
	}

	if o.TailLines != nil {

		// query param tailLines
		var qrTailLines string

		if o.TailLines != nil {
			qrTailLines = *o.TailLines
		}
		qTailLines := qrTailLines
		if qTailLines != "" {

			if err := r.SetQueryParam("tailLines", qTailLines); err != nil {
				return err
			}
		}
	}

	if o.UntilTime != nil {

		// query param untilTime
		var qrUntilTime string

		if o.UntilTime != nil {
			qrUntilTime = *o.UntilTime
		}
		qUntilTime := qrUntilTime
		if qUntilTime != "" {

			if err := r.SetQueryParam("untilTime", qUntilTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
