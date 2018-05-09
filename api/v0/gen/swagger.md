# Optikon


<a name="overview"></a>
## Overview
Optikon is an edge computing API for Kubernetes. It exposes Helm and the cluster registry for a comprehensive UI


### Version information
*Version* : 0.0.1


### Contact information
*Contact Email* : stelouie@cisco.com


### URI scheme
*BasePath* : /v0  
*Schemes* : HTTP, HTTPS


### Tags

* clusters : Everything about your clusters
* releases : Access to all helm releases




<a name="paths"></a>
## Paths

<a name="addcluster"></a>
### Add a new cluster to optikon
```
POST /clusters
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**body**  <br>*optional*|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.Cluster](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-cluster)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**201**|Created|[APIResponse](#apiresponse)|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### Tags

* clusters


<a name="getclusters"></a>
### Returns all clusters
```
GET /clusters
```


#### Description
Returns a list of clusters


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|< [io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.Cluster](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-cluster) > array|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Produces

* `application/json`


#### Tags

* clusters


<a name="getclusterbyid"></a>
### Find cluster by ID
```
GET /clusters/{clusterId}
```


#### Description
Returns a single cluster


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**clusterId**  <br>*required*|ID of cluster to return|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.Cluster](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-cluster)|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**404**|The specified resource was not found|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Produces

* `application/json`


#### Tags

* clusters


<a name="updatecluster"></a>
### Update an existing cluster
```
PUT /clusters/{clusterId}
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**clusterId**  <br>*required*|Cluster id to update|string|
|**Body**|**body**  <br>*optional*||[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.Cluster](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-cluster)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|No Content|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**404**|The specified resource was not found|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### Tags

* clusters


<a name="deletecluster"></a>
### Deletes a cluster
```
DELETE /clusters/{clusterId}
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**clusterId**  <br>*required*|Cluster id to delete|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|No Content|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**404**|The specified resource was not found|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Produces

* `application/json`


#### Tags

* clusters


<a name="addreleases"></a>
### Add a new release to the clusters
```
POST /releases
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**labels**  <br>*optional*|The node labels to identify applicable clusters|string|
|**FormData**|**chartTar**  <br>*required*|The file to upload|file|
|**FormData**|**name**  <br>*required*|The name of the helm release|string|
|**FormData**|**namespace**  <br>*required*|The kubernetes namespace to be used|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**201**|Created|[APIResponse](#apiresponse)|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Consumes

* `multipart/form-data`


#### Produces

* `application/json`


#### Tags

* releases


<a name="getreleases"></a>
### Returns all releases
```
GET /releases
```


#### Description
Returns a list of releases


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**labels**  <br>*optional*|The node labels to identify applicable clusters|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|< [release.Release](#release-release) > array|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Produces

* `application/json`


#### Tags

* releases


<a name="getreleasebyid"></a>
### Find release by ID
```
GET /releases/{releaseId}
```


#### Description
Returns a single release


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**releaseId**  <br>*required*|ID of release to return|string|
|**Query**|**labels**  <br>*optional*|The node labels to identify applicable clusters|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|[release.Release](#release-release)|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**404**|The specified resource was not found|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Produces

* `application/json`


#### Tags

* releases


<a name="updaterelease"></a>
### Update an existing release
```
PUT /releases/{releaseId}
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**releaseId**  <br>*required*|ID of release to update|string|
|**Query**|**labels**  <br>*optional*|The node labels to identify applicable clusters|string|
|**FormData**|**chartTar**  <br>*required*|The file to upload|file|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|No Content|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**404**|The specified resource was not found|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Consumes

* `multipart/form-data`


#### Produces

* `application/json`


#### Tags

* releases


<a name="deleterelease"></a>
### Delete release by ID
```
DELETE /releases/{releaseId}
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**releaseId**  <br>*required*|ID of release to return|string|
|**Query**|**labels**  <br>*optional*|The node labels to identify applicable clusters|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|No Content|
|**400**|Bad Request|[APIResponse](#apiresponse)|
|**401**|Unauthorized|[APIResponse](#apiresponse)|
|**404**|The specified resource was not found|[APIResponse](#apiresponse)|
|**500**|Internal Server Error|[APIResponse](#apiresponse)|


#### Produces

* `application/json`


#### Tags

* releases




<a name="definitions"></a>
## Definitions

<a name="apiresponse"></a>
### APIResponse

|Name|Schema|
|---|---|
|**code**  <br>*required*|string|
|**message**  <br>*required*|string|


<a name="chart-chart"></a>
### chart.Chart

|Name|Schema|
|---|---|
|**Config**  <br>*optional*|[chart.Config](#chart-config)|
|**Dependencies**  <br>*optional*|< [chart.Chart](#chart-chart) > array|
|**Files**  <br>*optional*|< [protobuf.Any](#protobuf-any) > array|
|**Metadata**  <br>*optional*|[chart.Metadata](#chart-metadata)|
|**Template**  <br>*optional*|< [chart.Template](#chart-template) > array|


<a name="chart-config"></a>
### chart.Config

|Name|Schema|
|---|---|
|**Raw**  <br>*optional*|string|
|**Values**  <br>*optional*|[chart.Map](#chart-map)|


<a name="chart-maintainer"></a>
### chart.Maintainer

|Name|Schema|
|---|---|
|**Email**  <br>*optional*|string|
|**Name**  <br>*optional*|string|
|**Url**  <br>*optional*|string|


<a name="chart-map"></a>
### chart.Map
*Type* : < [chart.Value](#chart-value) > array


<a name="chart-metadata"></a>
### chart.Metadata

|Name|Schema|
|---|---|
|**Annotations**  <br>*optional*|[chart.Map](#chart-map)|
|**ApiVersion**  <br>*optional*|string|
|**AppVersion**  <br>*optional*|string|
|**Condition**  <br>*optional*|string|
|**Deprecated**  <br>*optional*|boolean|
|**Description**  <br>*optional*|string|
|**Engine**  <br>*optional*|enum (UNKNOWN, GOTPL)|
|**Home**  <br>*optional*|string|
|**Icon**  <br>*optional*|string|
|**Keywords**  <br>*optional*|< string > array|
|**KubeVersion**  <br>*optional*|string|
|**Maintainers**  <br>*optional*|< [chart.Maintainer](#chart-maintainer) > array|
|**Name**  <br>*optional*|string|
|**Sources**  <br>*optional*|< string > array|
|**Tags**  <br>*optional*|string|
|**TillerVersion**  <br>*optional*|string|
|**Version**  <br>*optional*|string|


<a name="chart-template"></a>
### chart.Template

|Name|Description|Schema|
|---|---|---|
|**Data**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**Name**  <br>*optional*||string|


<a name="chart-value"></a>
### chart.Value

|Name|Schema|
|---|---|
|**Key**  <br>*optional*|string|
|**Value**  <br>*optional*|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-initializer"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.Initializer
Initializer is information about an initializer that has not yet completed.


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*required*|name of the process that is responsible for initializing this object.|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-initializers"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.Initializers
Initializers tracks the progress of initialization.


|Name|Description|Schema|
|---|---|---|
|**pending**  <br>*required*|Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.|< [io.k8s.apimachinery.pkg.apis.meta.v1.Initializer](#io-k8s-apimachinery-pkg-apis-meta-v1-initializer) > array|
|**result**  <br>*optional*|If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.|[io.k8s.apimachinery.pkg.apis.meta.v1.Status](#io-k8s-apimachinery-pkg-apis-meta-v1-status)|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-listmeta"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta
ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.


|Name|Description|Schema|
|---|---|---|
|**continue**  <br>*optional*|continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response.|string|
|**resourceVersion**  <br>*optional*|String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency|string|
|**selfLink**  <br>*optional*|selfLink is a URL representing this object. Populated by the system. Read-only.|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-objectmeta"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.


|Name|Description|Schema|
|---|---|---|
|**annotations**  <br>*optional*|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations|< string, string > map|
|**clusterName**  <br>*optional*|The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.|string|
|**creationTimestamp**  <br>*optional*|CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.<br><br>Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata|[io.k8s.apimachinery.pkg.apis.meta.v1.Time](#io-k8s-apimachinery-pkg-apis-meta-v1-time)|
|**deletionGracePeriodSeconds**  <br>*optional*|Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.|integer (int64)|
|**deletionTimestamp**  <br>*optional*|DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field. Once set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.<br><br>Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata|[io.k8s.apimachinery.pkg.apis.meta.v1.Time](#io-k8s-apimachinery-pkg-apis-meta-v1-time)|
|**finalizers**  <br>*optional*|Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.|< string > array|
|**generateName**  <br>*optional*|GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.<br><br>If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).<br><br>Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency|string|
|**generation**  <br>*optional*|A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.|integer (int64)|
|**initializers**  <br>*optional*|An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.<br><br>When an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.|[io.k8s.apimachinery.pkg.apis.meta.v1.Initializers](#io-k8s-apimachinery-pkg-apis-meta-v1-initializers)|
|**labels**  <br>*optional*|Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels|< string, string > map|
|**name**  <br>*optional*|Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names|string|
|**namespace**  <br>*optional*|Namespace defines the space within each name must be unique. An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.<br><br>Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces|string|
|**ownerReferences**  <br>*optional*|List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.|< [io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference](#io-k8s-apimachinery-pkg-apis-meta-v1-ownerreference) > array|
|**resourceVersion**  <br>*optional*|An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.<br><br>Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency|string|
|**selfLink**  <br>*optional*|SelfLink is a URL representing this object. Populated by the system. Read-only.|string|
|**uid**  <br>*optional*|UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.<br><br>Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-ownerreference"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
OwnerReference contains enough information to let you identify an owning object. Currently, an owning object must be in the same namespace, so there is no namespace field.


|Name|Description|Schema|
|---|---|---|
|**apiVersion**  <br>*required*|API version of the referent.|string|
|**blockOwnerDeletion**  <br>*optional*|If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.|boolean|
|**controller**  <br>*optional*|If true, this reference points to the managing controller.|boolean|
|**kind**  <br>*required*|Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds|string|
|**name**  <br>*required*|Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names|string|
|**uid**  <br>*required*|UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-status"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.Status
Status is a return value for calls that don't return other objects.


|Name|Description|Schema|
|---|---|---|
|**apiVersion**  <br>*optional*|APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources|string|
|**code**  <br>*optional*|Suggested HTTP return code for this status, 0 if not set.|integer (int32)|
|**details**  <br>*optional*|Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.|[io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails](#io-k8s-apimachinery-pkg-apis-meta-v1-statusdetails)|
|**kind**  <br>*optional*|Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds|string|
|**message**  <br>*optional*|A human-readable description of the status of this operation.|string|
|**metadata**  <br>*optional*|Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds|[io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta](#io-k8s-apimachinery-pkg-apis-meta-v1-listmeta)|
|**reason**  <br>*optional*|A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.|string|
|**status**  <br>*optional*|Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-statuscause"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause
StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.


|Name|Description|Schema|
|---|---|---|
|**field**  <br>*optional*|The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.<br><br>Examples:<br>  "name" - the field "name" on the current resource<br>  "items[0].name" - the field "name" on the first array entry in "items"|string|
|**message**  <br>*optional*|A human-readable description of the cause of the error.  This field may be presented as-is to a reader.|string|
|**reason**  <br>*optional*|A machine-readable description of the cause of the error. If this value is empty there is no information available.|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-statusdetails"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails
StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.


|Name|Description|Schema|
|---|---|---|
|**causes**  <br>*optional*|The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.|< [io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause](#io-k8s-apimachinery-pkg-apis-meta-v1-statuscause) > array|
|**group**  <br>*optional*|The group attribute of the resource associated with the status StatusReason.|string|
|**kind**  <br>*optional*|The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds|string|
|**name**  <br>*optional*|The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).|string|
|**retryAfterSeconds**  <br>*optional*|If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.|integer (int32)|
|**uid**  <br>*optional*|UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids|string|


<a name="io-k8s-apimachinery-pkg-apis-meta-v1-time"></a>
### io.k8s.apimachinery.pkg.apis.meta.v1.Time
*Type* : string (date-time)


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-authinfo"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.AuthInfo
AuthInfo holds public information that describes how a client can get credentials to access the cluster. For example, OAuth2 client registration endpoints and supported flows, or Kerberos servers locations.

It should not hold any private or sensitive information.


|Name|Description|Schema|
|---|---|---|
|**providers**  <br>*optional*|AuthProviders is a list of configurations for auth providers.|< [io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.AuthProviderConfig](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-authproviderconfig) > array|


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-authproviderconfig"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.AuthProviderConfig
AuthProviderConfig contains the information necessary for a client to authenticate to a Kubernetes API server. It is modeled after k8s.io/client-go/tools/clientcmd/api/v1.AuthProviderConfig.


|Name|Description|Schema|
|---|---|---|
|**config**  <br>*optional*|Config is a map of values that contains the information necessary for a client to determine how to authenticate to a Kubernetes API server.|< string, string > map|
|**name**  <br>*optional*|Name is the name of this configuration.|string|
|**type**  <br>*optional*|Type contains type information about this auth provider. Clients of the cluster registry should use this field to differentiate between different kinds of authentication providers.|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.AuthProviderType](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-authprovidertype)|


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-authprovidertype"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.AuthProviderType
AuthProviderType contains metadata about the auth provider. It should be used by clients to differentiate between different kinds of auth providers, and to select a relevant provider for the client's configuration. For example, a controller would look for a provider type that denotes a service account that it should use to access the cluster, whereas a user would look for a provider type that denotes an authentication system from which they should request a token.


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*optional*|Name is the name of the auth provider.|string|


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-cluster"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.Cluster
Cluster contains information about a cluster in a cluster registry.


|Name|Description|Schema|
|---|---|---|
|**apiVersion**  <br>*optional*|APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources|string|
|**kind**  <br>*optional*|Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds|string|
|**metadata**  <br>*optional*|Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata|[io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta](#io-k8s-apimachinery-pkg-apis-meta-v1-objectmeta)|
|**spec**  <br>*optional*|Spec is the specification of the cluster. This may or may not be reconciled by an active controller.|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.ClusterSpec](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-clusterspec)|
|**status**  <br>*optional*|Status is the status of the cluster. It is optional, and can be left nil to imply that the cluster status is not being reported.|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.ClusterStatus](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-clusterstatus)|


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-clusterspec"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.ClusterSpec
ClusterSpec contains the specification of a cluster.


|Name|Description|Schema|
|---|---|---|
|**authInfo**  <br>*optional*|AuthInfo contains public information that can be used to authenticate to and authorize with this cluster. It is not meant to store private information (e.g., tokens or client certificates) and cluster registry implementations are not expected to provide hardened storage for secrets.|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.AuthInfo](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-authinfo)|
|**kubernetesApiEndpoints**  <br>*optional*|KubernetesAPIEndpoints represents the endpoints of the API server for this cluster.|[io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.KubernetesAPIEndpoints](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-kubernetesapiendpoints)|


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-clusterstatus"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.ClusterStatus
ClusterStatus contains the status of a cluster.


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-kubernetesapiendpoints"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.KubernetesAPIEndpoints
KubernetesAPIEndpoints represents the endpoints for one and only one Kubernetes API server.


|Name|Description|Schema|
|---|---|---|
|**caBundle**  <br>*optional*|CABundle contains the certificate authority information.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**serverEndpoints**  <br>*optional*|ServerEndpoints specifies the address(es) of the Kubernetes API server’s network identity or identities.|< [io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.ServerAddressByClientCIDR](#io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-serveraddressbyclientcidr) > array|


<a name="io-k8s-cluster-registry-pkg-apis-clusterregistry-v1alpha1-serveraddressbyclientcidr"></a>
### io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.ServerAddressByClientCIDR
ServerAddressByClientCIDR helps clients determine the server address that they should use, depending on the ClientCIDR that they match.


|Name|Description|Schema|
|---|---|---|
|**clientCIDR**  <br>*optional*|The CIDR with which clients can match their IP to figure out if they should use the corresponding server address.|string|
|**serverAddress**  <br>*optional*|Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.|string|


<a name="protobuf-any"></a>
### protobuf.Any

|Name|Description|Schema|
|---|---|---|
|**Type_Url**  <br>*optional*||string|
|**Value**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="release-files"></a>
### release.Files

|Name|Description|Schema|
|---|---|---|
|**Key**  <br>*optional*||string|
|**Value**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="release-hook"></a>
### release.Hook

|Name|Schema|
|---|---|
|**DeletePolicies**  <br>*optional*|< enum (SUCCEEDED, FAILED, BEFORE_HOOK_CREATION) > array|
|**Events**  <br>*optional*|< enum (UNKNOWN, PRE_INSTALL, POST_INSTALL, PRE_DELETE, POST_DELETE, PRE_UPGRADE, POST_UPGRADE, PRE_ROLLBACK, POST_ROLLBACK, RELEASE_TEST_SUCCESS, RELEASE_TEST_FAILURE) > array|
|**Kind**  <br>*optional*|string|
|**LastRun**  <br>*optional*|string|
|**Manifest**  <br>*optional*|string|
|**Name**  <br>*optional*|string|
|**Path**  <br>*optional*|string|
|**Weight**  <br>*optional*|integer (int32)|


<a name="release-info"></a>
### release.Info

|Name|Schema|
|---|---|
|**Deleted**  <br>*optional*|string|
|**Description**  <br>*optional*|string|
|**FirstDeployed**  <br>*optional*|string|
|**LastDeployed**  <br>*optional*|string|
|**Status**  <br>*optional*|[release.Status](#release-status)|


<a name="release-release"></a>
### release.Release

|Name|Schema|
|---|---|
|**Chart**  <br>*optional*|[chart.Chart](#chart-chart)|
|**Config**  <br>*optional*|[chart.Config](#chart-config)|
|**Hooks**  <br>*optional*|< [release.Hook](#release-hook) > array|
|**Info**  <br>*optional*|[release.Info](#release-info)|
|**Manifest**  <br>*optional*|string|
|**Name**  <br>*optional*|string|
|**Namespace**  <br>*optional*|string|
|**OnCluster**  <br>*optional*|string|
|**Version**  <br>*optional*|integer (int32)|


<a name="release-status"></a>
### release.Status

|Name|Schema|
|---|---|
|**Code**  <br>*optional*|enum (UNKNOWN, DEPLOYED, DELETED, SUSPENDED, FAILED, DELETING, PENDING_INSTALL, PENDING_UPGRADE, PENDING_ROLLBACK)|
|**LastTestSuiteRun**  <br>*optional*|[release.TestSuite](#release-testsuite)|
|**Notes**  <br>*optional*|string|
|**Resources**  <br>*optional*|string|


<a name="release-testrun"></a>
### release.TestRun

|Name|Schema|
|---|---|
|**CompletedAt**  <br>*optional*|string|
|**Info**  <br>*optional*|string|
|**Name**  <br>*optional*|string|
|**StartedAt**  <br>*optional*|string|
|**Status**  <br>*optional*|string|


<a name="release-testsuite"></a>
### release.TestSuite

|Name|Schema|
|---|---|
|**CompletedAt**  <br>*optional*|string|
|**Results**  <br>*optional*|< [release.TestRun](#release-testrun) > array|
|**StartedAt**  <br>*optional*|string|





