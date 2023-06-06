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

// DNS struct for DNS
type DNS struct {
	// requested domain name (gets converted to punycode if required)
	Dns string `json:"dns"`
	// Append base domain to wildcard DNS. Adds sample.org when *.sample.org is requested. Applies only to certificate policies which allow wildcard issuance. This flag has no effect if the certificate policy does not include the wildcard option. 
	IncludeBaseDomainForWildcard *bool `json:"includeBaseDomainForWildcard,omitempty"`
	// Append www to requested DNS.  Adds www.sample.org when sample.org is requested. When enabled, www is prefixed to all requested DNS. 
	IncludeWWWDomain *bool `json:"includeWWWDomain,omitempty"`
}

// NewDNS instantiates a new DNS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNS(dns string) *DNS {
	this := DNS{}
	this.Dns = dns
	return &this
}

// NewDNSWithDefaults instantiates a new DNS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSWithDefaults() *DNS {
	this := DNS{}
	return &this
}

// GetDns returns the Dns field value
func (o *DNS) GetDns() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dns
}

// GetDnsOk returns a tuple with the Dns field value
// and a boolean to check if the value has been set.
func (o *DNS) GetDnsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dns, true
}

// SetDns sets field value
func (o *DNS) SetDns(v string) {
	o.Dns = v
}

// GetIncludeBaseDomainForWildcard returns the IncludeBaseDomainForWildcard field value if set, zero value otherwise.
func (o *DNS) GetIncludeBaseDomainForWildcard() bool {
	if o == nil || o.IncludeBaseDomainForWildcard == nil {
		var ret bool
		return ret
	}
	return *o.IncludeBaseDomainForWildcard
}

// GetIncludeBaseDomainForWildcardOk returns a tuple with the IncludeBaseDomainForWildcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNS) GetIncludeBaseDomainForWildcardOk() (*bool, bool) {
	if o == nil || o.IncludeBaseDomainForWildcard == nil {
		return nil, false
	}
	return o.IncludeBaseDomainForWildcard, true
}

// HasIncludeBaseDomainForWildcard returns a boolean if a field has been set.
func (o *DNS) HasIncludeBaseDomainForWildcard() bool {
	if o != nil && o.IncludeBaseDomainForWildcard != nil {
		return true
	}

	return false
}

// SetIncludeBaseDomainForWildcard gets a reference to the given bool and assigns it to the IncludeBaseDomainForWildcard field.
func (o *DNS) SetIncludeBaseDomainForWildcard(v bool) {
	o.IncludeBaseDomainForWildcard = &v
}

// GetIncludeWWWDomain returns the IncludeWWWDomain field value if set, zero value otherwise.
func (o *DNS) GetIncludeWWWDomain() bool {
	if o == nil || o.IncludeWWWDomain == nil {
		var ret bool
		return ret
	}
	return *o.IncludeWWWDomain
}

// GetIncludeWWWDomainOk returns a tuple with the IncludeWWWDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNS) GetIncludeWWWDomainOk() (*bool, bool) {
	if o == nil || o.IncludeWWWDomain == nil {
		return nil, false
	}
	return o.IncludeWWWDomain, true
}

// HasIncludeWWWDomain returns a boolean if a field has been set.
func (o *DNS) HasIncludeWWWDomain() bool {
	if o != nil && o.IncludeWWWDomain != nil {
		return true
	}

	return false
}

// SetIncludeWWWDomain gets a reference to the given bool and assigns it to the IncludeWWWDomain field.
func (o *DNS) SetIncludeWWWDomain(v bool) {
	o.IncludeWWWDomain = &v
}

func (o DNS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dns"] = o.Dns
	}
	if o.IncludeBaseDomainForWildcard != nil {
		toSerialize["includeBaseDomainForWildcard"] = o.IncludeBaseDomainForWildcard
	}
	if o.IncludeWWWDomain != nil {
		toSerialize["includeWWWDomain"] = o.IncludeWWWDomain
	}
	return json.Marshal(toSerialize)
}

type NullableDNS struct {
	value *DNS
	isSet bool
}

func (v NullableDNS) Get() *DNS {
	return v.value
}

func (v *NullableDNS) Set(val *DNS) {
	v.value = val
	v.isSet = true
}

func (v NullableDNS) IsSet() bool {
	return v.isSet
}

func (v *NullableDNS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNS(val *DNS) *NullableDNS {
	return &NullableDNS{value: val, isSet: true}
}

func (v NullableDNS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


