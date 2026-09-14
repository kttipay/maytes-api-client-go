# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCheckout**](DefaultAPI.md#CancelCheckout) | **Post** /api/merchant/v1/checkouts/{checkoutUuid}/cancel | Cancel authorized checkout
[**CaptureCheckout**](DefaultAPI.md#CaptureCheckout) | **Post** /api/merchant/v1/checkouts/{checkoutUuid}/capture | Capture authorized checkout
[**CreateCheckout**](DefaultAPI.md#CreateCheckout) | **Post** /api/merchant/v1/checkouts | Create a new checkout
[**GetCheckout**](DefaultAPI.md#GetCheckout) | **Get** /api/merchant/v1/checkouts/{checkoutUuid} | Get checkout details
[**GetHealth**](DefaultAPI.md#GetHealth) | **Get** /api/health | Health check
[**GetOAuthToken**](DefaultAPI.md#GetOAuthToken) | **Post** /oauth/token | Issue an access token
[**GetSettlement**](DefaultAPI.md#GetSettlement) | **Get** /api/merchant/v1/settlements/{uuid} | Get a settlement by UUID
[**ListSettlements**](DefaultAPI.md#ListSettlements) | **Get** /api/merchant/v1/settlements | List settlements for this merchant
[**RefundCheckout**](DefaultAPI.md#RefundCheckout) | **Post** /api/merchant/v1/checkouts/{checkoutUuid}/refund | Refund a captured checkout



## CancelCheckout

> CancelCheckoutResponse CancelCheckout(ctx, checkoutUuid).CancelCheckoutRequest(cancelCheckoutRequest).Execute()

Cancel authorized checkout

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	checkoutUuid := "checkoutUuid_example" // string | 
	cancelCheckoutRequest := *openapiclient.NewCancelCheckoutRequest() // CancelCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CancelCheckout(context.Background(), checkoutUuid).CancelCheckoutRequest(cancelCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CancelCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelCheckout`: CancelCheckoutResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CancelCheckout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelCheckoutRequest** | [**CancelCheckoutRequest**](CancelCheckoutRequest.md) |  | 

### Return type

[**CancelCheckoutResponse**](CancelCheckoutResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaptureCheckout

> CaptureCheckoutResponse CaptureCheckout(ctx, checkoutUuid).CaptureCheckoutRequest(captureCheckoutRequest).Execute()

Capture authorized checkout

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	checkoutUuid := "checkoutUuid_example" // string | 
	captureCheckoutRequest := *openapiclient.NewCaptureCheckoutRequest() // CaptureCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CaptureCheckout(context.Background(), checkoutUuid).CaptureCheckoutRequest(captureCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CaptureCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaptureCheckout`: CaptureCheckoutResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CaptureCheckout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCaptureCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captureCheckoutRequest** | [**CaptureCheckoutRequest**](CaptureCheckoutRequest.md) |  | 

### Return type

[**CaptureCheckoutResponse**](CaptureCheckoutResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCheckout

> CreateCheckoutResponse CreateCheckout(ctx).CreateCheckoutRequest(createCheckoutRequest).Execute()

Create a new checkout

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	createCheckoutRequest := *openapiclient.NewCreateCheckoutRequest(int32(123), "Currency_example", "ReturnUrl_example", "CancelUrl_example", []openapiclient.LineItem{openapiclient.LineItem{DigitalGoodsLineItem: openapiclient.NewDigitalGoodsLineItem("ItemRef_example", "Name_example", int32(123), int32(123), "Currency_example", "Category_example")}}) // CreateCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCheckout(context.Background()).CreateCheckoutRequest(createCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCheckout`: CreateCheckoutResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCheckout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckoutRequest** | [**CreateCheckoutRequest**](CreateCheckoutRequest.md) |  | 

### Return type

[**CreateCheckoutResponse**](CreateCheckoutResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckout

> GetCheckoutResponse GetCheckout(ctx, checkoutUuid).MerchantOrderId(merchantOrderId).Execute()

Get checkout details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	checkoutUuid := "checkoutUuid_example" // string | 
	merchantOrderId := "merchantOrderId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCheckout(context.Background(), checkoutUuid).MerchantOrderId(merchantOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckout`: GetCheckoutResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCheckout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantOrderId** | **string** |  | 

### Return type

[**GetCheckoutResponse**](GetCheckoutResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealth

> HealthResponse GetHealth(ctx).Execute()

Health check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealth`: HealthResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthToken

> OAuthTokenResponse GetOAuthToken(ctx).OAuthTokenRequest(oAuthTokenRequest).Execute()

Issue an access token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	oAuthTokenRequest := *openapiclient.NewOAuthTokenRequest("GrantType_example", "ClientId_example") // OAuthTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOAuthToken(context.Background()).OAuthTokenRequest(oAuthTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuthToken`: OAuthTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOAuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuthTokenRequest** | [**OAuthTokenRequest**](OAuthTokenRequest.md) |  | 

### Return type

[**OAuthTokenResponse**](OAuthTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettlement

> GetSettlementResponse GetSettlement(ctx, uuid).Execute()

Get a settlement by UUID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	uuid := "uuid_example" // string | Settlement UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSettlement(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSettlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettlement`: GetSettlementResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSettlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Settlement UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSettlementResponse**](GetSettlementResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSettlements

> ListSettlementsResponse ListSettlements(ctx).Offset(offset).Limit(limit).SortDir(sortDir).SortBy(sortBy).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Status(status).Execute()

List settlements for this merchant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	offset := TODO // interface{} | Pagination offset (default 0) (optional)
	limit := TODO // interface{} | Max records to return (default 50, max 200) (optional)
	sortDir := "sortDir_example" // string | Sort direction (default: asc) (optional)
	sortBy := "sortBy_example" // string | Field to sort by (default: created_at) (optional)
	createdBefore := TODO // interface{} | Include settlements created at or before this ISO-8601 timestamp (optional)
	createdAfter := TODO // interface{} | Include settlements created at or after this ISO-8601 timestamp (optional)
	status := "status_example" // string | Filter by status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListSettlements(context.Background()).Offset(offset).Limit(limit).SortDir(sortDir).SortBy(sortBy).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSettlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSettlements`: ListSettlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSettlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSettlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | [**interface{}**](interface{}.md) | Pagination offset (default 0) | 
 **limit** | [**interface{}**](interface{}.md) | Max records to return (default 50, max 200) | 
 **sortDir** | **string** | Sort direction (default: asc) | 
 **sortBy** | **string** | Field to sort by (default: created_at) | 
 **createdBefore** | [**interface{}**](interface{}.md) | Include settlements created at or before this ISO-8601 timestamp | 
 **createdAfter** | [**interface{}**](interface{}.md) | Include settlements created at or after this ISO-8601 timestamp | 
 **status** | **string** | Filter by status | 

### Return type

[**ListSettlementsResponse**](ListSettlementsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundCheckout

> RefundCheckout(ctx, checkoutUuid).CreateRefundRequest(createRefundRequest).Execute()

Refund a captured checkout



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kttipay/maytes-api-client-go"
)

func main() {
	checkoutUuid := "checkoutUuid_example" // string | 
	createRefundRequest := *openapiclient.NewCreateRefundRequest("MerchantRefundRef_example") // CreateRefundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RefundCheckout(context.Background(), checkoutUuid).CreateRefundRequest(createRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RefundCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRefundRequest** | [**CreateRefundRequest**](CreateRefundRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

