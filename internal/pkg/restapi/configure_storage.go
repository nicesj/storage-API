// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"futuremobile.net/m/v2/internal/pkg/restapi/pkg/operations"
	"futuremobile.net/m/v2/internal/pkg/restapi/pkg/operations/storage"
)

//go:generate swagger generate server --target ../../../internal --name Storage --spec ../../../api/swagger.yaml --api-package pkg/operations --model-package pkg/models --server-package pkg/restapi

func configureFlags(api *operations.StorageAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.StorageAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.ApplicationOctectStreamProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("applicationOctectStream producer has not yet been implemented")
	})
	api.TxtProducer = runtime.TextProducer()

	api.OAuth2Auth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (OAuth2) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.StorageStorageReadFileHandler == nil {
		api.StorageStorageReadFileHandler = storage.StorageReadFileHandlerFunc(func(params storage.StorageReadFileParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageReadFile has not yet been implemented")
		})
	}
	if api.StorageStorageReadFolderHandler == nil {
		api.StorageStorageReadFolderHandler = storage.StorageReadFolderHandlerFunc(func(params storage.StorageReadFolderParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageReadFolder has not yet been implemented")
		})
	}
	if api.StorageStorageWriteFileHandler == nil {
		api.StorageStorageWriteFileHandler = storage.StorageWriteFileHandlerFunc(func(params storage.StorageWriteFileParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageWriteFile has not yet been implemented")
		})
	}

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
