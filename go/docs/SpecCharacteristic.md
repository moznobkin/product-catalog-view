# SpecCharacteristic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class entity name | [optional] [default to null]
**CharacteristicType** | **string** |  | [optional] [default to null]
**Code** | **string** | Code of Characteristics for external use. It will be the system generated code and not required in input | [optional] [default to null]
**Configurable** | **bool** | If true, the Boolean indicates that the SpecCharacteristic is configurable | [optional] [default to null]
**ConfigurationStage** | **string** | A String. This is the place where we will define the place where characteristic will be defined. | [optional] [default to null]
**CreatedBy** | **string** |  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Description** | **string** | A narrative that explains in detail what the SpecCharacteristic is | [optional] [default to null]
**DisplayName** | **string** | A String. Here we can define the display name of the Characteristic. | [optional] [default to null]
**DisplaySequence** | **int32** |  | [optional] [default to null]
**Extensible** | **bool** | An indicator that specifies that the values for the characteristic can be extended by adding new values when instantiating a characteristic for an Entity. | [optional] [default to null]
**ExternalID** | **string** |  | [optional] [default to null]
**Href** | **string** |  | [optional] [default to null]
**Id** | **string** | Unique identifier of Characteristics for external use. It will be the system generated code and not required in input. | [optional] [default to null]
**IsUnique** | **bool** | An indicator that specifies if a value is unique for the specification. Possible values are; \&quot;unique while value is in effect\&quot; and \&quot;unique whether value is in effect or not\&quot; | [optional] [default to null]
**IsVisible** | **bool** |  | [optional] [default to null]
**LastUpdate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastUpdatedBy** | **string** |  | [optional] [default to null]
**MaxCardinality** | **int32** | The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality. | [optional] [default to null]
**MinCardinality** | **int32** | The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality. | [optional] [default to null]
**Name** | **string** | A word, term, or phrase by which this characteristic specification is known and distinguished from other characteristic specifications. | [optional] [default to null]
**ReferenceId** | **string** |  | [optional] [default to null]
**Regex** | **string** | A rule or principle represented in regular expression used to derive the value of a characteristic value. | [optional] [default to null]
**SpecCharRelationship** | [**[]SpecCharRelationship**](SpecCharRelationship.md) | A list of spec char relationships (SpecCharRelationship [*]). An aggregation, migration, substitution, dependency or exclusivity relationship between/among Specification Characteristics. | [optional] [default to null]
**SpecCharacteristicValue** | [**[]SpecCharacteristicValue**](SpecCharacteristicValue.md) | A list of service spec characteristic values (SpecCharacteristicValue [*]). A SpecCharacteristicValue object is used to define a set of attributes, each of which can be assigned to a corresponding set of attributes in a SpecCharacteristic object. The values of the attributes in the SpecCharacteristicValue object describe the values of the attributes that a corresponding SpecCharacteristic object can take on. | [optional] [default to null]
**ValueType** | **string** | A kind of value that the characteristic can take on, such as numeric, text and so forth | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


