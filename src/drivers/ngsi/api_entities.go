/*
 * FIWARE-NGSI v2 Specification
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type EntitiesApiService service

/*
EntitiesApiService Create Entity
The payload is an object representing the entity to be created. The object follows the JSON entity representation format (described in a \&quot;JSON Entity Representation\&quot; section). Response: * Successful operation uses 201 Created (if upsert option is not used) or 204 No Content (if   upsert option is used). Response includes a &#x60;Location&#x60; header with the URL of the   created entity. * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param contentType
 * @param optional nil or *EntitiesApiCreateEntityOpts - Optional Parameters:
     * @param "Options" (optional.String) -  Options dictionary

*/

type EntitiesApiCreateEntityOpts struct {
	Options optional.String
}

func (a *EntitiesApiService) CreateEntity(ctx context.Context, body CreateEntityRequest, contentType string, localVarOptionals *EntitiesApiCreateEntityOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/entities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
		localVarQueryParams.Add("options", parameterToString(localVarOptionals.Options.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["Content-Length"] = "[1024]"
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
EntitiesApiService List Entities
Retrieves a list of entities that match different criteria by id, type, pattern matching (either id or type) and/or those which match a query or geographical query (see [Simple Query Language](#simple_query_language) and  [Geographical Queries](#geographical_queries)). A given entity has to match all the criteria to be retrieved (i.e., the criteria is combined in a logical AND way). Note that pattern matching query parameters are incompatible (i.e. mutually exclusive) with their corresponding exact matching parameters, i.e. &#x60;idPattern&#x60; with &#x60;id&#x60; and &#x60;typePattern&#x60; with &#x60;type&#x60;. The response payload is an array containing one object per matching entity. Each entity follows the JSON entity representation format (described in \&quot;JSON Entity Representation\&quot; section). Response code: * Successful operation uses 200 OK * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *EntitiesApiListEntitiesOpts - Optional Parameters:
     * @param "Id" (optional.String) -  A comma-separated list of elements. Retrieve entities whose ID matches one of the elements in the list. Incompatible with &#x60;idPattern&#x60;.
     * @param "Type_" (optional.String) -  comma-separated list of elements. Retrieve entities whose type matches one of the elements in the list. Incompatible with &#x60;typePattern&#x60;.
     * @param "IdPattern" (optional.String) -  A correctly formated regular expression. Retrieve entities whose ID matches the regular expression. Incompatible with &#x60;id&#x60;.
     * @param "TypePattern" (optional.String) -  A correctly formated regular expression. Retrieve entities whose type matches the regular expression. Incompatible with &#x60;type&#x60;.
     * @param "Q" (optional.String) -  A query expression, composed of a list of statements separated by &#x60;;&#x60;, i.e., q&#x3D;statement1;statement2;statement3. See [Simple Query Language specification](#simple_query_language).
     * @param "Mq" (optional.String) -  A query expression for attribute metadata, composed of a list of statements separated by &#x60;;&#x60;, i.e., mq&#x3D;statement1;statement2;statement3. See [Simple Query Language specification](#simple_query_language).
     * @param "Georel" (optional.String) -  Spatial relationship between matching entities and a reference shape. See [Geographical Queries](#geographical_queries).
     * @param "Geometry" (optional.String) -  Geografical area to which the query is restricted. See [Geographical Queries](#geographical_queries).
     * @param "Coords" (optional.String) -  List of latitude-longitude pairs of coordinates separated by &#x27;;&#x27;. See [Geographical Queries](#geographical_queries).
     * @param "Limit" (optional.Float64) -  Limits the number of entities to be retrieved
     * @param "Offset" (optional.Float64) -  Establishes the offset from where entities are retrieved
     * @param "Attrs" (optional.String) -  Comma-separated list of attribute names whose data are to be included in the response. The attributes are retrieved in the order specified by this parameter. If this parameter is not included, the attributes are retrieved in arbitrary order. See \&quot;Filtering out attributes and metadata\&quot; section for more detail.
     * @param "Metadata" (optional.String) -  A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail.
     * @param "OrderBy" (optional.String) -  Criteria for ordering results. See \&quot;Ordering Results\&quot; section for details.
     * @param "Options" (optional.String) -  Options dictionary
@return []ListEntitiesResponse
*/

type EntitiesApiListEntitiesOpts struct {
	Id          optional.String
	Type_       optional.String
	IdPattern   optional.String
	TypePattern optional.String
	Q           optional.String
	Mq          optional.String
	Georel      optional.String
	Geometry    optional.String
	Coords      optional.String
	Limit       optional.Float64
	Offset      optional.Float64
	Attrs       optional.String
	Metadata    optional.String
	OrderBy     optional.String
	Options     optional.String
}

func (a *EntitiesApiService) ListEntities(ctx context.Context, localVarOptionals *EntitiesApiListEntitiesOpts) ([]ListEntitiesResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []ListEntitiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/entities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdPattern.IsSet() {
		localVarQueryParams.Add("idPattern", parameterToString(localVarOptionals.IdPattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TypePattern.IsSet() {
		localVarQueryParams.Add("typePattern", parameterToString(localVarOptionals.TypePattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Q.IsSet() {
		localVarQueryParams.Add("q", parameterToString(localVarOptionals.Q.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mq.IsSet() {
		localVarQueryParams.Add("mq", parameterToString(localVarOptionals.Mq.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Georel.IsSet() {
		localVarQueryParams.Add("georel", parameterToString(localVarOptionals.Georel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Geometry.IsSet() {
		localVarQueryParams.Add("geometry", parameterToString(localVarOptionals.Geometry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Coords.IsSet() {
		localVarQueryParams.Add("coords", parameterToString(localVarOptionals.Coords.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Attrs.IsSet() {
		localVarQueryParams.Add("attrs", parameterToString(localVarOptionals.Attrs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Metadata.IsSet() {
		localVarQueryParams.Add("metadata", parameterToString(localVarOptionals.Metadata.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderBy.IsSet() {
		localVarQueryParams.Add("orderBy", parameterToString(localVarOptionals.OrderBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
		localVarQueryParams.Add("options", parameterToString(localVarOptionals.Options.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []ListEntitiesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
EntitiesApiService Remove Entity
Delete the entity. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityId Id of the entity to be deleted
 * @param optional nil or *EntitiesApiRemoveEntityOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.

*/

type EntitiesApiRemoveEntityOpts struct {
	Type_ optional.String
}

func (a *EntitiesApiService) RemoveEntity(ctx context.Context, entityId string, localVarOptionals *EntitiesApiRemoveEntityOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}"
	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
EntitiesApiService Replace all entity attributes
The request payload is an object representing the new entity attributes. The object follows the JSON entity representation format (described in a \&quot;JSON Entity Representation\&quot; above), except that &#x60;id&#x60; and &#x60;type&#x60; are not allowed. The attributes previously existing in the entity are removed and replaced by the ones in the request. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param contentType
 * @param entityId Id of the entity in question.
 * @param optional nil or *EntitiesApiReplaceAllEntityAttributesOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.
     * @param "Options" (optional.String) -  Operations options

*/

type EntitiesApiReplaceAllEntityAttributesOpts struct {
	Type_   optional.String
	Options optional.String
}

/*
EntitiesApiService Retrieve Entity
The response is an object representing the entity identified by the ID. The object follows the JSON entity representation format (described in \&quot;JSON Entity Representation\&quot; section). This operation must return one entity element only, but there may be more than one entity with the same ID (e.g. entities with same ID but different types). In such case, an error message is returned, with the HTTP status code set to 409 Conflict. Response: * Successful operation uses 200 OK * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityId Id of the entity to be retrieved
 * @param optional nil or *EntitiesApiRetrieveEntityOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.
     * @param "Attrs" (optional.String) -  Comma-separated list of attribute names whose data must be included in the response. The attributes are retrieved in the order specified by this parameter. See \&quot;Filtering out attributes and metadata\&quot; section for more detail. If this parameter is not included, the attributes are retrieved in arbitrary order, and all the attributes of the entity are included in the response.
     * @param "Metadata" (optional.String) -  A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail.
     * @param "Options" (optional.String) -  Options dictionary
@return RetrieveEntityResponse
*/

type EntitiesApiRetrieveEntityOpts struct {
	Type_    optional.String
	Attrs    optional.String
	Metadata optional.String
	Options  optional.String
}

// func (a *EntitiesApiService) RetrieveEntity(ctx context.Context, entityId string, localVarOptionals *EntitiesApiRetrieveEntityOpts) (RetrieveEntityResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue RetrieveEntityResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
// 		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Attrs.IsSet() {
// 		localVarQueryParams.Add("attrs", parameterToString(localVarOptionals.Attrs.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Metadata.IsSet() {
// 		localVarQueryParams.Add("metadata", parameterToString(localVarOptionals.Metadata.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
// 		localVarQueryParams.Add("options", parameterToString(localVarOptionals.Options.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"application/json"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v RetrieveEntityResponse
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

/*
EntitiesApiService Retrieve Entity Attributes
This request is similar to retreiving the whole entity, however this one omits the &#x60;id&#x60; and &#x60;type&#x60; fields. Just like the general request of getting an entire entity, this operation must return only one entity element. If more than one entity with the same ID is found (e.g. entities with same ID but different type), an error message is returned, with the HTTP status code set to 409 Conflict. Response: * Successful operation uses 200 OK * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityId Id of the entity to be retrieved
 * @param optional nil or *EntitiesApiRetrieveEntityAttributesOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.
     * @param "Attrs" (optional.String) -  Comma-separated list of attribute names whose data are to be included in the response. The attributes are retrieved in the order specified by this parameter. If this parameter is not included, the attributes are retrieved in arbitrary order, and all the attributes of the entity are included in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail.
     * @param "Metadata" (optional.String) -  A list of metadata names to include in the response. See \&quot;Filtering out attributes and metadata\&quot; section for more detail.
     * @param "Options" (optional.String) -  Options dictionary
@return RetrieveEntityAttributesResponse
*/

type EntitiesApiRetrieveEntityAttributesOpts struct {
	Type_    optional.String
	Attrs    optional.String
	Metadata optional.String
	Options  optional.String
}

// func (a *EntitiesApiService) RetrieveEntityAttributes(ctx context.Context, entityId string, localVarOptionals *EntitiesApiRetrieveEntityAttributesOpts) (RetrieveEntityAttributesResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue RetrieveEntityAttributesResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}/attrs"
// 	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
// 		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Attrs.IsSet() {
// 		localVarQueryParams.Add("attrs", parameterToString(localVarOptionals.Attrs.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Metadata.IsSet() {
// 		localVarQueryParams.Add("metadata", parameterToString(localVarOptionals.Metadata.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
// 		localVarQueryParams.Add("options", parameterToString(localVarOptionals.Options.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"application/json"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v RetrieveEntityAttributesResponse
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// EntitiesApiService Update Existing Entity Attributes
// The request payload is an object representing the attributes to update. The object follows the JSON entity representation format (described in \&quot;JSON Entity Representation\&quot; section), except that &#x60;id&#x60; and &#x60;type&#x60; are not allowed. The entity attributes are updated with the ones in the payload. In addition to that, if one or more attributes in the payload doesn&#x27;t exist in the entity, an error is returned. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
//  * @param contentType
//  * @param entityId Id of the entity to be updated
//  * @param optional nil or *EntitiesApiUpdateExistingEntityAttributesOpts - Optional Parameters:
//      * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.
//      * @param "Options" (optional.String) -  Operations options

// */

// type EntitiesApiUpdateExistingEntityAttributesOpts struct {
// 	Type_   optional.String
// 	Options optional.String
// }

// func (a *EntitiesApiService) UpdateExistingEntityAttributes(ctx context.Context, body UpdateExistingEntityAttributesRequest, contentType string, entityId string, localVarOptionals *EntitiesApiUpdateExistingEntityAttributesOpts) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Patch")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}/attrs"
// 	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
// 		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
// 		localVarQueryParams.Add("options", parameterToString(localVarOptionals.Options.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// EntitiesApiService Update or Append Entity Attributes
// The request payload is an object representing the attributes to append or update. The object follows the JSON entity representation format (described in \&quot;JSON Entity Representation\&quot; section), except that &#x60;id&#x60; and &#x60;type&#x60; are not allowed. The entity attributes are updated with the ones in the payload, depending on whether the &#x60;append&#x60; operation option is used or not. * If &#x60;append&#x60; is not used: the entity attributes are updated (if they previously exist) or appended   (if they don&#x27;t previously exist) with the ones in the payload. * If &#x60;append&#x60; is used (i.e. strict append semantics): all the attributes in the payload not   previously existing in the entity are appended. In addition to that, in case some of the   attributes in the payload already exist in the entity, an error is returned. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
//  * @param contentType
//  * @param entityId Entity id to be updated
//  * @param optional nil or *EntitiesApiUpdateOrAppendEntityAttributesOpts - Optional Parameters:
//      * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.
//      * @param "Options" (optional.String) -  Operations options

// */

// type EntitiesApiUpdateOrAppendEntityAttributesOpts struct {
// 	Type_   optional.String
// 	Options optional.String
// }

// func (a *EntitiesApiService) UpdateOrAppendEntityAttributes(ctx context.Context, body UpdateOrAppendEntityAttributesRequest, contentType string, entityId string, localVarOptionals *EntitiesApiUpdateOrAppendEntityAttributesOpts) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Post")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}/attrs"
// 	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
// 		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
// 		localVarQueryParams.Add("options", parameterToString(localVarOptionals.Options.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }
