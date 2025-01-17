/*
 * Neon CRM API Reference
 *
 * Welcome to the Neon CRM API v2. This version is a complete rebuild of the existing [API v1](https://developer.neoncrm.com/api) and adopts many modern RESTful API conventions. API v2 also introduces new methods for working with custom fields and events, in addition to updating and extending the functionality of all the existing methods from v1.  API v2 has two possible base URLs, depending on the type of Neon CRM instance you are using: - **Production & Sandbox Instances:** [https://api.neoncrm.com/v2](https://api.neoncrm.com/v2) - **Trial Instances:** [https://trial.neoncrm.com/v2](https://trial.neoncrm.com/v2)  ### Authorization API v2 authorizes all requests using [HTTP basic authentication](https://en.wikipedia.org/wiki/Basic_access_authentication). Use the values below for the username and password pair in the `Authorization` header of your request. - *Username* – The organization ID of your Neon CRM instance - *Password* – An API key associated with a user in the instance  Learn [how to obtain an org ID and API key](https://developer.neoncrm.com/authentication/) from your Neon CRM instance.  ### Versioning This API is versioned to allow for ongoing fixes and updates to the API that may include backwards-incompatible changes. Specify the API version by including the header `NEON-API-VERSION: 2.X` with every request. The current version is listed at the top of these docs.  Learn more about [versioning in API v2](https://developer.neoncrm.com/versioning/).  ### Legacy Events vs New Events Currently, only [legacy events](https://support.neonone.com/hc/en-us/articles/18441378851981-Legacy-Events-Vs-Events) are supported within Neon CRM's API. [Follow this page for updates.](https://developer.neoncrm.com/updates/)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package neon

import (
	"time"
)

type SubMembership struct {
	Id                   string      `json:"id,omitempty"`
	ParentId             string      `json:"parentId,omitempty"`
	AccountId            string      `json:"accountId,omitempty"`
	MembershipLevel      *IdNamePair `json:"membershipLevel,omitempty"`
	MembershipTerm       *IdNamePair `json:"membershipTerm,omitempty"`
	AutoRenewal          bool        `json:"autoRenewal,omitempty"`
	Source               *IdNamePair `json:"source,omitempty"`
	ChangeType           string      `json:"changeType,omitempty"`
	TermUnit             string      `json:"termUnit,omitempty"`
	TermDuration         int32       `json:"termDuration,omitempty"`
	EnrollType           string      `json:"enrollType,omitempty"`
	TransactionDate      time.Time   `json:"transactionDate,omitempty"`
	TermStartDate        time.Time   `json:"termStartDate,omitempty"`
	TermEndDate          time.Time   `json:"termEndDate,omitempty"`
	Fee                  float64     `json:"fee,omitempty"`
	CouponCode           string      `json:"couponCode,omitempty"`
	SendAcknowledgeEmail bool        `json:"sendAcknowledgeEmail,omitempty"`
	Status               string      `json:"status,omitempty"`
	// 1: Complimentary memberships because of donations made.<br>2: Complimentary memberships through the combination of past membership fees.<br>3: Complimentary memberships through the combination of donations made and previous membership fees.
	Complimentary               int32             `json:"complimentary,omitempty"`
	MembershipCustomFields      []CustomField     `json:"membershipCustomFields,omitempty"`
	CraInfo                     *CraInfo          `json:"craInfo,omitempty"`
	TaxDeductibleInfo           *TaxDeducibleInfo `json:"taxDeductibleInfo,omitempty"`
	Timestamps                  *Timestamps       `json:"timestamps,omitempty"`
	Origin                      *Origin           `json:"origin,omitempty"`
	SendAutoRenewalEnabledEmail bool              `json:"sendAutoRenewalEnabledEmail,omitempty"`
}
