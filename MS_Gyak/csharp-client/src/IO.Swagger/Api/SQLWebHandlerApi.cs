/* 
 * Homework server
 *
 * This is a simple server for my homework. It can make request towards a PSQL database, and get full tables of data from it.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: szmgepesz@gmail.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using RestSharp;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace IO.Swagger.Api
{
    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public interface ISQLWebHandlerApi : IApiAccessor
    {
        #region Synchronous Operations
        /// <summary>
        /// Create a table with the given code.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns></returns>
        void CreateTable (int? tableType);

        /// <summary>
        /// Create a table with the given code.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>ApiResponse of Object(void)</returns>
        ApiResponse<Object> CreateTableWithHttpInfo (int? tableType);
        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>List&lt;DataPointDescription&gt;</returns>
        List<DataPointDescription> GetDataPointDescriptionTable ();

        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>ApiResponse of List&lt;DataPointDescription&gt;</returns>
        ApiResponse<List<DataPointDescription>> GetDataPointDescriptionTableWithHttpInfo ();
        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>List&lt;DataPoint&gt;</returns>
        List<DataPoint> GetDataPointTable (int? tableType);

        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>ApiResponse of List&lt;DataPoint&gt;</returns>
        ApiResponse<List<DataPoint>> GetDataPointTableWithHttpInfo (int? tableType);
        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>List&lt;Module&gt;</returns>
        List<Module> GetModule ();

        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>ApiResponse of List&lt;Module&gt;</returns>
        ApiResponse<List<Module>> GetModuleWithHttpInfo ();
        #endregion Synchronous Operations
        #region Asynchronous Operations
        /// <summary>
        /// Create a table with the given code.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of void</returns>
        System.Threading.Tasks.Task CreateTableAsync (int? tableType);

        /// <summary>
        /// Create a table with the given code.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of ApiResponse</returns>
        System.Threading.Tasks.Task<ApiResponse<Object>> CreateTableAsyncWithHttpInfo (int? tableType);
        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of List&lt;DataPointDescription&gt;</returns>
        System.Threading.Tasks.Task<List<DataPointDescription>> GetDataPointDescriptionTableAsync ();

        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of ApiResponse (List&lt;DataPointDescription&gt;)</returns>
        System.Threading.Tasks.Task<ApiResponse<List<DataPointDescription>>> GetDataPointDescriptionTableAsyncWithHttpInfo ();
        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of List&lt;DataPoint&gt;</returns>
        System.Threading.Tasks.Task<List<DataPoint>> GetDataPointTableAsync (int? tableType);

        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of ApiResponse (List&lt;DataPoint&gt;)</returns>
        System.Threading.Tasks.Task<ApiResponse<List<DataPoint>>> GetDataPointTableAsyncWithHttpInfo (int? tableType);
        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of List&lt;Module&gt;</returns>
        System.Threading.Tasks.Task<List<Module>> GetModuleAsync ();

        /// <summary>
        /// Get the specified table from the SQL server.
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of ApiResponse (List&lt;Module&gt;)</returns>
        System.Threading.Tasks.Task<ApiResponse<List<Module>>> GetModuleAsyncWithHttpInfo ();
        #endregion Asynchronous Operations
    }

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public partial class SQLWebHandlerApi : ISQLWebHandlerApi
    {
        private IO.Swagger.Client.ExceptionFactory _exceptionFactory = (name, response) => null;

        /// <summary>
        /// Initializes a new instance of the <see cref="SQLWebHandlerApi"/> class.
        /// </summary>
        /// <returns></returns>
        public SQLWebHandlerApi(String basePath)
        {
            this.Configuration = new IO.Swagger.Client.Configuration { BasePath = basePath };

            ExceptionFactory = IO.Swagger.Client.Configuration.DefaultExceptionFactory;
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SQLWebHandlerApi"/> class
        /// using Configuration object
        /// </summary>
        /// <param name="configuration">An instance of Configuration</param>
        /// <returns></returns>
        public SQLWebHandlerApi(IO.Swagger.Client.Configuration configuration = null)
        {
            if (configuration == null) // use the default one in Configuration
                this.Configuration = IO.Swagger.Client.Configuration.Default;
            else
                this.Configuration = configuration;

            ExceptionFactory = IO.Swagger.Client.Configuration.DefaultExceptionFactory;
        }

        /// <summary>
        /// Gets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        public String GetBasePath()
        {
            return this.Configuration.ApiClient.RestClient.BaseUrl.ToString();
        }

        /// <summary>
        /// Sets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        [Obsolete("SetBasePath is deprecated, please do 'Configuration.ApiClient = new ApiClient(\"http://new-path\")' instead.")]
        public void SetBasePath(String basePath)
        {
            // do nothing
        }

        /// <summary>
        /// Gets or sets the configuration object
        /// </summary>
        /// <value>An instance of the Configuration</value>
        public IO.Swagger.Client.Configuration Configuration {get; set;}

        /// <summary>
        /// Provides a factory method hook for the creation of exceptions.
        /// </summary>
        public IO.Swagger.Client.ExceptionFactory ExceptionFactory
        {
            get
            {
                if (_exceptionFactory != null && _exceptionFactory.GetInvocationList().Length > 1)
                {
                    throw new InvalidOperationException("Multicast delegate for ExceptionFactory is unsupported.");
                }
                return _exceptionFactory;
            }
            set { _exceptionFactory = value; }
        }

        /// <summary>
        /// Gets the default header.
        /// </summary>
        /// <returns>Dictionary of HTTP header</returns>
        [Obsolete("DefaultHeader is deprecated, please use Configuration.DefaultHeader instead.")]
        public IDictionary<String, String> DefaultHeader()
        {
            return new ReadOnlyDictionary<string, string>(this.Configuration.DefaultHeader);
        }

        /// <summary>
        /// Add default header.
        /// </summary>
        /// <param name="key">Header field name.</param>
        /// <param name="value">Header field value.</param>
        /// <returns></returns>
        [Obsolete("AddDefaultHeader is deprecated, please use Configuration.AddDefaultHeader instead.")]
        public void AddDefaultHeader(string key, string value)
        {
            this.Configuration.AddDefaultHeader(key, value);
        }

        /// <summary>
        /// Create a table with the given code. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns></returns>
        public void CreateTable (int? tableType)
        {
             CreateTableWithHttpInfo(tableType);
        }

        /// <summary>
        /// Create a table with the given code. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>ApiResponse of Object(void)</returns>
        public ApiResponse<Object> CreateTableWithHttpInfo (int? tableType)
        {
            // verify the required parameter 'tableType' is set
            if (tableType == null)
                throw new ApiException(400, "Missing required parameter 'tableType' when calling SQLWebHandlerApi->CreateTable");

            var localVarPath = "/CreateTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
                "application/json"
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            if (tableType != null) localVarQueryParams.AddRange(this.Configuration.ApiClient.ParameterToKeyValuePairs("", "TableType", tableType)); // query parameter


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) this.Configuration.ApiClient.CallApi(localVarPath,
                Method.POST, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("CreateTable", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<Object>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                null);
        }

        /// <summary>
        /// Create a table with the given code. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of void</returns>
        public async System.Threading.Tasks.Task CreateTableAsync (int? tableType)
        {
             await CreateTableAsyncWithHttpInfo(tableType);

        }

        /// <summary>
        /// Create a table with the given code. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of ApiResponse</returns>
        public async System.Threading.Tasks.Task<ApiResponse<Object>> CreateTableAsyncWithHttpInfo (int? tableType)
        {
            // verify the required parameter 'tableType' is set
            if (tableType == null)
                throw new ApiException(400, "Missing required parameter 'tableType' when calling SQLWebHandlerApi->CreateTable");

            var localVarPath = "/CreateTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
                "application/json"
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            if (tableType != null) localVarQueryParams.AddRange(this.Configuration.ApiClient.ParameterToKeyValuePairs("", "TableType", tableType)); // query parameter


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) await this.Configuration.ApiClient.CallApiAsync(localVarPath,
                Method.POST, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("CreateTable", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<Object>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                null);
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>List&lt;DataPointDescription&gt;</returns>
        public List<DataPointDescription> GetDataPointDescriptionTable ()
        {
             ApiResponse<List<DataPointDescription>> localVarResponse = GetDataPointDescriptionTableWithHttpInfo();
             return localVarResponse.Data;
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>ApiResponse of List&lt;DataPointDescription&gt;</returns>
        public ApiResponse< List<DataPointDescription> > GetDataPointDescriptionTableWithHttpInfo ()
        {

            var localVarPath = "/GetDataPointDescriptionTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);



            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) this.Configuration.ApiClient.CallApi(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("GetDataPointDescriptionTable", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<List<DataPointDescription>>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (List<DataPointDescription>) this.Configuration.ApiClient.Deserialize(localVarResponse, typeof(List<DataPointDescription>)));
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of List&lt;DataPointDescription&gt;</returns>
        public async System.Threading.Tasks.Task<List<DataPointDescription>> GetDataPointDescriptionTableAsync ()
        {
             ApiResponse<List<DataPointDescription>> localVarResponse = await GetDataPointDescriptionTableAsyncWithHttpInfo();
             return localVarResponse.Data;

        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of ApiResponse (List&lt;DataPointDescription&gt;)</returns>
        public async System.Threading.Tasks.Task<ApiResponse<List<DataPointDescription>>> GetDataPointDescriptionTableAsyncWithHttpInfo ()
        {

            var localVarPath = "/GetDataPointDescriptionTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);



            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) await this.Configuration.ApiClient.CallApiAsync(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("GetDataPointDescriptionTable", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<List<DataPointDescription>>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (List<DataPointDescription>) this.Configuration.ApiClient.Deserialize(localVarResponse, typeof(List<DataPointDescription>)));
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>List&lt;DataPoint&gt;</returns>
        public List<DataPoint> GetDataPointTable (int? tableType)
        {
             ApiResponse<List<DataPoint>> localVarResponse = GetDataPointTableWithHttpInfo(tableType);
             return localVarResponse.Data;
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>ApiResponse of List&lt;DataPoint&gt;</returns>
        public ApiResponse< List<DataPoint> > GetDataPointTableWithHttpInfo (int? tableType)
        {
            // verify the required parameter 'tableType' is set
            if (tableType == null)
                throw new ApiException(400, "Missing required parameter 'tableType' when calling SQLWebHandlerApi->GetDataPointTable");

            var localVarPath = "/GetDataPointTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            if (tableType != null) localVarQueryParams.AddRange(this.Configuration.ApiClient.ParameterToKeyValuePairs("", "TableType", tableType)); // query parameter


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) this.Configuration.ApiClient.CallApi(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("GetDataPointTable", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<List<DataPoint>>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (List<DataPoint>) this.Configuration.ApiClient.Deserialize(localVarResponse, typeof(List<DataPoint>)));
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of List&lt;DataPoint&gt;</returns>
        public async System.Threading.Tasks.Task<List<DataPoint>> GetDataPointTableAsync (int? tableType)
        {
             ApiResponse<List<DataPoint>> localVarResponse = await GetDataPointTableAsyncWithHttpInfo(tableType);
             return localVarResponse.Data;

        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="tableType">Type of the table that have to be created.</param>
        /// <returns>Task of ApiResponse (List&lt;DataPoint&gt;)</returns>
        public async System.Threading.Tasks.Task<ApiResponse<List<DataPoint>>> GetDataPointTableAsyncWithHttpInfo (int? tableType)
        {
            // verify the required parameter 'tableType' is set
            if (tableType == null)
                throw new ApiException(400, "Missing required parameter 'tableType' when calling SQLWebHandlerApi->GetDataPointTable");

            var localVarPath = "/GetDataPointTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            if (tableType != null) localVarQueryParams.AddRange(this.Configuration.ApiClient.ParameterToKeyValuePairs("", "TableType", tableType)); // query parameter


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) await this.Configuration.ApiClient.CallApiAsync(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("GetDataPointTable", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<List<DataPoint>>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (List<DataPoint>) this.Configuration.ApiClient.Deserialize(localVarResponse, typeof(List<DataPoint>)));
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>List&lt;Module&gt;</returns>
        public List<Module> GetModule ()
        {
             ApiResponse<List<Module>> localVarResponse = GetModuleWithHttpInfo();
             return localVarResponse.Data;
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>ApiResponse of List&lt;Module&gt;</returns>
        public ApiResponse< List<Module> > GetModuleWithHttpInfo ()
        {

            var localVarPath = "/GetModuleTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);



            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) this.Configuration.ApiClient.CallApi(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("GetModule", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<List<Module>>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (List<Module>) this.Configuration.ApiClient.Deserialize(localVarResponse, typeof(List<Module>)));
        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of List&lt;Module&gt;</returns>
        public async System.Threading.Tasks.Task<List<Module>> GetModuleAsync ()
        {
             ApiResponse<List<Module>> localVarResponse = await GetModuleAsyncWithHttpInfo();
             return localVarResponse.Data;

        }

        /// <summary>
        /// Get the specified table from the SQL server. 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <returns>Task of ApiResponse (List&lt;Module&gt;)</returns>
        public async System.Threading.Tasks.Task<ApiResponse<List<Module>>> GetModuleAsyncWithHttpInfo ()
        {

            var localVarPath = "/GetModuleTable";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new List<KeyValuePair<String, String>>();
            var localVarHeaderParams = new Dictionary<String, String>(this.Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
            };
            String localVarHttpContentType = this.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = this.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);



            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) await this.Configuration.ApiClient.CallApiAsync(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("GetModule", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<List<Module>>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (List<Module>) this.Configuration.ApiClient.Deserialize(localVarResponse, typeof(List<Module>)));
        }

    }
}
