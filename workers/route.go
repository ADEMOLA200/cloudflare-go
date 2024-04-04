// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RouteService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRouteService] method instead.
type RouteService struct {
	Options []option.RequestOption
}

// NewRouteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRouteService(opts ...option.RequestOption) (r *RouteService) {
	r = &RouteService{}
	r.Options = opts
	return
}

// Creates a route that maps a URL pattern to a Worker.
func (r *RouteService) New(ctx context.Context, params RouteNewParams, opts ...option.RequestOption) (res *RouteNewResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *RouteService) Update(ctx context.Context, routeID string, params RouteUpdateParams, opts ...option.RequestOption) (res *WorkersRoute, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", params.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns routes for a zone.
func (r *RouteService) List(ctx context.Context, query RouteListParams, opts ...option.RequestOption) (res *pagination.SinglePage[WorkersRoute], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/workers/routes", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns routes for a zone.
func (r *RouteService) ListAutoPaging(ctx context.Context, query RouteListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WorkersRoute] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a route.
func (r *RouteService) Delete(ctx context.Context, routeID string, params RouteDeleteParams, opts ...option.RequestOption) (res *RouteDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", params.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *RouteService) Get(ctx context.Context, routeID string, query RouteGetParams, opts ...option.RequestOption) (res *WorkersRoute, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", query.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersRoute struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string           `json:"script,required"`
	JSON   workersRouteJSON `json:"-"`
}

// workersRouteJSON contains the JSON metadata for the struct [WorkersRoute]
type workersRouteJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersRouteJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [workers.RouteNewResponseUnknown] or [shared.UnionString].
type RouteNewResponseUnion interface {
	ImplementsWorkersRouteNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RouteNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [workers.RouteDeleteResponseUnknown] or [shared.UnionString].
type RouteDeleteResponseUnion interface {
	ImplementsWorkersRouteDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RouteDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RouteNewParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r RouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteNewResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeNewResponseEnvelopeJSON    `json:"-"`
}

// routeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelope]
type routeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteNewResponseEnvelopeSuccess bool

const (
	RouteNewResponseEnvelopeSuccessTrue RouteNewResponseEnvelopeSuccess = true
)

func (r RouteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteUpdateParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r RouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   WorkersRoute          `json:"result,required"`
	// Whether the API call was successful
	Success RouteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeUpdateResponseEnvelopeJSON    `json:"-"`
}

// routeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteUpdateResponseEnvelope]
type routeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteUpdateResponseEnvelopeSuccess bool

const (
	RouteUpdateResponseEnvelopeSuccessTrue RouteUpdateResponseEnvelopeSuccess = true
)

func (r RouteUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r RouteDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RouteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   RouteDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeDeleteResponseEnvelopeJSON    `json:"-"`
}

// routeDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteDeleteResponseEnvelope]
type routeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteDeleteResponseEnvelopeSuccess bool

const (
	RouteDeleteResponseEnvelopeSuccessTrue RouteDeleteResponseEnvelopeSuccess = true
)

func (r RouteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   WorkersRoute          `json:"result,required"`
	// Whether the API call was successful
	Success RouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeGetResponseEnvelopeJSON    `json:"-"`
}

// routeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelope]
type routeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteGetResponseEnvelopeSuccess bool

const (
	RouteGetResponseEnvelopeSuccessTrue RouteGetResponseEnvelopeSuccess = true
)

func (r RouteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
