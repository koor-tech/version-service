# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/version_service.proto](#api_v1_version_service-proto)
    - [DetailedProductVersions](#api-v1-DetailedProductVersions)
    - [DetailedVersion](#api-v1-DetailedVersion)
    - [OperatorRequest](#api-v1-OperatorRequest)
    - [OperatorResponse](#api-v1-OperatorResponse)
    - [ProductVersions](#api-v1-ProductVersions)
    - [VersionMatrix](#api-v1-VersionMatrix)
    - [VersionMatrix.CephEntry](#api-v1-VersionMatrix-CephEntry)
    - [VersionMatrix.KoorOperatorEntry](#api-v1-VersionMatrix-KoorOperatorEntry)
    - [VersionMatrix.KsdEntry](#api-v1-VersionMatrix-KsdEntry)
  
    - [VersionService](#api-v1-VersionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_version_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/version_service.proto
Koor Version Service

This defines the messages and services of the Koor Version Service, which is used to retreive information about product and dependency versions.


<a name="api-v1-DetailedProductVersions"></a>

### DetailedProductVersions
Represents a map of products to detailed versions, which include images or helm charts.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| koor_operator | [DetailedVersion](#api-v1-DetailedVersion) |  | The detailed Koor Operator version. |
| ksd | [DetailedVersion](#api-v1-DetailedVersion) |  | The detailed Koor Storage Distribution version. |
| ceph | [DetailedVersion](#api-v1-DetailedVersion) |  | The detailed Ceph version. |






<a name="api-v1-DetailedVersion"></a>

### DetailedVersion
Defines a detailed version of a product, which includes a container image or a helm chart.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | The product version, must be a valid semver. |
| image_uri | [string](#string) |  | The URI of the container image. |
| image_hash | [string](#string) |  | The hash of the container image. |
| helm_repository | [string](#string) |  | The URI of the helm repository. |
| helm_chart | [string](#string) |  | The name of the helm chart in the repository. |






<a name="api-v1-OperatorRequest"></a>

### OperatorRequest
Represents an operator request message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versions | [ProductVersions](#api-v1-ProductVersions) |  | A map of products to current versions. |






<a name="api-v1-OperatorResponse"></a>

### OperatorResponse
Represents an operator response message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versions | [DetailedProductVersions](#api-v1-DetailedProductVersions) |  | A map of products to the newest available versions with deiails. |






<a name="api-v1-ProductVersions"></a>

### ProductVersions
Represents a map of products to version strings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kube | [string](#string) |  | Kubernetes version, must be a valid semver. |
| koor_operator | [string](#string) |  | Koor Operator version, must be a valid semver. |
| ksd | [string](#string) |  | Koor Storage Distribution version, must be a valid semver. |
| ceph | [string](#string) |  | Ceph version, must be a valid semver. |






<a name="api-v1-VersionMatrix"></a>

### VersionMatrix
Represents a map of products with all available images or helm chart versions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| koor_operator | [VersionMatrix.KoorOperatorEntry](#api-v1-VersionMatrix-KoorOperatorEntry) | repeated | A map detailing the available Koor Operator versions. The keys are version strings and the values are version details. |
| ksd | [VersionMatrix.KsdEntry](#api-v1-VersionMatrix-KsdEntry) | repeated | A map detailing the available Koor Storage Distribution versions. The keys are version strings and the values are version details. |
| ceph | [VersionMatrix.CephEntry](#api-v1-VersionMatrix-CephEntry) | repeated | A map detailing the available Ceph versions. The keys are version strings and the values are version details. |






<a name="api-v1-VersionMatrix-CephEntry"></a>

### VersionMatrix.CephEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [DetailedVersion](#api-v1-DetailedVersion) |  |  |






<a name="api-v1-VersionMatrix-KoorOperatorEntry"></a>

### VersionMatrix.KoorOperatorEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [DetailedVersion](#api-v1-DetailedVersion) |  |  |






<a name="api-v1-VersionMatrix-KsdEntry"></a>

### VersionMatrix.KsdEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [DetailedVersion](#api-v1-DetailedVersion) |  |  |





 

 

 


<a name="api-v1-VersionService"></a>

### VersionService
The main service of the Koor Version Service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Operator | [OperatorRequest](#api-v1-OperatorRequest) | [OperatorResponse](#api-v1-OperatorResponse) | Used by the Koor operator to get the latest version of each depencency product based on the current product versions. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

