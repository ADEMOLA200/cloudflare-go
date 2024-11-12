// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
)

type ASN = int64

type ASNParam = int64

type AuditLog struct {
	// A string that uniquely identifies the audit log.
	ID     string         `json:"id"`
	Action AuditLogAction `json:"action"`
	Actor  AuditLogActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string           `json:"oldValue"`
	Owner    AuditLogOwner    `json:"owner"`
	Resource AuditLogResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time    `json:"when" format:"date-time"`
	JSON auditLogJSON `json:"-"`
}

// auditLogJSON contains the JSON metadata for the struct [AuditLog]
type auditLogJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Interface   apijson.Field
	Metadata    apijson.Field
	NewValue    apijson.Field
	OldValue    apijson.Field
	Owner       apijson.Field
	Resource    apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogJSON) RawJSON() string {
	return r.raw
}

type AuditLogAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string             `json:"type"`
	JSON auditLogActionJSON `json:"-"`
}

// auditLogActionJSON contains the JSON metadata for the struct [AuditLogAction]
type auditLogActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogActionJSON) RawJSON() string {
	return r.raw
}

type AuditLogActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type AuditLogActorType `json:"type"`
	JSON auditLogActorJSON `json:"-"`
}

// auditLogActorJSON contains the JSON metadata for the struct [AuditLogActor]
type auditLogActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogActorJSON) RawJSON() string {
	return r.raw
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type AuditLogActorType string

const (
	AuditLogActorTypeUser       AuditLogActorType = "user"
	AuditLogActorTypeAdmin      AuditLogActorType = "admin"
	AuditLogActorTypeCloudflare AuditLogActorType = "Cloudflare"
)

func (r AuditLogActorType) IsKnown() bool {
	switch r {
	case AuditLogActorTypeUser, AuditLogActorTypeAdmin, AuditLogActorTypeCloudflare:
		return true
	}
	return false
}

type AuditLogOwner struct {
	// Identifier
	ID   string            `json:"id"`
	JSON auditLogOwnerJSON `json:"-"`
}

// auditLogOwnerJSON contains the JSON metadata for the struct [AuditLogOwner]
type auditLogOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogOwnerJSON) RawJSON() string {
	return r.raw
}

type AuditLogResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string               `json:"type"`
	JSON auditLogResourceJSON `json:"-"`
}

// auditLogResourceJSON contains the JSON metadata for the struct
// [AuditLogResource]
type auditLogResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogResourceJSON) RawJSON() string {
	return r.raw
}

// The Certificate Authority that will issue the certificate
type CertificateCA string

const (
	CertificateCADigicert    CertificateCA = "digicert"
	CertificateCAGoogle      CertificateCA = "google"
	CertificateCALetsEncrypt CertificateCA = "lets_encrypt"
	CertificateCASSLCom      CertificateCA = "ssl_com"
)

func (r CertificateCA) IsKnown() bool {
	switch r {
	case CertificateCADigicert, CertificateCAGoogle, CertificateCALetsEncrypt, CertificateCASSLCom:
		return true
	}
	return false
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type CertificateRequestType string

const (
	CertificateRequestTypeOriginRSA          CertificateRequestType = "origin-rsa"
	CertificateRequestTypeOriginECC          CertificateRequestType = "origin-ecc"
	CertificateRequestTypeKeylessCertificate CertificateRequestType = "keyless-certificate"
)

func (r CertificateRequestType) IsKnown() bool {
	switch r {
	case CertificateRequestTypeOriginRSA, CertificateRequestTypeOriginECC, CertificateRequestTypeKeylessCertificate:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CloudflareTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []CloudflareTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status CloudflareTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType CloudflareTunnelTunType `json:"tun_type"`
	JSON    cloudflareTunnelJSON    `json:"-"`
}

// cloudflareTunnelJSON contains the JSON metadata for the struct
// [CloudflareTunnel]
type cloudflareTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r CloudflareTunnel) ImplementsWARPConnectorWARPConnectorNewResponse() {}

func (r CloudflareTunnel) ImplementsWARPConnectorWARPConnectorListResponse() {}

func (r CloudflareTunnel) ImplementsWARPConnectorWARPConnectorDeleteResponse() {}

func (r CloudflareTunnel) ImplementsWARPConnectorWARPConnectorEditResponse() {}

func (r CloudflareTunnel) ImplementsWARPConnectorWARPConnectorGetResponse() {}

func (r CloudflareTunnel) ImplementsZeroTrustTunnelNewResponse() {}

func (r CloudflareTunnel) ImplementsZeroTrustTunnelListResponse() {}

func (r CloudflareTunnel) ImplementsZeroTrustTunnelDeleteResponse() {}

func (r CloudflareTunnel) ImplementsZeroTrustTunnelEditResponse() {}

func (r CloudflareTunnel) ImplementsZeroTrustTunnelGetResponse() {}

type CloudflareTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string                         `json:"uuid" format:"uuid"`
	JSON cloudflareTunnelConnectionJSON `json:"-"`
}

// cloudflareTunnelConnectionJSON contains the JSON metadata for the struct
// [CloudflareTunnelConnection]
type cloudflareTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudflareTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudflareTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type CloudflareTunnelStatus string

const (
	CloudflareTunnelStatusInactive CloudflareTunnelStatus = "inactive"
	CloudflareTunnelStatusDegraded CloudflareTunnelStatus = "degraded"
	CloudflareTunnelStatusHealthy  CloudflareTunnelStatus = "healthy"
	CloudflareTunnelStatusDown     CloudflareTunnelStatus = "down"
)

func (r CloudflareTunnelStatus) IsKnown() bool {
	switch r {
	case CloudflareTunnelStatusInactive, CloudflareTunnelStatusDegraded, CloudflareTunnelStatusHealthy, CloudflareTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type CloudflareTunnelTunType string

const (
	CloudflareTunnelTunTypeCfdTunnel     CloudflareTunnelTunType = "cfd_tunnel"
	CloudflareTunnelTunTypeWARPConnector CloudflareTunnelTunType = "warp_connector"
	CloudflareTunnelTunTypeIPSec         CloudflareTunnelTunType = "ip_sec"
	CloudflareTunnelTunTypeGRE           CloudflareTunnelTunType = "gre"
	CloudflareTunnelTunTypeCNI           CloudflareTunnelTunType = "cni"
)

func (r CloudflareTunnelTunType) IsKnown() bool {
	switch r {
	case CloudflareTunnelTunTypeCfdTunnel, CloudflareTunnelTunTypeWARPConnector, CloudflareTunnelTunTypeIPSec, CloudflareTunnelTunTypeGRE, CloudflareTunnelTunTypeCNI:
		return true
	}
	return false
}

type ErrorData struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	JSON    errorDataJSON `json:"-"`
}

// errorDataJSON contains the JSON metadata for the struct [ErrorData]
type errorDataJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorDataJSON) RawJSON() string {
	return r.raw
}

type Permission = string

type PermissionGrant struct {
	Read  bool                `json:"read"`
	Write bool                `json:"write"`
	JSON  permissionGrantJSON `json:"-"`
}

// permissionGrantJSON contains the JSON metadata for the struct [PermissionGrant]
type permissionGrantJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionGrantJSON) RawJSON() string {
	return r.raw
}

type PermissionGrantParam struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r PermissionGrantParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rate plan applied to the subscription.
type RatePlan struct {
	// The ID of the rate plan.
	ID string `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency string `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged bool `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract bool `json:"is_contract"`
	// The full name of the rate plan.
	PublicName string `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope string `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets []string     `json:"sets"`
	JSON ratePlanJSON `json:"-"`
}

// ratePlanJSON contains the JSON metadata for the struct [RatePlan]
type ratePlanJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	IsContract        apijson.Field
	PublicName        apijson.Field
	Scope             apijson.Field
	Sets              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanJSON) RawJSON() string {
	return r.raw
}

// The rate plan applied to the subscription.
type RatePlanParam struct {
	// The ID of the rate plan.
	ID param.Field[string] `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency param.Field[string] `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged param.Field[bool] `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract param.Field[bool] `json:"is_contract"`
	// The full name of the rate plan.
	PublicName param.Field[string] `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope param.Field[string] `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets param.Field[[]string] `json:"sets"`
}

func (r RatePlanParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseInfo struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    responseInfoJSON `json:"-"`
}

// responseInfoJSON contains the JSON metadata for the struct [ResponseInfo]
type responseInfoJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseInfoJSON) RawJSON() string {
	return r.raw
}

// Direction to order DNS records in.
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func (r SortDirection) IsKnown() bool {
	switch r {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

type Subscription struct {
	// Subscription identifier tag.
	ID string `json:"id"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan RatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionState `json:"state"`
	JSON  subscriptionJSON  `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID                 apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionFrequency string

const (
	SubscriptionFrequencyWeekly    SubscriptionFrequency = "weekly"
	SubscriptionFrequencyMonthly   SubscriptionFrequency = "monthly"
	SubscriptionFrequencyQuarterly SubscriptionFrequency = "quarterly"
	SubscriptionFrequencyYearly    SubscriptionFrequency = "yearly"
)

func (r SubscriptionFrequency) IsKnown() bool {
	switch r {
	case SubscriptionFrequencyWeekly, SubscriptionFrequencyMonthly, SubscriptionFrequencyQuarterly, SubscriptionFrequencyYearly:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionState string

const (
	SubscriptionStateTrial           SubscriptionState = "Trial"
	SubscriptionStateProvisioned     SubscriptionState = "Provisioned"
	SubscriptionStatePaid            SubscriptionState = "Paid"
	SubscriptionStateAwaitingPayment SubscriptionState = "AwaitingPayment"
	SubscriptionStateCancelled       SubscriptionState = "Cancelled"
	SubscriptionStateFailed          SubscriptionState = "Failed"
	SubscriptionStateExpired         SubscriptionState = "Expired"
)

func (r SubscriptionState) IsKnown() bool {
	switch r {
	case SubscriptionStateTrial, SubscriptionStateProvisioned, SubscriptionStatePaid, SubscriptionStateAwaitingPayment, SubscriptionStateCancelled, SubscriptionStateFailed, SubscriptionStateExpired:
		return true
	}
	return false
}

type SubscriptionParam struct {
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[RatePlanParam] `json:"rate_plan"`
}

func (r SubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
