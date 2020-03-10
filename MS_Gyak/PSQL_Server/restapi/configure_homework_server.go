// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"PSQL_Server/restapi/operations"
	"PSQL_Server/restapi/operations/create_table"
	"PSQL_Server/restapi/operations/get_table"
	"PSQL_Server/restapi/operations/home_page"
)

//go:generate swagger generate server --target ../src/PSQL_Server --name  --spec ../swagger_server/swagger.yaml --principal SQL_API --skip-models

func configureFlags(api *operations.HomeworkServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HomeworkServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("html producer has not yet been implemented")
	})

	api.CreateTableCreateTableHandler = create_table.CreateTableHandlerFunc(func(params create_table.CreateTableParams) middleware.Responder {
		return middleware.NotImplemented("operation create_table.CreateTable has not yet been implemented")
	})
	api.GetTableGetTableHandler = get_table.GetTableHandlerFunc(func(params get_table.GetTableParams) middleware.Responder {
		return middleware.NotImplemented("operation get_table.GetTable has not yet been implemented")
	})
	api.HomePageGetHomePageHandler = home_page.GetHomePageHandlerFunc(func(params home_page.GetHomePageParams) middleware.Responder {
		return middleware.NotImplemented("operation home_page.GetHomePage has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
