/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Contains Product Offer details
type GetProductOfferResponse struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// Schema location of this resource
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// An agreement represents a contract or arrangement, either written or verbal and sometimes enforceable by law, such as a service level agreement or a customer price agreement. An agreement involves a number of other business entities, such as products, services, and resources and/or their specifications.
	Agreement []AgreementRef `json:"agreement,omitempty"`
	// Complements the description of an element (for instance a product) through video, pictures...
	Attachment []Attachment `json:"attachment,omitempty"`
	// A type of ProductOffering that belongs to a grouping of ProductOfferings made available to the market. It inherits of all attributes of ProductOffering.
	BundledProductOffering []BundledProductOffering `json:"bundledProductOffering,omitempty"`
	// The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other categories and/or product offerings, resource or service candidates.
	Category []CategoryRef `json:"category,omitempty"`
	// The channel defines the channel for selling product offerings.
	Channel []ChannelRef `json:"channel,omitempty"`
	// Unique code which is used to communicate with external systems.
	Code string `json:"code,omitempty"`
	// User Name of the User who has created this entity 
	CreatedBy string `json:"createdBy,omitempty"`
	// DateTime at which user has created this entity 
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// Description of the productOffering
	Description string `json:"description,omitempty"`
	// A String. This is ID/Primary key/Unique ID of the entity which is loaded/inserted/imported in BM Catalog.
	ExternalID string `json:"externalID,omitempty"`
	// Href of the productOffering
	Href string `json:"href,omitempty"`
	// Unique identifier of the productOffering
	Id string `json:"id,omitempty"`
	// isBundle determines whether a productOffering represents a single productOffering (false), or a bundle of productOfferings (true).
	IsBundle bool `json:"isBundle,omitempty"`
	// A flag indicating if this product offer can be sold stand-alone for sale or not. If this flag is false it indicates that the offer can only be sold within a bundle.
	IsSellable bool `json:"isSellable,omitempty"`
	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
	// User Name of the User who has last updated this entity 
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// provides references to the corresponding market segment as target of product offerings. A market segment is grouping of Parties, GeographicAreas, SalesChannels, and so forth.
	MarketSegment []MarketSegmentRef `json:"marketSegment,omitempty"`
	// Name of the productOffering
	Name string `json:"name,omitempty"`
	OfferLabels []OfferLabels `json:"offerLabels,omitempty"`
	// Place defines the places where the products are sold or delivered.
	Place []PlaceRef `json:"place,omitempty"`
	// An amount, usually of money, that is asked for or allowed when a ProductOffering is bought, rented, or leased. The price is valid for a defined period of time and may not represent the actual price paid by a customer.
	ProductOfferingPrice []ProductOfferingPrice `json:"productOfferingPrice,omitempty"`
	// A condition under which a ProductOffering is made available to Customers. For instance, a productOffering can be offered with multiple commitment periods.
	ProductOfferingTerm []ProductOfferingTerm `json:"productOfferingTerm,omitempty"`
	// A Product Specification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
	ProductSpecification *ProductSpecification `json:"productSpecification,omitempty"`
	// A resource candidate is an entity that makes a ResourceSpecification available to a catalog.
	ResourceCandidate *ResourceCandidateRef `json:"resourceCandidate,omitempty"`
	// ServiceCandidate is an entity that makes a ServiceSpecification available to a catalog.
	ServiceCandidate *ServiceCandidateRef `json:"serviceCandidate,omitempty"`
	// A service level agreement (SLA) is a type of agreement that represents a formal negotiated agreement between two parties designed to create a common understanding about products, services, priorities, responsibilities, and so forth. The SLA is a set of appropriate procedures and targets formally or informally agreed between parties in order to achieve and maintain specified Quality of Service.
	ServiceLevelAgreement *SlaRef `json:"serviceLevelAgreement,omitempty"`
	// A use of the SpecificationCharacteristicValue by a ProductOffering to which additional properties (attributes) apply or override the properties of similar properties contained in SpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding specification. The available characteristic values for a SpecificationCharacteristic in a specification can be modified at the ProductOffering level. For example, a characteristic 'Color' might have values White, Blue, Green, and Red. But, the list of values can be restricted to e.g. White and Blue in an associated product offering. It should be noted that the list of values in 'SpecificationCharacteristicValueUse' is a strict subset of the list of values as defined in the corresponding specification characteristics.
	SpecCharValueUse []ConfigurableSpecificationCharacteristicValueUse `json:"specCharValueUse,omitempty"`
	// A string providing a complementary information on the value of the lifecycle status attribute.
	StatusReason string `json:"statusReason,omitempty"`
	// The period for which the productOffering is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// ProductOffering version
	Version string `json:"version,omitempty"`
}