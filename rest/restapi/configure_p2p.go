// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/subutai-io/p2p/rest/restapi/operations"
	"github.com/subutai-io/p2p/rest/restapi/operations/daemon"
	"github.com/subutai-io/p2p/rest/restapi/operations/instances"
	"github.com/subutai-io/p2p/rest/restapi/operations/swarm"
)

//go:generate swagger generate server --target .. --name p2pApi --spec ../swagger.yml

func configureFlags(api *operations.P2pAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.P2pAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.XMLProducer = runtime.XMLProducer()

	api.InstancesCloseInstanceHandler = instances.CloseInstanceHandlerFunc(func(params instances.CloseInstanceParams) middleware.Responder {
		return middleware.NotImplemented("operation instances.CloseInstance has not yet been implemented")
	})
	api.InstancesCreateInstanceHandler = instances.CreateInstanceHandlerFunc(func(params instances.CreateInstanceParams) middleware.Responder {
		return middleware.NotImplemented("operation instances.CreateInstance has not yet been implemented")
	})
	api.DaemonDaemonInfoHandler = daemon.DaemonInfoHandlerFunc(func(params daemon.DaemonInfoParams) middleware.Responder {
		return middleware.NotImplemented("operation daemon.DaemonInfo has not yet been implemented")
	})
	api.DaemonDaemonOptionsHandler = daemon.DaemonOptionsHandlerFunc(func(params daemon.DaemonOptionsParams) middleware.Responder {
		return middleware.NotImplemented("operation daemon.DaemonOptions has not yet been implemented")
	})
	api.InstancesListInstancesHandler = instances.ListInstancesHandlerFunc(func(params instances.ListInstancesParams) middleware.Responder {
		return middleware.NotImplemented("operation instances.ListInstances has not yet been implemented")
	})
	api.SwarmSwarmOptionsHandler = swarm.SwarmOptionsHandlerFunc(func(params swarm.SwarmOptionsParams) middleware.Responder {
		return middleware.NotImplemented("operation swarm.SwarmOptions has not yet been implemented")
	})
	api.SwarmSwarmStatusHandler = swarm.SwarmStatusHandlerFunc(func(params swarm.SwarmStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation swarm.SwarmStatus has not yet been implemented")
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
func configureServer(s *graceful.Server, scheme, addr string) {
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
