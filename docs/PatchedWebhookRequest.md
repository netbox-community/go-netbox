# PatchedWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeCreate** | Pointer to **bool** | Triggers when a matching object is created. | [optional] 
**TypeUpdate** | Pointer to **bool** | Triggers when a matching object is updated. | [optional] 
**TypeDelete** | Pointer to **bool** | Triggers when a matching object is deleted. | [optional] 
**TypeJobStart** | Pointer to **bool** | Triggers when a job for a matching object is started. | [optional] 
**TypeJobEnd** | Pointer to **bool** | Triggers when a job for a matching object terminates. | [optional] 
**PayloadUrl** | Pointer to **string** | This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HttpMethod** | Pointer to **string** | * &#x60;GET&#x60; - GET * &#x60;POST&#x60; - POST * &#x60;PUT&#x60; - PUT * &#x60;PATCH&#x60; - PATCH * &#x60;DELETE&#x60; - DELETE | [optional] 
**HttpContentType** | Pointer to **string** | The complete list of official content types is available &lt;a href&#x3D;\&quot;https://www.iana.org/assignments/media-types/media-types.xhtml\&quot;&gt;here&lt;/a&gt;. | [optional] 
**AdditionalHeaders** | Pointer to **string** | User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format &lt;code&gt;Name: Value&lt;/code&gt;. Jinja2 template processing is supported with the same context as the request body (below). | [optional] 
**BodyTemplate** | Pointer to **string** | Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: &lt;code&gt;event&lt;/code&gt;, &lt;code&gt;model&lt;/code&gt;, &lt;code&gt;timestamp&lt;/code&gt;, &lt;code&gt;username&lt;/code&gt;, &lt;code&gt;request_id&lt;/code&gt;, and &lt;code&gt;data&lt;/code&gt;. | [optional] 
**Secret** | Pointer to **string** | When provided, the request will include a &#39;X-Hook-Signature&#39; header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request. | [optional] 
**Conditions** | Pointer to **map[string]interface{}** | A set of conditions which determine whether the webhook will be generated. | [optional] 
**SslVerification** | Pointer to **bool** | Enable SSL certificate verification. Disable with caution! | [optional] 
**CaFilePath** | Pointer to **NullableString** | The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults. | [optional] 

## Methods

### NewPatchedWebhookRequest

`func NewPatchedWebhookRequest() *PatchedWebhookRequest`

NewPatchedWebhookRequest instantiates a new PatchedWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWebhookRequestWithDefaults

`func NewPatchedWebhookRequestWithDefaults() *PatchedWebhookRequest`

NewPatchedWebhookRequestWithDefaults instantiates a new PatchedWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *PatchedWebhookRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedWebhookRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedWebhookRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedWebhookRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeCreate

`func (o *PatchedWebhookRequest) GetTypeCreate() bool`

GetTypeCreate returns the TypeCreate field if non-nil, zero value otherwise.

### GetTypeCreateOk

`func (o *PatchedWebhookRequest) GetTypeCreateOk() (*bool, bool)`

GetTypeCreateOk returns a tuple with the TypeCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCreate

`func (o *PatchedWebhookRequest) SetTypeCreate(v bool)`

SetTypeCreate sets TypeCreate field to given value.

### HasTypeCreate

`func (o *PatchedWebhookRequest) HasTypeCreate() bool`

HasTypeCreate returns a boolean if a field has been set.

### GetTypeUpdate

`func (o *PatchedWebhookRequest) GetTypeUpdate() bool`

GetTypeUpdate returns the TypeUpdate field if non-nil, zero value otherwise.

### GetTypeUpdateOk

`func (o *PatchedWebhookRequest) GetTypeUpdateOk() (*bool, bool)`

GetTypeUpdateOk returns a tuple with the TypeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUpdate

`func (o *PatchedWebhookRequest) SetTypeUpdate(v bool)`

SetTypeUpdate sets TypeUpdate field to given value.

### HasTypeUpdate

`func (o *PatchedWebhookRequest) HasTypeUpdate() bool`

HasTypeUpdate returns a boolean if a field has been set.

### GetTypeDelete

`func (o *PatchedWebhookRequest) GetTypeDelete() bool`

GetTypeDelete returns the TypeDelete field if non-nil, zero value otherwise.

### GetTypeDeleteOk

`func (o *PatchedWebhookRequest) GetTypeDeleteOk() (*bool, bool)`

GetTypeDeleteOk returns a tuple with the TypeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDelete

`func (o *PatchedWebhookRequest) SetTypeDelete(v bool)`

SetTypeDelete sets TypeDelete field to given value.

### HasTypeDelete

`func (o *PatchedWebhookRequest) HasTypeDelete() bool`

HasTypeDelete returns a boolean if a field has been set.

### GetTypeJobStart

`func (o *PatchedWebhookRequest) GetTypeJobStart() bool`

GetTypeJobStart returns the TypeJobStart field if non-nil, zero value otherwise.

### GetTypeJobStartOk

`func (o *PatchedWebhookRequest) GetTypeJobStartOk() (*bool, bool)`

GetTypeJobStartOk returns a tuple with the TypeJobStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeJobStart

`func (o *PatchedWebhookRequest) SetTypeJobStart(v bool)`

SetTypeJobStart sets TypeJobStart field to given value.

### HasTypeJobStart

`func (o *PatchedWebhookRequest) HasTypeJobStart() bool`

HasTypeJobStart returns a boolean if a field has been set.

### GetTypeJobEnd

`func (o *PatchedWebhookRequest) GetTypeJobEnd() bool`

GetTypeJobEnd returns the TypeJobEnd field if non-nil, zero value otherwise.

### GetTypeJobEndOk

`func (o *PatchedWebhookRequest) GetTypeJobEndOk() (*bool, bool)`

GetTypeJobEndOk returns a tuple with the TypeJobEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeJobEnd

`func (o *PatchedWebhookRequest) SetTypeJobEnd(v bool)`

SetTypeJobEnd sets TypeJobEnd field to given value.

### HasTypeJobEnd

`func (o *PatchedWebhookRequest) HasTypeJobEnd() bool`

HasTypeJobEnd returns a boolean if a field has been set.

### GetPayloadUrl

`func (o *PatchedWebhookRequest) GetPayloadUrl() string`

GetPayloadUrl returns the PayloadUrl field if non-nil, zero value otherwise.

### GetPayloadUrlOk

`func (o *PatchedWebhookRequest) GetPayloadUrlOk() (*string, bool)`

GetPayloadUrlOk returns a tuple with the PayloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadUrl

`func (o *PatchedWebhookRequest) SetPayloadUrl(v string)`

SetPayloadUrl sets PayloadUrl field to given value.

### HasPayloadUrl

`func (o *PatchedWebhookRequest) HasPayloadUrl() bool`

HasPayloadUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHttpMethod

`func (o *PatchedWebhookRequest) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *PatchedWebhookRequest) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *PatchedWebhookRequest) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *PatchedWebhookRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpContentType

`func (o *PatchedWebhookRequest) GetHttpContentType() string`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *PatchedWebhookRequest) GetHttpContentTypeOk() (*string, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *PatchedWebhookRequest) SetHttpContentType(v string)`

SetHttpContentType sets HttpContentType field to given value.

### HasHttpContentType

`func (o *PatchedWebhookRequest) HasHttpContentType() bool`

HasHttpContentType returns a boolean if a field has been set.

### GetAdditionalHeaders

`func (o *PatchedWebhookRequest) GetAdditionalHeaders() string`

GetAdditionalHeaders returns the AdditionalHeaders field if non-nil, zero value otherwise.

### GetAdditionalHeadersOk

`func (o *PatchedWebhookRequest) GetAdditionalHeadersOk() (*string, bool)`

GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHeaders

`func (o *PatchedWebhookRequest) SetAdditionalHeaders(v string)`

SetAdditionalHeaders sets AdditionalHeaders field to given value.

### HasAdditionalHeaders

`func (o *PatchedWebhookRequest) HasAdditionalHeaders() bool`

HasAdditionalHeaders returns a boolean if a field has been set.

### GetBodyTemplate

`func (o *PatchedWebhookRequest) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *PatchedWebhookRequest) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *PatchedWebhookRequest) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.

### HasBodyTemplate

`func (o *PatchedWebhookRequest) HasBodyTemplate() bool`

HasBodyTemplate returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetConditions

`func (o *PatchedWebhookRequest) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PatchedWebhookRequest) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PatchedWebhookRequest) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PatchedWebhookRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PatchedWebhookRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PatchedWebhookRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetSslVerification

`func (o *PatchedWebhookRequest) GetSslVerification() bool`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *PatchedWebhookRequest) GetSslVerificationOk() (*bool, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *PatchedWebhookRequest) SetSslVerification(v bool)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *PatchedWebhookRequest) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetCaFilePath

`func (o *PatchedWebhookRequest) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *PatchedWebhookRequest) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *PatchedWebhookRequest) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *PatchedWebhookRequest) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.

### SetCaFilePathNil

`func (o *PatchedWebhookRequest) SetCaFilePathNil(b bool)`

 SetCaFilePathNil sets the value for CaFilePath to be an explicit nil

### UnsetCaFilePath
`func (o *PatchedWebhookRequest) UnsetCaFilePath()`

UnsetCaFilePath ensures that no value is present for CaFilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


