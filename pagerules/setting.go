// File generated from our OpenAPI spec by Stainless.

package pagerules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Returns a list of settings (and their details) that Page Rules can apply to
// matching requests.
func (r *SettingService) List(ctx context.Context, query SettingListParams, opts ...option.RequestOption) (res *ZonesSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingListResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZonesSettings []interface{}

type SettingListResponse = interface{}

type SettingListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingListResponseEnvelope struct {
	Errors   []SettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingListResponseEnvelopeMessages `json:"messages,required"`
	// Settings available for the zone.
	Result ZonesSettings `json:"result,required"`
	// Whether the API call was successful
	Success SettingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingListResponseEnvelopeJSON    `json:"-"`
}

// settingListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingListResponseEnvelope]
type settingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    settingListResponseEnvelopeErrorsJSON `json:"-"`
}

// settingListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SettingListResponseEnvelopeErrors]
type settingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingListResponseEnvelopeMessagesJSON `json:"-"`
}

// settingListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingListResponseEnvelopeMessages]
type settingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingListResponseEnvelopeSuccess bool

const (
	SettingListResponseEnvelopeSuccessTrue SettingListResponseEnvelopeSuccess = true
)
