# HomeworkServer.SQLWebHandlerApi

All URIs are relative to *http://petstore.swagger.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTable**](SQLWebHandlerApi.md#createTable) | **POST** /CreateTable | Create a table with the given code.
[**getDataPointDescriptionTable**](SQLWebHandlerApi.md#getDataPointDescriptionTable) | **GET** /GetDataPointDescriptionTable | Get the specified table from the SQL server.
[**getDataPointTable**](SQLWebHandlerApi.md#getDataPointTable) | **GET** /GetDataPointTable | Get the specified table from the SQL server.
[**getModule**](SQLWebHandlerApi.md#getModule) | **GET** /GetModuleTable | Get the specified table from the SQL server.


<a name="createTable"></a>
# **createTable**
> createTable(tableType)

Create a table with the given code.



### Example
```javascript
var HomeworkServer = require('homework_server');

var apiInstance = new HomeworkServer.SQLWebHandlerApi();

var tableType = 56; // Number | Type of the table that have to be created.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.createTable(tableType, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableType** | **Number**| Type of the table that have to be created. | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getDataPointDescriptionTable"></a>
# **getDataPointDescriptionTable**
> [DataPointDescription] getDataPointDescriptionTable()

Get the specified table from the SQL server.



### Example
```javascript
var HomeworkServer = require('homework_server');

var apiInstance = new HomeworkServer.SQLWebHandlerApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getDataPointDescriptionTable(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[DataPointDescription]**](DataPointDescription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getDataPointTable"></a>
# **getDataPointTable**
> [DataPoint] getDataPointTable(tableType)

Get the specified table from the SQL server.



### Example
```javascript
var HomeworkServer = require('homework_server');

var apiInstance = new HomeworkServer.SQLWebHandlerApi();

var tableType = 56; // Number | Type of the table that have to be created.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getDataPointTable(tableType, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableType** | **Number**| Type of the table that have to be created. | 

### Return type

[**[DataPoint]**](DataPoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getModule"></a>
# **getModule**
> [Module] getModule()

Get the specified table from the SQL server.



### Example
```javascript
var HomeworkServer = require('homework_server');

var apiInstance = new HomeworkServer.SQLWebHandlerApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getModule(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Module]**](Module.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

