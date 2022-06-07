# {{classname}}

All URIs are relative to *http://orion.lab.fiware.org/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttributeData**](AttributesApi.md#GetAttributeData) | **Get** /v2/entities/{entityId}/attrs/{attrName} | Get attribute data
[**RemoveASingleAttribute**](AttributesApi.md#RemoveASingleAttribute) | **Delete** /v2/entities/{entityId}/attrs/{attrName} | Remove a Single Attribute
[**UpdateAttributeData**](AttributesApi.md#UpdateAttributeData) | **Put** /v2/entities/{entityId}/attrs/{attrName} | Update Attribute Data

# **GetAttributeData**
> GetAttributeDataResponse GetAttributeData(ctx, entityId, attrName, optional)
Get attribute data

Returns a JSON object with the attribute data of the attribute. The object follows the JSON representation for attributes (described in \"JSON Attribute Representation\" section). Response: * Successful operation uses 200 OK. * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| Id of the entity | 
  **attrName** | **string**| Name of the attribute to be retrieved. | 
 **optional** | ***AttributesApiGetAttributeDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttributesApiGetAttributeDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 
 **metadata** | **optional.String**| A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. | 

### Return type

[**GetAttributeDataResponse**](GetAttributeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveASingleAttribute**
> RemoveASingleAttribute(ctx, entityId, attrName, optional)
Remove a Single Attribute

Removes an entity attribute. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| Id of the entity. | 
  **attrName** | **string**| Attribute name. | 
 **optional** | ***AttributesApiRemoveASingleAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttributesApiRemoveASingleAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAttributeData**
> UpdateAttributeData(ctx, body, entityId, attrName, contentType, optional)
Update Attribute Data

The request payload is an object representing the new attribute data. Previous attribute data is replaced by the one in the request. The object follows the JSON representation for attributes (described in \"JSON Attribute Representation\" section). Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAttributeDataRequest**](UpdateAttributeDataRequest.md)|  | 
  **entityId** | **string**| Id of the entity to update | 
  **attrName** | **string**| Attribute name | 
  **contentType** | **string**|  | 
 **optional** | ***AttributesApiUpdateAttributeDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttributesApiUpdateAttributeDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **type_** | **optional.**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

