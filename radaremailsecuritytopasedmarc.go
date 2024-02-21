// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityTopAseDmarcService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseDmarcService] method instead.
type RadarEmailSecurityTopAseDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseDmarcService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseDmarcService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseDmarcService) {
	r = &RadarEmailSecurityTopAseDmarcService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by emails DMARC validation.
func (r *RadarEmailSecurityTopAseDmarcService) Get(ctx context.Context, dmarc RadarEmailSecurityTopAseDmarcGetParamsDmarc, query RadarEmailSecurityTopAseDmarcGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseDmarcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseDmarcGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseDmarcGetResponse struct {
	Meta RadarEmailSecurityTopAseDmarcGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseDmarcGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseDmarcGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseDmarcGetResponse]
type radarEmailSecurityTopAseDmarcGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDmarcGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseDmarcGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseDmarcGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDmarcGetResponseMeta]
type radarEmailSecurityTopAseDmarcGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDmarcGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseDmarcGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseDmarcGetResponseMetaDateRange]
type radarEmailSecurityTopAseDmarcGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDmarcGetResponseTop0 struct {
	ClientAsn    int64                                            `json:"clientASN,required"`
	ClientAsName string                                           `json:"clientASName,required"`
	Value        string                                           `json:"value,required"`
	JSON         radarEmailSecurityTopAseDmarcGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDmarcGetResponseTop0]
type radarEmailSecurityTopAseDmarcGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDmarcGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseDmarcGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseDmarcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseDmarcGetParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseDmarcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseDmarcGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseDmarcGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopAseDmarcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DMARC.
type RadarEmailSecurityTopAseDmarcGetParamsDmarc string

const (
	RadarEmailSecurityTopAseDmarcGetParamsDmarcPass RadarEmailSecurityTopAseDmarcGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseDmarcGetParamsDmarcNone RadarEmailSecurityTopAseDmarcGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseDmarcGetParamsDmarcFail RadarEmailSecurityTopAseDmarcGetParamsDmarc = "FAIL"
)

type RadarEmailSecurityTopAseDmarcGetParamsArc string

const (
	RadarEmailSecurityTopAseDmarcGetParamsArcPass RadarEmailSecurityTopAseDmarcGetParamsArc = "PASS"
	RadarEmailSecurityTopAseDmarcGetParamsArcNone RadarEmailSecurityTopAseDmarcGetParamsArc = "NONE"
	RadarEmailSecurityTopAseDmarcGetParamsArcFail RadarEmailSecurityTopAseDmarcGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseDmarcGetParamsDateRange string

const (
	RadarEmailSecurityTopAseDmarcGetParamsDateRange1d         RadarEmailSecurityTopAseDmarcGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange2d         RadarEmailSecurityTopAseDmarcGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange7d         RadarEmailSecurityTopAseDmarcGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange14d        RadarEmailSecurityTopAseDmarcGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange28d        RadarEmailSecurityTopAseDmarcGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange12w        RadarEmailSecurityTopAseDmarcGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange24w        RadarEmailSecurityTopAseDmarcGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange52w        RadarEmailSecurityTopAseDmarcGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange1dControl  RadarEmailSecurityTopAseDmarcGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange2dControl  RadarEmailSecurityTopAseDmarcGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange7dControl  RadarEmailSecurityTopAseDmarcGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange14dControl RadarEmailSecurityTopAseDmarcGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange28dControl RadarEmailSecurityTopAseDmarcGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange12wControl RadarEmailSecurityTopAseDmarcGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseDmarcGetParamsDateRange24wControl RadarEmailSecurityTopAseDmarcGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseDmarcGetParamsDKIM string

const (
	RadarEmailSecurityTopAseDmarcGetParamsDKIMPass RadarEmailSecurityTopAseDmarcGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseDmarcGetParamsDKIMNone RadarEmailSecurityTopAseDmarcGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseDmarcGetParamsDKIMFail RadarEmailSecurityTopAseDmarcGetParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseDmarcGetParamsFormat string

const (
	RadarEmailSecurityTopAseDmarcGetParamsFormatJson RadarEmailSecurityTopAseDmarcGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseDmarcGetParamsFormatCsv  RadarEmailSecurityTopAseDmarcGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseDmarcGetParamsSPF string

const (
	RadarEmailSecurityTopAseDmarcGetParamsSPFPass RadarEmailSecurityTopAseDmarcGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseDmarcGetParamsSPFNone RadarEmailSecurityTopAseDmarcGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseDmarcGetParamsSPFFail RadarEmailSecurityTopAseDmarcGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseDmarcGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseDmarcGetResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecurityTopAseDmarcGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseDmarcGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseDmarcGetResponseEnvelope]
type radarEmailSecurityTopAseDmarcGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDmarcGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
