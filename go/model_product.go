/*
SwissSign RA REST API

See https://github.com/SwissSign-AG/RaApi/README.md

API version: 2.2.1
Contact: ssc@swisssign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package swisssign_ra_api.v2

import (
	"encoding/json"
)

// Product struct for Product
type Product struct {
	// Product UUID reference
	Uuid string `json:"uuid"`
	// Certificate product name
	ProductName *string `json:"productName,omitempty"`
	// Product additional description
	ProductDescription NullableString `json:"productDescription,omitempty"`
	// Key generation protocol
	KeyGenerationType string `json:"keyGenerationType"`
	KeyType KeyType `json:"keyType"`
	// Indicate if the certificate sends a notification on issuance
	IssuanceNotification bool `json:"issuanceNotification"`
	// Indicate if the certificate sends a notification on revocation
	RevocationNotification bool `json:"revocationNotification"`
	// Indicate if the certificate requires an authorization on issuance/revocation
	Authorization bool `json:"authorization"`
	// Indicate if the certificate sends a notification for renewal
	RenewalRule bool `json:"renewalRule"`
	// Indicate if the certificate is published to the public LDAP after issuance
	PublishCertificate bool `json:"publishCertificate"`
	// Indicate if the certificate publication can be overridden
	ClientPublishCertificateOverride bool `json:"clientPublishCertificateOverride"`
	// Indicate the default value if publication override is enabled
	ClientPublishCertificateOverrideDefault bool `json:"clientPublishCertificateOverrideDefault"`
	// Indicate if the certificate product has an expiration date
	ExpirationDate NullableString `json:"expirationDate,omitempty"`
	// When enabled, additional certificate issuance notification recipients can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalIssuanceNotificationRecipients bool `json:"allowAdditionalIssuanceNotificationRecipients"`
	// When enabled, additional certificate revocation notification recipients can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalRevocationNotificationRecipients bool `json:"allowAdditionalRevocationNotificationRecipients"`
	// When enabled, additional certificate renewal notification recipients can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalRenewalNotificationRecipients bool `json:"allowAdditionalRenewalNotificationRecipients"`
	// When enabled, additional authorization notification recipients can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalAuthorizationNotificationRecipients bool `json:"allowAdditionalAuthorizationNotificationRecipients"`
	// When enabled, additional authorization notification recipients (for accepted requests) can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalAuthorizationAcceptedNotificationRecipients bool `json:"allowAdditionalAuthorizationAcceptedNotificationRecipients"`
	// When enabled, additional authorization notification recipients (for rejected requests) can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalAuthorizationRejectedNotificationRecipients bool `json:"allowAdditionalAuthorizationRejectedNotificationRecipients"`
	// When enabled, indicates CAB DNS or HTTP domain validation is required.
	IsCABDNSValidationRequired bool `json:"isCABDNSValidationRequired"`
	// When enabled, indicates that additional notification recipients can be added to the certificate order. Additional recipients are skipped when disabled.
	AllowAdditionalCABDNSNotificationRecipients bool `json:"allowAdditionalCABDNSNotificationRecipients"`
	// When enabled, indicates CAB DNS via constructed email link to domain owner is required.
	IsCABDNSEmailLinkValidationRequired bool `json:"isCABDNSEmailLinkValidationRequired"`
	// When enabled, indicates that the recipient must validate the email box via a link.
	IsEmailBoxValidationRequired bool `json:"isEmailBoxValidationRequired"`
	// When enabled, registration documents must provided with the certificate order. Documents are skipped when disabled.
	RequiresRegistrationDocuments bool `json:"requiresRegistrationDocuments"`
	// When enabled, registration documents must provided when submitting the certificate order. When disabled, documents can be added to the certificate order at a later time via the RA UI.
	RequiresRegistrationDocumentsOnRegister bool `json:"requiresRegistrationDocumentsOnRegister"`
	// PDF registration document are allowed.
	AllowRegistrationDocumentsPDF bool `json:"allowRegistrationDocumentsPDF"`
	// JPG/PNG registration images are allowed.
	AllowRegistrationDocumentsJPG bool `json:"allowRegistrationDocumentsJPG"`
	// Indicates if a revocation code is issued for the recipient (link to self service revocation).
	IsGenerateRevocationCode bool `json:"isGenerateRevocationCode"`
	ProductValidity *ProductValidity `json:"productValidity,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(uuid string, keyGenerationType string, keyType KeyType, issuanceNotification bool, revocationNotification bool, authorization bool, renewalRule bool, publishCertificate bool, clientPublishCertificateOverride bool, clientPublishCertificateOverrideDefault bool, allowAdditionalIssuanceNotificationRecipients bool, allowAdditionalRevocationNotificationRecipients bool, allowAdditionalRenewalNotificationRecipients bool, allowAdditionalAuthorizationNotificationRecipients bool, allowAdditionalAuthorizationAcceptedNotificationRecipients bool, allowAdditionalAuthorizationRejectedNotificationRecipients bool, isCABDNSValidationRequired bool, allowAdditionalCABDNSNotificationRecipients bool, isCABDNSEmailLinkValidationRequired bool, isEmailBoxValidationRequired bool, requiresRegistrationDocuments bool, requiresRegistrationDocumentsOnRegister bool, allowRegistrationDocumentsPDF bool, allowRegistrationDocumentsJPG bool, isGenerateRevocationCode bool) *Product {
	this := Product{}
	this.Uuid = uuid
	this.KeyGenerationType = keyGenerationType
	this.KeyType = keyType
	this.IssuanceNotification = issuanceNotification
	this.RevocationNotification = revocationNotification
	this.Authorization = authorization
	this.RenewalRule = renewalRule
	this.PublishCertificate = publishCertificate
	this.ClientPublishCertificateOverride = clientPublishCertificateOverride
	this.ClientPublishCertificateOverrideDefault = clientPublishCertificateOverrideDefault
	this.AllowAdditionalIssuanceNotificationRecipients = allowAdditionalIssuanceNotificationRecipients
	this.AllowAdditionalRevocationNotificationRecipients = allowAdditionalRevocationNotificationRecipients
	this.AllowAdditionalRenewalNotificationRecipients = allowAdditionalRenewalNotificationRecipients
	this.AllowAdditionalAuthorizationNotificationRecipients = allowAdditionalAuthorizationNotificationRecipients
	this.AllowAdditionalAuthorizationAcceptedNotificationRecipients = allowAdditionalAuthorizationAcceptedNotificationRecipients
	this.AllowAdditionalAuthorizationRejectedNotificationRecipients = allowAdditionalAuthorizationRejectedNotificationRecipients
	this.IsCABDNSValidationRequired = isCABDNSValidationRequired
	this.AllowAdditionalCABDNSNotificationRecipients = allowAdditionalCABDNSNotificationRecipients
	this.IsCABDNSEmailLinkValidationRequired = isCABDNSEmailLinkValidationRequired
	this.IsEmailBoxValidationRequired = isEmailBoxValidationRequired
	this.RequiresRegistrationDocuments = requiresRegistrationDocuments
	this.RequiresRegistrationDocumentsOnRegister = requiresRegistrationDocumentsOnRegister
	this.AllowRegistrationDocumentsPDF = allowRegistrationDocumentsPDF
	this.AllowRegistrationDocumentsJPG = allowRegistrationDocumentsJPG
	this.IsGenerateRevocationCode = isGenerateRevocationCode
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *Product) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Product) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Product) SetUuid(v string) {
	o.Uuid = v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *Product) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *Product) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *Product) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetProductDescription() string {
	if o == nil || o.ProductDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProductDescription.Get()
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetProductDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductDescription.Get(), o.ProductDescription.IsSet()
}

// HasProductDescription returns a boolean if a field has been set.
func (o *Product) HasProductDescription() bool {
	if o != nil && o.ProductDescription.IsSet() {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given NullableString and assigns it to the ProductDescription field.
func (o *Product) SetProductDescription(v string) {
	o.ProductDescription.Set(&v)
}
// SetProductDescriptionNil sets the value for ProductDescription to be an explicit nil
func (o *Product) SetProductDescriptionNil() {
	o.ProductDescription.Set(nil)
}

// UnsetProductDescription ensures that no value is present for ProductDescription, not even an explicit nil
func (o *Product) UnsetProductDescription() {
	o.ProductDescription.Unset()
}

// GetKeyGenerationType returns the KeyGenerationType field value
func (o *Product) GetKeyGenerationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyGenerationType
}

// GetKeyGenerationTypeOk returns a tuple with the KeyGenerationType field value
// and a boolean to check if the value has been set.
func (o *Product) GetKeyGenerationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyGenerationType, true
}

// SetKeyGenerationType sets field value
func (o *Product) SetKeyGenerationType(v string) {
	o.KeyGenerationType = v
}

// GetKeyType returns the KeyType field value
func (o *Product) GetKeyType() KeyType {
	if o == nil {
		var ret KeyType
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *Product) GetKeyTypeOk() (*KeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *Product) SetKeyType(v KeyType) {
	o.KeyType = v
}

// GetIssuanceNotification returns the IssuanceNotification field value
func (o *Product) GetIssuanceNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IssuanceNotification
}

// GetIssuanceNotificationOk returns a tuple with the IssuanceNotification field value
// and a boolean to check if the value has been set.
func (o *Product) GetIssuanceNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuanceNotification, true
}

// SetIssuanceNotification sets field value
func (o *Product) SetIssuanceNotification(v bool) {
	o.IssuanceNotification = v
}

// GetRevocationNotification returns the RevocationNotification field value
func (o *Product) GetRevocationNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RevocationNotification
}

// GetRevocationNotificationOk returns a tuple with the RevocationNotification field value
// and a boolean to check if the value has been set.
func (o *Product) GetRevocationNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevocationNotification, true
}

// SetRevocationNotification sets field value
func (o *Product) SetRevocationNotification(v bool) {
	o.RevocationNotification = v
}

// GetAuthorization returns the Authorization field value
func (o *Product) GetAuthorization() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value
// and a boolean to check if the value has been set.
func (o *Product) GetAuthorizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorization, true
}

// SetAuthorization sets field value
func (o *Product) SetAuthorization(v bool) {
	o.Authorization = v
}

// GetRenewalRule returns the RenewalRule field value
func (o *Product) GetRenewalRule() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RenewalRule
}

// GetRenewalRuleOk returns a tuple with the RenewalRule field value
// and a boolean to check if the value has been set.
func (o *Product) GetRenewalRuleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RenewalRule, true
}

// SetRenewalRule sets field value
func (o *Product) SetRenewalRule(v bool) {
	o.RenewalRule = v
}

// GetPublishCertificate returns the PublishCertificate field value
func (o *Product) GetPublishCertificate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublishCertificate
}

// GetPublishCertificateOk returns a tuple with the PublishCertificate field value
// and a boolean to check if the value has been set.
func (o *Product) GetPublishCertificateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishCertificate, true
}

// SetPublishCertificate sets field value
func (o *Product) SetPublishCertificate(v bool) {
	o.PublishCertificate = v
}

// GetClientPublishCertificateOverride returns the ClientPublishCertificateOverride field value
func (o *Product) GetClientPublishCertificateOverride() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClientPublishCertificateOverride
}

// GetClientPublishCertificateOverrideOk returns a tuple with the ClientPublishCertificateOverride field value
// and a boolean to check if the value has been set.
func (o *Product) GetClientPublishCertificateOverrideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientPublishCertificateOverride, true
}

// SetClientPublishCertificateOverride sets field value
func (o *Product) SetClientPublishCertificateOverride(v bool) {
	o.ClientPublishCertificateOverride = v
}

// GetClientPublishCertificateOverrideDefault returns the ClientPublishCertificateOverrideDefault field value
func (o *Product) GetClientPublishCertificateOverrideDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClientPublishCertificateOverrideDefault
}

// GetClientPublishCertificateOverrideDefaultOk returns a tuple with the ClientPublishCertificateOverrideDefault field value
// and a boolean to check if the value has been set.
func (o *Product) GetClientPublishCertificateOverrideDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientPublishCertificateOverrideDefault, true
}

// SetClientPublishCertificateOverrideDefault sets field value
func (o *Product) SetClientPublishCertificateOverrideDefault(v bool) {
	o.ClientPublishCertificateOverrideDefault = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetExpirationDate() string {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *Product) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableString and assigns it to the ExpirationDate field.
func (o *Product) SetExpirationDate(v string) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *Product) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *Product) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetAllowAdditionalIssuanceNotificationRecipients returns the AllowAdditionalIssuanceNotificationRecipients field value
func (o *Product) GetAllowAdditionalIssuanceNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalIssuanceNotificationRecipients
}

// GetAllowAdditionalIssuanceNotificationRecipientsOk returns a tuple with the AllowAdditionalIssuanceNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalIssuanceNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalIssuanceNotificationRecipients, true
}

// SetAllowAdditionalIssuanceNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalIssuanceNotificationRecipients(v bool) {
	o.AllowAdditionalIssuanceNotificationRecipients = v
}

// GetAllowAdditionalRevocationNotificationRecipients returns the AllowAdditionalRevocationNotificationRecipients field value
func (o *Product) GetAllowAdditionalRevocationNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalRevocationNotificationRecipients
}

// GetAllowAdditionalRevocationNotificationRecipientsOk returns a tuple with the AllowAdditionalRevocationNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalRevocationNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalRevocationNotificationRecipients, true
}

// SetAllowAdditionalRevocationNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalRevocationNotificationRecipients(v bool) {
	o.AllowAdditionalRevocationNotificationRecipients = v
}

// GetAllowAdditionalRenewalNotificationRecipients returns the AllowAdditionalRenewalNotificationRecipients field value
func (o *Product) GetAllowAdditionalRenewalNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalRenewalNotificationRecipients
}

// GetAllowAdditionalRenewalNotificationRecipientsOk returns a tuple with the AllowAdditionalRenewalNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalRenewalNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalRenewalNotificationRecipients, true
}

// SetAllowAdditionalRenewalNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalRenewalNotificationRecipients(v bool) {
	o.AllowAdditionalRenewalNotificationRecipients = v
}

// GetAllowAdditionalAuthorizationNotificationRecipients returns the AllowAdditionalAuthorizationNotificationRecipients field value
func (o *Product) GetAllowAdditionalAuthorizationNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalAuthorizationNotificationRecipients
}

// GetAllowAdditionalAuthorizationNotificationRecipientsOk returns a tuple with the AllowAdditionalAuthorizationNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalAuthorizationNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalAuthorizationNotificationRecipients, true
}

// SetAllowAdditionalAuthorizationNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalAuthorizationNotificationRecipients(v bool) {
	o.AllowAdditionalAuthorizationNotificationRecipients = v
}

// GetAllowAdditionalAuthorizationAcceptedNotificationRecipients returns the AllowAdditionalAuthorizationAcceptedNotificationRecipients field value
func (o *Product) GetAllowAdditionalAuthorizationAcceptedNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalAuthorizationAcceptedNotificationRecipients
}

// GetAllowAdditionalAuthorizationAcceptedNotificationRecipientsOk returns a tuple with the AllowAdditionalAuthorizationAcceptedNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalAuthorizationAcceptedNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalAuthorizationAcceptedNotificationRecipients, true
}

// SetAllowAdditionalAuthorizationAcceptedNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalAuthorizationAcceptedNotificationRecipients(v bool) {
	o.AllowAdditionalAuthorizationAcceptedNotificationRecipients = v
}

// GetAllowAdditionalAuthorizationRejectedNotificationRecipients returns the AllowAdditionalAuthorizationRejectedNotificationRecipients field value
func (o *Product) GetAllowAdditionalAuthorizationRejectedNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalAuthorizationRejectedNotificationRecipients
}

// GetAllowAdditionalAuthorizationRejectedNotificationRecipientsOk returns a tuple with the AllowAdditionalAuthorizationRejectedNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalAuthorizationRejectedNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalAuthorizationRejectedNotificationRecipients, true
}

// SetAllowAdditionalAuthorizationRejectedNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalAuthorizationRejectedNotificationRecipients(v bool) {
	o.AllowAdditionalAuthorizationRejectedNotificationRecipients = v
}

// GetIsCABDNSValidationRequired returns the IsCABDNSValidationRequired field value
func (o *Product) GetIsCABDNSValidationRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCABDNSValidationRequired
}

// GetIsCABDNSValidationRequiredOk returns a tuple with the IsCABDNSValidationRequired field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsCABDNSValidationRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCABDNSValidationRequired, true
}

// SetIsCABDNSValidationRequired sets field value
func (o *Product) SetIsCABDNSValidationRequired(v bool) {
	o.IsCABDNSValidationRequired = v
}

// GetAllowAdditionalCABDNSNotificationRecipients returns the AllowAdditionalCABDNSNotificationRecipients field value
func (o *Product) GetAllowAdditionalCABDNSNotificationRecipients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalCABDNSNotificationRecipients
}

// GetAllowAdditionalCABDNSNotificationRecipientsOk returns a tuple with the AllowAdditionalCABDNSNotificationRecipients field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowAdditionalCABDNSNotificationRecipientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalCABDNSNotificationRecipients, true
}

// SetAllowAdditionalCABDNSNotificationRecipients sets field value
func (o *Product) SetAllowAdditionalCABDNSNotificationRecipients(v bool) {
	o.AllowAdditionalCABDNSNotificationRecipients = v
}

// GetIsCABDNSEmailLinkValidationRequired returns the IsCABDNSEmailLinkValidationRequired field value
func (o *Product) GetIsCABDNSEmailLinkValidationRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCABDNSEmailLinkValidationRequired
}

// GetIsCABDNSEmailLinkValidationRequiredOk returns a tuple with the IsCABDNSEmailLinkValidationRequired field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsCABDNSEmailLinkValidationRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCABDNSEmailLinkValidationRequired, true
}

// SetIsCABDNSEmailLinkValidationRequired sets field value
func (o *Product) SetIsCABDNSEmailLinkValidationRequired(v bool) {
	o.IsCABDNSEmailLinkValidationRequired = v
}

// GetIsEmailBoxValidationRequired returns the IsEmailBoxValidationRequired field value
func (o *Product) GetIsEmailBoxValidationRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEmailBoxValidationRequired
}

// GetIsEmailBoxValidationRequiredOk returns a tuple with the IsEmailBoxValidationRequired field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsEmailBoxValidationRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEmailBoxValidationRequired, true
}

// SetIsEmailBoxValidationRequired sets field value
func (o *Product) SetIsEmailBoxValidationRequired(v bool) {
	o.IsEmailBoxValidationRequired = v
}

// GetRequiresRegistrationDocuments returns the RequiresRegistrationDocuments field value
func (o *Product) GetRequiresRegistrationDocuments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequiresRegistrationDocuments
}

// GetRequiresRegistrationDocumentsOk returns a tuple with the RequiresRegistrationDocuments field value
// and a boolean to check if the value has been set.
func (o *Product) GetRequiresRegistrationDocumentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiresRegistrationDocuments, true
}

// SetRequiresRegistrationDocuments sets field value
func (o *Product) SetRequiresRegistrationDocuments(v bool) {
	o.RequiresRegistrationDocuments = v
}

// GetRequiresRegistrationDocumentsOnRegister returns the RequiresRegistrationDocumentsOnRegister field value
func (o *Product) GetRequiresRegistrationDocumentsOnRegister() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequiresRegistrationDocumentsOnRegister
}

// GetRequiresRegistrationDocumentsOnRegisterOk returns a tuple with the RequiresRegistrationDocumentsOnRegister field value
// and a boolean to check if the value has been set.
func (o *Product) GetRequiresRegistrationDocumentsOnRegisterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiresRegistrationDocumentsOnRegister, true
}

// SetRequiresRegistrationDocumentsOnRegister sets field value
func (o *Product) SetRequiresRegistrationDocumentsOnRegister(v bool) {
	o.RequiresRegistrationDocumentsOnRegister = v
}

// GetAllowRegistrationDocumentsPDF returns the AllowRegistrationDocumentsPDF field value
func (o *Product) GetAllowRegistrationDocumentsPDF() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowRegistrationDocumentsPDF
}

// GetAllowRegistrationDocumentsPDFOk returns a tuple with the AllowRegistrationDocumentsPDF field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowRegistrationDocumentsPDFOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowRegistrationDocumentsPDF, true
}

// SetAllowRegistrationDocumentsPDF sets field value
func (o *Product) SetAllowRegistrationDocumentsPDF(v bool) {
	o.AllowRegistrationDocumentsPDF = v
}

// GetAllowRegistrationDocumentsJPG returns the AllowRegistrationDocumentsJPG field value
func (o *Product) GetAllowRegistrationDocumentsJPG() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowRegistrationDocumentsJPG
}

// GetAllowRegistrationDocumentsJPGOk returns a tuple with the AllowRegistrationDocumentsJPG field value
// and a boolean to check if the value has been set.
func (o *Product) GetAllowRegistrationDocumentsJPGOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowRegistrationDocumentsJPG, true
}

// SetAllowRegistrationDocumentsJPG sets field value
func (o *Product) SetAllowRegistrationDocumentsJPG(v bool) {
	o.AllowRegistrationDocumentsJPG = v
}

// GetIsGenerateRevocationCode returns the IsGenerateRevocationCode field value
func (o *Product) GetIsGenerateRevocationCode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGenerateRevocationCode
}

// GetIsGenerateRevocationCodeOk returns a tuple with the IsGenerateRevocationCode field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsGenerateRevocationCodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGenerateRevocationCode, true
}

// SetIsGenerateRevocationCode sets field value
func (o *Product) SetIsGenerateRevocationCode(v bool) {
	o.IsGenerateRevocationCode = v
}

// GetProductValidity returns the ProductValidity field value if set, zero value otherwise.
func (o *Product) GetProductValidity() ProductValidity {
	if o == nil || o.ProductValidity == nil {
		var ret ProductValidity
		return ret
	}
	return *o.ProductValidity
}

// GetProductValidityOk returns a tuple with the ProductValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductValidityOk() (*ProductValidity, bool) {
	if o == nil || o.ProductValidity == nil {
		return nil, false
	}
	return o.ProductValidity, true
}

// HasProductValidity returns a boolean if a field has been set.
func (o *Product) HasProductValidity() bool {
	if o != nil && o.ProductValidity != nil {
		return true
	}

	return false
}

// SetProductValidity gets a reference to the given ProductValidity and assigns it to the ProductValidity field.
func (o *Product) SetProductValidity(v ProductValidity) {
	o.ProductValidity = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ProductName != nil {
		toSerialize["productName"] = o.ProductName
	}
	if o.ProductDescription.IsSet() {
		toSerialize["productDescription"] = o.ProductDescription.Get()
	}
	if true {
		toSerialize["keyGenerationType"] = o.KeyGenerationType
	}
	if true {
		toSerialize["keyType"] = o.KeyType
	}
	if true {
		toSerialize["issuanceNotification"] = o.IssuanceNotification
	}
	if true {
		toSerialize["revocationNotification"] = o.RevocationNotification
	}
	if true {
		toSerialize["authorization"] = o.Authorization
	}
	if true {
		toSerialize["renewalRule"] = o.RenewalRule
	}
	if true {
		toSerialize["publishCertificate"] = o.PublishCertificate
	}
	if true {
		toSerialize["clientPublishCertificateOverride"] = o.ClientPublishCertificateOverride
	}
	if true {
		toSerialize["clientPublishCertificateOverrideDefault"] = o.ClientPublishCertificateOverrideDefault
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expirationDate"] = o.ExpirationDate.Get()
	}
	if true {
		toSerialize["allowAdditionalIssuanceNotificationRecipients"] = o.AllowAdditionalIssuanceNotificationRecipients
	}
	if true {
		toSerialize["allowAdditionalRevocationNotificationRecipients"] = o.AllowAdditionalRevocationNotificationRecipients
	}
	if true {
		toSerialize["allowAdditionalRenewalNotificationRecipients"] = o.AllowAdditionalRenewalNotificationRecipients
	}
	if true {
		toSerialize["allowAdditionalAuthorizationNotificationRecipients"] = o.AllowAdditionalAuthorizationNotificationRecipients
	}
	if true {
		toSerialize["allowAdditionalAuthorizationAcceptedNotificationRecipients"] = o.AllowAdditionalAuthorizationAcceptedNotificationRecipients
	}
	if true {
		toSerialize["allowAdditionalAuthorizationRejectedNotificationRecipients"] = o.AllowAdditionalAuthorizationRejectedNotificationRecipients
	}
	if true {
		toSerialize["isCABDNSValidationRequired"] = o.IsCABDNSValidationRequired
	}
	if true {
		toSerialize["allowAdditionalCABDNSNotificationRecipients"] = o.AllowAdditionalCABDNSNotificationRecipients
	}
	if true {
		toSerialize["isCABDNSEmailLinkValidationRequired"] = o.IsCABDNSEmailLinkValidationRequired
	}
	if true {
		toSerialize["isEmailBoxValidationRequired"] = o.IsEmailBoxValidationRequired
	}
	if true {
		toSerialize["requiresRegistrationDocuments"] = o.RequiresRegistrationDocuments
	}
	if true {
		toSerialize["requiresRegistrationDocumentsOnRegister"] = o.RequiresRegistrationDocumentsOnRegister
	}
	if true {
		toSerialize["allowRegistrationDocumentsPDF"] = o.AllowRegistrationDocumentsPDF
	}
	if true {
		toSerialize["allowRegistrationDocumentsJPG"] = o.AllowRegistrationDocumentsJPG
	}
	if true {
		toSerialize["isGenerateRevocationCode"] = o.IsGenerateRevocationCode
	}
	if o.ProductValidity != nil {
		toSerialize["productValidity"] = o.ProductValidity
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


