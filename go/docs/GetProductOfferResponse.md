# GetProductOfferResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | Schema location of this resource | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class entity name | [optional] [default to null]
**Agreement** | [**[]AgreementRef**](AgreementRef.md) | An agreement represents a contract or arrangement, either written or verbal and sometimes enforceable by law, such as a service level agreement or a customer price agreement. An agreement involves a number of other business entities, such as products, services, and resources and/or their specifications. | [optional] [default to null]
**Attachment** | [**[]Attachment**](Attachment.md) | Complements the description of an element (for instance a product) through video, pictures... | [optional] [default to null]
**BundledProductOffering** | [**[]BundledProductOffering**](BundledProductOffering.md) | A type of ProductOffering that belongs to a grouping of ProductOfferings made available to the market. It inherits of all attributes of ProductOffering. | [optional] [default to null]
**Category** | [**[]CategoryRef**](CategoryRef.md) | The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other categories and/or product offerings, resource or service candidates. | [optional] [default to null]
**Channel** | [**[]ChannelRef**](ChannelRef.md) | The channel defines the channel for selling product offerings. | [optional] [default to null]
**Code** | **string** | Unique code which is used to communicate with external systems. | [optional] [default to null]
**CreatedBy** | **string** | User Name of the User who has created this entity  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | DateTime at which user has created this entity  | [optional] [default to null]
**Description** | **string** | Description of the productOffering | [optional] [default to null]
**ExternalID** | **string** | A String. This is ID/Primary key/Unique ID of the entity which is loaded/inserted/imported in BM Catalog. | [optional] [default to null]
**Href** | **string** | Href of the productOffering | [optional] [default to null]
**Id** | **string** | Unique identifier of the productOffering | [optional] [default to null]
**IsBundle** | **bool** | isBundle determines whether a productOffering represents a single productOffering (false), or a bundle of productOfferings (true). | [optional] [default to null]
**IsSellable** | **bool** | A flag indicating if this product offer can be sold stand-alone for sale or not. If this flag is false it indicates that the offer can only be sold within a bundle. | [optional] [default to null]
**LastUpdate** | [**time.Time**](time.Time.md) | Date and time of the last update | [optional] [default to null]
**LastUpdatedBy** | **string** | User Name of the User who has last updated this entity  | [optional] [default to null]
**LifecycleStatus** | **string** | Used to indicate the current lifecycle status | [optional] [default to null]
**MarketSegment** | [**[]MarketSegmentRef**](MarketSegmentRef.md) | provides references to the corresponding market segment as target of product offerings. A market segment is grouping of Parties, GeographicAreas, SalesChannels, and so forth. | [optional] [default to null]
**Name** | **string** | Name of the productOffering | [optional] [default to null]
**OfferLabels** | [**[]OfferLabels**](OfferLabels.md) |  | [optional] [default to null]
**Place** | [**[]PlaceRef**](PlaceRef.md) | Place defines the places where the products are sold or delivered. | [optional] [default to null]
**ProductOfferingPrice** | [**[]ProductOfferingPrice**](ProductOfferingPrice.md) | An amount, usually of money, that is asked for or allowed when a ProductOffering is bought, rented, or leased. The price is valid for a defined period of time and may not represent the actual price paid by a customer. | [optional] [default to null]
**ProductOfferingTerm** | [**[]ProductOfferingTerm**](ProductOfferingTerm.md) | A condition under which a ProductOffering is made available to Customers. For instance, a productOffering can be offered with multiple commitment periods. | [optional] [default to null]
**ProductSpecification** | [***ProductSpecification**](ProductSpecification.md) | A Product Specification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role. | [optional] [default to null]
**ResourceCandidate** | [***ResourceCandidateRef**](ResourceCandidateRef.md) | A resource candidate is an entity that makes a ResourceSpecification available to a catalog. | [optional] [default to null]
**ServiceCandidate** | [***ServiceCandidateRef**](ServiceCandidateRef.md) | ServiceCandidate is an entity that makes a ServiceSpecification available to a catalog. | [optional] [default to null]
**ServiceLevelAgreement** | [***SlaRef**](SLARef.md) | A service level agreement (SLA) is a type of agreement that represents a formal negotiated agreement between two parties designed to create a common understanding about products, services, priorities, responsibilities, and so forth. The SLA is a set of appropriate procedures and targets formally or informally agreed between parties in order to achieve and maintain specified Quality of Service. | [optional] [default to null]
**SpecCharValueUse** | [**[]ConfigurableSpecificationCharacteristicValueUse**](ConfigurableSpecificationCharacteristicValueUse.md) | A use of the SpecificationCharacteristicValue by a ProductOffering to which additional properties (attributes) apply or override the properties of similar properties contained in SpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding specification. The available characteristic values for a SpecificationCharacteristic in a specification can be modified at the ProductOffering level. For example, a characteristic &#39;Color&#39; might have values White, Blue, Green, and Red. But, the list of values can be restricted to e.g. White and Blue in an associated product offering. It should be noted that the list of values in &#39;SpecificationCharacteristicValueUse&#39; is a strict subset of the list of values as defined in the corresponding specification characteristics. | [optional] [default to null]
**StatusReason** | **string** | A string providing a complementary information on the value of the lifecycle status attribute. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the productOffering is valid | [optional] [default to null]
**Version** | **string** | ProductOffering version | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


