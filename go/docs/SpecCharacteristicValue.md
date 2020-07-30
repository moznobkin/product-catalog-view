# SpecCharacteristicValue

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseType** | **string** | A string. The (immediate) base class type of this REST resource | [optional] [default to null]
**Type_** | **string** | A string class type of target specification | [optional] [default to null]
**IsDefault** | **bool** | Indicates if the value is the default value for a characteristic | [optional] [default to null]
**RangeInterval** | **string** | An indicator that specifies the inclusion or exclusion of the valueFrom and valueTo attributes. If applicable, possible values are \&quot;open\&quot;, \&quot;closed\&quot;, \&quot;closedBottom\&quot; and \&quot;closedTop\&quot;. | [optional] [default to null]
**Regex** | **string** | A regular expression constraint for given value | [optional] [default to null]
**UnitOfMeasure** | **string** | A length, surface, volume, dry measure, liquid measure, money, weight, time, and the like. In general, a determinate quantity or magnitude of the kind designated, taken as a standard of comparison for others of the same kind, in assigning to them numerical values, as 1 foot, 1 yard, 1 mile, 1 square foot. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period of time for which a value is applicable | [optional] [default to null]
**Value** | **string** | A discrete value that the characteristic can take on, or the actual value of the characteristic | [optional] [default to null]
**ValueFrom** | **int32** | The low range value that a characteristic can take on | [optional] [default to null]
**ValueTo** | **int32** | The upper range value that a characteristic can take on | [optional] [default to null]
**ValueType** | **string** | A kind of value that the characteristic can take on, such as numeric, text, and so forth | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


