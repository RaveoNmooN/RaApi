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
	"fmt"
)

// AdditionalRecipientType Notification type for additional recipient.   - DNS_OWNER_CHECK_EMAIL_LINK: send notification for DNS validation (email link) to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - DNS_CAB: send notification for CAB DNS change or CAB Agreed upon to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - ISSUANCE: send notification for certificate issuance to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - REVOCATION: send notification for certificate revocation to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - RENEWAL: send notification for certificate renewal to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - RECOVERY: send notification for key recovery to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - ISSUANCE_AUTHORIZATION: send notification for certificate issuance authorization to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - ISSUANCE_AUTHORIZATION_ACCEPTED: send notification for accepted certificate issuance authorization to additional recipient. Requires notification enabled for 'Additional recipient' on notification template.   - ISSUANCE_AUTHORIZATION_REJECTED: send notification for rejected certificate issuance authorization to additional recipient. Requires notification enabled for 'Additional recipient' on notification template. 
type AdditionalRecipientType string

// List of AdditionalRecipientType
const (
	DNS_OWNER_CHECK_EMAIL_LINK AdditionalRecipientType = "DNS_OWNER_CHECK_EMAIL_LINK"
	DNS_CAB AdditionalRecipientType = "DNS_CAB"
	ISSUANCE AdditionalRecipientType = "ISSUANCE"
	REVOCATION AdditionalRecipientType = "REVOCATION"
	RENEWAL AdditionalRecipientType = "RENEWAL"
	RECOVERY AdditionalRecipientType = "RECOVERY"
	ISSUANCE_AUTHORIZATION AdditionalRecipientType = "ISSUANCE_AUTHORIZATION"
	ISSUANCE_AUTHORIZATION_ACCEPTED AdditionalRecipientType = "ISSUANCE_AUTHORIZATION_ACCEPTED"
	ISSUANCE_AUTHORIZATION_REJECTED AdditionalRecipientType = "ISSUANCE_AUTHORIZATION_REJECTED"
)

// All allowed values of AdditionalRecipientType enum
var AllowedAdditionalRecipientTypeEnumValues = []AdditionalRecipientType{
	"DNS_OWNER_CHECK_EMAIL_LINK",
	"DNS_CAB",
	"ISSUANCE",
	"REVOCATION",
	"RENEWAL",
	"RECOVERY",
	"ISSUANCE_AUTHORIZATION",
	"ISSUANCE_AUTHORIZATION_ACCEPTED",
	"ISSUANCE_AUTHORIZATION_REJECTED",
}

func (v *AdditionalRecipientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdditionalRecipientType(value)
	for _, existing := range AllowedAdditionalRecipientTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdditionalRecipientType", value)
}

// NewAdditionalRecipientTypeFromValue returns a pointer to a valid AdditionalRecipientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdditionalRecipientTypeFromValue(v string) (*AdditionalRecipientType, error) {
	ev := AdditionalRecipientType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdditionalRecipientType: valid values are %v", v, AllowedAdditionalRecipientTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdditionalRecipientType) IsValid() bool {
	for _, existing := range AllowedAdditionalRecipientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdditionalRecipientType value
func (v AdditionalRecipientType) Ptr() *AdditionalRecipientType {
	return &v
}

type NullableAdditionalRecipientType struct {
	value *AdditionalRecipientType
	isSet bool
}

func (v NullableAdditionalRecipientType) Get() *AdditionalRecipientType {
	return v.value
}

func (v *NullableAdditionalRecipientType) Set(val *AdditionalRecipientType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalRecipientType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalRecipientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalRecipientType(val *AdditionalRecipientType) *NullableAdditionalRecipientType {
	return &NullableAdditionalRecipientType{value: val, isSet: true}
}

func (v NullableAdditionalRecipientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalRecipientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

