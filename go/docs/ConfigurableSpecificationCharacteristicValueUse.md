# ConfigurableSpecificationCharacteristicValueUse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** |  hyperlink reference to the schema describing this resource. | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class entity name | [optional] [default to null]
**ConfigurationStage** | **string** | A String. This is used to get configurationStage at productOffering level | [optional] [default to null]
**Description** | **string** | A narrative that explains in detail what the productSpecificationCharacteristic is | [optional] [default to null]
**DisplayName** | **string** | A String.  This is used to get displayName at productOffering level | [optional] [default to null]
**DisplaySequence** | **string** | This defines on what order the characteristics will be exposed. | [optional] [default to null]
**IsVisible** | **string** | This defines the characteristics will be exposed or not. | [optional] [default to null]
**MaxCardinality** | **int32** | The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality. | [optional] [default to null]
**MinCardinality** | **int32** | The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality. | [optional] [default to null]
**Name** | **string** | Name of the associated productSpecificationCharacteristic | [default to null]
**ReferenceId** | **string** |  | [optional] [default to null]
**SpecCharacteristicValue** | [**[]SpecificationCharacteristicValue**](SpecificationCharacteristicValue.md) | A number or text that can be assigned to a SpecificationCharacteristic. | [default to null]
**Specification** | [***SpecificationRef**](SpecificationRef.md) | A Specification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role. | [optional] [default to null]
**SpecificationType** | **string** | When sub-classing, this defines the sub-class specification type | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the productSpecificationCharacteristic is valid | [optional] [default to null]
**ValueType** | **string** | A kind of value that the characteristic can take on, such as numeric, text and so forth | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


