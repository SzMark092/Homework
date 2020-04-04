# IO.Swagger.Api.SQLWebHandlerApi

All URIs are relative to *http://petstore.swagger.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTable**](SQLWebHandlerApi.md#createtable) | **POST** /CreateTable | Create a table with the given code.
[**GetDataPointDescriptionTable**](SQLWebHandlerApi.md#getdatapointdescriptiontable) | **GET** /GetDataPointDescriptionTable | Get the specified table from the SQL server.
[**GetDataPointTable**](SQLWebHandlerApi.md#getdatapointtable) | **GET** /GetDataPointTable | Get the specified table from the SQL server.
[**GetModule**](SQLWebHandlerApi.md#getmodule) | **GET** /GetModuleTable | Get the specified table from the SQL server.


<a name="createtable"></a>
# **CreateTable**
> void CreateTable (int? tableType)

Create a table with the given code.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateTableExample
    {
        public void main()
        {
            var apiInstance = new SQLWebHandlerApi();
            var tableType = 56;  // int? | Type of the table that have to be created.

            try
            {
                // Create a table with the given code.
                apiInstance.CreateTable(tableType);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SQLWebHandlerApi.CreateTable: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableType** | **int?**| Type of the table that have to be created. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getdatapointdescriptiontable"></a>
# **GetDataPointDescriptionTable**
> List<DataPointDescription> GetDataPointDescriptionTable ()

Get the specified table from the SQL server.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetDataPointDescriptionTableExample
    {
        public void main()
        {
            var apiInstance = new SQLWebHandlerApi();

            try
            {
                // Get the specified table from the SQL server.
                List&lt;DataPointDescription&gt; result = apiInstance.GetDataPointDescriptionTable();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SQLWebHandlerApi.GetDataPointDescriptionTable: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<DataPointDescription>**](DataPointDescription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getdatapointtable"></a>
# **GetDataPointTable**
> List<DataPoint> GetDataPointTable (int? tableType)

Get the specified table from the SQL server.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetDataPointTableExample
    {
        public void main()
        {
            var apiInstance = new SQLWebHandlerApi();
            var tableType = 56;  // int? | Type of the table that have to be created.

            try
            {
                // Get the specified table from the SQL server.
                List&lt;DataPoint&gt; result = apiInstance.GetDataPointTable(tableType);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SQLWebHandlerApi.GetDataPointTable: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableType** | **int?**| Type of the table that have to be created. | 

### Return type

[**List<DataPoint>**](DataPoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getmodule"></a>
# **GetModule**
> List<Module> GetModule ()

Get the specified table from the SQL server.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetModuleExample
    {
        public void main()
        {
            var apiInstance = new SQLWebHandlerApi();

            try
            {
                // Get the specified table from the SQL server.
                List&lt;Module&gt; result = apiInstance.GetModule();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SQLWebHandlerApi.GetModule: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<Module>**](Module.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

