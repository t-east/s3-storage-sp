# {{classname}}

All URIs are relative to *http://orion.lab.fiware.org/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntity**](EntitiesApi.md#CreateEntity) | **Post** /v2/entities | Create Entity
[**ListEntities**](EntitiesApi.md#ListEntities) | **Get** /v2/entities | List Entities
[**RemoveEntity**](EntitiesApi.md#RemoveEntity) | **Delete** /v2/entities/{entityId} | Remove Entity
[**ReplaceAllEntityAttributes**](EntitiesApi.md#ReplaceAllEntityAttributes) | **Put** /v2/entities/{entityId}/attrs | Replace all entity attributes
[**RetrieveEntity**](EntitiesApi.md#RetrieveEntity) | **Get** /v2/entities/{entityId} | Retrieve Entity
[**RetrieveEntityAttributes**](EntitiesApi.md#RetrieveEntityAttributes) | **Get** /v2/entities/{entityId}/attrs | Retrieve Entity Attributes
[**UpdateExistingEntityAttributes**](EntitiesApi.md#UpdateExistingEntityAttributes) | **Patch** /v2/entities/{entityId}/attrs | Update Existing Entity Attributes
[**UpdateOrAppendEntityAttributes**](EntitiesApi.md#UpdateOrAppendEntityAttributes) | **Post** /v2/entities/{entityId}/attrs | Update or Append Entity Attributes

# **CreateEntity**
> CreateEntity(ctx, body, contentType, optional)
Create Entity

The payload is an object representing the entity to be created. The object follows the JSON entity representation format (described in a \"JSON Entity Representation\" section). Response: * Successful operation uses 201 Created (if upsert option is not used) or 204 No Content (if   upsert option is used). Response includes a `Location` header with the URL of the   created entity. * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEntityRequest**](CreateEntityRequest.md)|  | 
  **contentType** | **string**|  | 
 **optional** | ***EntitiesApiCreateEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiCreateEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **options** | **optional.**| Options dictionary | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEntities**
> []ListEntitiesResponse ListEntities(ctx, optional)
List Entities

Retrieves a list of entities that match different criteria by id, type, pattern matching (either id or type) and/or those which match a query or geographical query (see [Simple Query Language](#simple_query_language) and  [Geographical Queries](#geographical_queries)). A given entity has to match all the criteria to be retrieved (i.e., the criteria is combined in a logical AND way). Note that pattern matching query parameters are incompatible (i.e. mutually exclusive) with their corresponding exact matching parameters, i.e. `idPattern` with `id` and `typePattern` with `type`. The response payload is an array containing one object per matching entity. Each entity follows the JSON entity representation format (described in \"JSON Entity Representation\" section). Response code: * Successful operation uses 200 OK * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntitiesApiListEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiListEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| A comma-separated list of elements. Retrieve entities whose ID matches one of the elements in the list. Incompatible with &#x60;idPattern&#x60;. | 
 **type_** | **optional.String**| comma-separated list of elements. Retrieve entities whose type matches one of the elements in the list. Incompatible with &#x60;typePattern&#x60;. | 
 **idPattern** | **optional.String**| A correctly formated regular expression. Retrieve entities whose ID matches the regular expression. Incompatible with &#x60;id&#x60;. | 
 **typePattern** | **optional.String**| A correctly formated regular expression. Retrieve entities whose type matches the regular expression. Incompatible with &#x60;type&#x60;. | 
 **q** | **optional.String**| A query expression, composed of a list of statements separated by &#x60;;&#x60;, i.e., q&#x3D;statement1;statement2;statement3. See [Simple Query Language specification](#simple_query_language). | 
 **mq** | **optional.String**| A query expression for attribute metadata, composed of a list of statements separated by &#x60;;&#x60;, i.e., mq&#x3D;statement1;statement2;statement3. See [Simple Query Language specification](#simple_query_language). | 
 **georel** | **optional.String**| Spatial relationship between matching entities and a reference shape. See [Geographical Queries](#geographical_queries). | 
 **geometry** | **optional.String**| Geografical area to which the query is restricted. See [Geographical Queries](#geographical_queries). | 
 **coords** | **optional.String**| List of latitude-longitude pairs of coordinates separated by &#x27;;&#x27;. See [Geographical Queries](#geographical_queries). | 
 **limit** | **optional.Float64**| Limits the number of entities to be retrieved | 
 **offset** | **optional.Float64**| Establishes the offset from where entities are retrieved | 
 **attrs** | **optional.String**| Comma-separated list of attribute names whose data are to be included in the response. The attributes are retrieved in the order specified by this parameter. If this parameter is not included, the attributes are retrieved in arbitrary order. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. | 
 **metadata** | **optional.String**| A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. | 
 **orderBy** | **optional.String**| Criteria for ordering results. See \&quot;Ordering Results\&quot; section for details. | 
 **options** | **optional.String**| Options dictionary | 

### Return type

[**[]ListEntitiesResponse**](ListEntitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEntity**
> RemoveEntity(ctx, entityId, optional)
Remove Entity

Delete the entity. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| Id of the entity to be deleted | 
 **optional** | ***EntitiesApiRemoveEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiRemoveEntityOpts struct
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

# **ReplaceAllEntityAttributes**
> ReplaceAllEntityAttributes(ctx, body, entityId, contentType, optional)
Replace all entity attributes

The request payload is an object representing the new entity attributes. The object follows the JSON entity representation format (described in a \"JSON Entity Representation\" above), except that `id` and `type` are not allowed. The attributes previously existing in the entity are removed and replaced by the ones in the request. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReplaceAllEntityAttributesRequest**](ReplaceAllEntityAttributesRequest.md)|  | 
  **entityId** | **string**| Id of the entity in question. | 
  **contentType** | **string**|  | 
 **optional** | ***EntitiesApiReplaceAllEntityAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiReplaceAllEntityAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **optional.**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 
 **options** | **optional.**| Operations options | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveEntity**
> RetrieveEntityResponse RetrieveEntity(ctx, entityId, optional)
Retrieve Entity

The response is an object representing the entity identified by the ID. The object follows the JSON entity representation format (described in \"JSON Entity Representation\" section). This operation must return one entity element only, but there may be more than one entity with the same ID (e.g. entities with same ID but different types). In such case, an error message is returned, with the HTTP status code set to 409 Conflict. Response: * Successful operation uses 200 OK * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| Id of the entity to be retrieved | 
 **optional** | ***EntitiesApiRetrieveEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiRetrieveEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 
 **attrs** | **optional.String**| Comma-separated list of attribute names whose data must be included in the response. The attributes are retrieved in the order specified by this parameter. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. If this parameter is not included, the attributes are retrieved in arbitrary order, and all the attributes of the entity are included in the response. | 
 **metadata** | **optional.String**| A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. | 
 **options** | **optional.String**| Options dictionary | 

### Return type

[**RetrieveEntityResponse**](RetrieveEntityResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveEntityAttributes**
> RetrieveEntityAttributesResponse RetrieveEntityAttributes(ctx, entityId, optional)
Retrieve Entity Attributes

This request is similar to retreiving the whole entity, however this one omits the `id` and `type` fields. Just like the general request of getting an entire entity, this operation must return only one entity element. If more than one entity with the same ID is found (e.g. entities with same ID but different type), an error message is returned, with the HTTP status code set to 409 Conflict. Response: * Successful operation uses 200 OK * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| Id of the entity to be retrieved | 
 **optional** | ***EntitiesApiRetrieveEntityAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiRetrieveEntityAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 
 **attrs** | **optional.String**| Comma-separated list of attribute names whose data are to be included in the response. The attributes are retrieved in the order specified by this parameter. If this parameter is not included, the attributes are retrieved in arbitrary order, and all the attributes of the entity are included in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. | 
 **metadata** | **optional.String**| A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. | 
 **options** | **optional.String**| Options dictionary | 

### Return type

[**RetrieveEntityAttributesResponse**](RetrieveEntityAttributesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExistingEntityAttributes**
> UpdateExistingEntityAttributes(ctx, body, entityId, contentType, optional)
Update Existing Entity Attributes

The request payload is an object representing the attributes to update. The object follows the JSON entity representation format (described in \"JSON Entity Representation\" section), except that `id` and `type` are not allowed. The entity attributes are updated with the ones in the payload. In addition to that, if one or more attributes in the payload doesn't exist in the entity, an error is returned. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateExistingEntityAttributesRequest**](UpdateExistingEntityAttributesRequest.md)|  | 
  **entityId** | **string**| Id of the entity to be updated | 
  **contentType** | **string**|  | 
 **optional** | ***EntitiesApiUpdateExistingEntityAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiUpdateExistingEntityAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **optional.**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 
 **options** | **optional.**| Operations options | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrAppendEntityAttributes**
> UpdateOrAppendEntityAttributes(ctx, body, entityId, contentType, optional)
Update or Append Entity Attributes

The request payload is an object representing the attributes to append or update. The object follows the JSON entity representation format (described in \"JSON Entity Representation\" section), except that `id` and `type` are not allowed. The entity attributes are updated with the ones in the payload, depending on whether the `append` operation option is used or not. * If `append` is not used: the entity attributes are updated (if they previously exist) or appended   (if they don't previously exist) with the ones in the payload. * If `append` is used (i.e. strict append semantics): all the attributes in the payload not   previously existing in the entity are appended. In addition to that, in case some of the   attributes in the payload already exist in the entity, an error is returned. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \"Error Responses\" for   more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrAppendEntityAttributesRequest**](UpdateOrAppendEntityAttributesRequest.md)|  | 
  **entityId** | **string**| Entity id to be updated | 
  **contentType** | **string**|  | 
 **optional** | ***EntitiesApiUpdateOrAppendEntityAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiUpdateOrAppendEntityAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **optional.**| Entity type, to avoid ambiguity in case there are several entities with the same entity id. | 
 **options** | **optional.**| Operations options | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

