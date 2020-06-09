# \ProductOfferApi

All URIs are relative to *https://virtserver.swaggerhub.com/moznobkin/Product_Catalog_View/0.0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductOfferById**](ProductOfferApi.md#GetProductOfferById) | **Get** /api/catalog-projection/v1/productOffering/{offerId} | Retrieves a ProductOffer by ID


# **GetProductOfferById**
> GetProductOfferResponse GetProductOfferById(ctx, offerId, optional)
Retrieves a ProductOffer by ID

This operation retrieves a ProductOffer entity by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **offerId** | **string**| offerId | 
 **optional** | ***ProductOfferApiGetProductOfferByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductOfferApiGetProductOfferByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| JWT Token | 

### Return type

[**GetProductOfferResponse**](GetProductOfferResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

