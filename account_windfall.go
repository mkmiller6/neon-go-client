/*
 * Neon CRM API Reference
 *
 * Welcome to the Neon CRM API v2. This version is a complete rebuild of the existing [API v1](https://developer.neoncrm.com/api) and adopts many modern RESTful API conventions. API v2 also introduces new methods for working with custom fields and events, in addition to updating and extending the functionality of all the existing methods from v1.  API v2 has two possible base URLs, depending on the type of Neon CRM instance you are using: - **Production & Sandbox Instances:** [https://api.neoncrm.com/v2](https://api.neoncrm.com/v2) - **Trial Instances:** [https://trial.neoncrm.com/v2](https://trial.neoncrm.com/v2)  ### Authorization API v2 authorizes all requests using [HTTP basic authentication](https://en.wikipedia.org/wiki/Basic_access_authentication). Use the values below for the username and password pair in the `Authorization` header of your request. - *Username* – The organization ID of your Neon CRM instance - *Password* – An API key associated with a user in the instance  Learn [how to obtain an org ID and API key](https://developer.neoncrm.com/authentication/) from your Neon CRM instance.  ### Versioning This API is versioned to allow for ongoing fixes and updates to the API that may include backwards-incompatible changes. Specify the API version by including the header `NEON-API-VERSION: 2.X` with every request. The current version is listed at the top of these docs.  Learn more about [versioning in API v2](https://developer.neoncrm.com/versioning/).  ### Legacy Events vs New Events Currently, only [legacy events](https://support.neonone.com/hc/en-us/articles/18441378851981-Legacy-Events-Vs-Events) are supported within Neon CRM's API. [Follow this page for updates.](https://developer.neoncrm.com/updates/)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package neon

type AccountWindfall struct {
	Id                      string   `json:"id,omitempty"`
	WindfallId              string   `json:"windfall_id,omitempty"`
	NetWorth                float64  `json:"net_worth,omitempty"`
	LowConfidenceNetWorth   float64  `json:"low_confidence_net_worth,omitempty"`
	HighConfidenceNetWorth  float64  `json:"high_confidence_net_worth,omitempty"`
	NetWorthLastCalculated  string   `json:"net_worth_last_calculated,omitempty"`
	RecentMover             bool     `json:"recent_mover,omitempty"`
	RecentlyDivorced        bool     `json:"recently_divorced,omitempty"`
	RecentDeathInFamily     bool     `json:"recent_death_in_family,omitempty"`
	PoliticalAffiliation    string   `json:"political_affiliation,omitempty"`
	PoliticalDonor          bool     `json:"political_donor,omitempty"`
	BoatOwner               bool     `json:"boat_owner,omitempty"`
	PlaneOwner              bool     `json:"plane_owner,omitempty"`
	SmallBusinessOwner      bool     `json:"small_business_owner,omitempty"`
	PhilanthropicGiver      bool     `json:"philanthropic_giver,omitempty"`
	PhilanthropicCause      []string `json:"philanthropic_cause,omitempty"`
	PhilanthropicFocus      []string `json:"philanthropic_focus,omitempty"`
	MultiPropertyOwner      bool     `json:"multi-property_owner,omitempty"`
	FoundationAssociation   bool     `json:"foundation_association,omitempty"`
	FoundationOfficer       bool     `json:"foundation_officer,omitempty"`
	TrustAssociation        bool     `json:"trust_association,omitempty"`
	NonprofitBoardMember    bool     `json:"nonprofit_board_member,omitempty"`
	TopPoliticalDonor       bool     `json:"top_political_donor,omitempty"`
	MatchConfidence         float64  `json:"match_confidence,omitempty"`
	SecMoneyInMotion        bool     `json:"sec_money_in_motion,omitempty"`
	LiquidityTrigger        bool     `json:"liquidity_trigger,omitempty"`
	LuxuryCarOwner          bool     `json:"luxury_car_owner,omitempty"`
	ImportedCarOwner        bool     `json:"imported_car_owner,omitempty"`
	RecentMortgage          bool     `json:"recent_mortgage,omitempty"`
	PrimaryAddressIncorrect bool     `json:"primary_address_incorrect,omitempty"`
}
