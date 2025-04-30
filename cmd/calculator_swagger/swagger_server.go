package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
	"path/filepath"
	"runtime"
)

// @title Calculator API
// @version 1.0
// @description API for performing calculations with the Calculator service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /

func SetViperConfig() {
	viper.SetEnvPrefix("swg")
	viper.SetDefault("port", "8082") // Different port from the main service
	err := viper.BindEnv("port")
	if err != nil {
		slog.Info("Used default port", "port", viper.GetString("port"))
	}
	err = viper.BindEnv("swaggerFile")
	if err != nil {
		slog.Info("Used default swaggerFile", "swaggerFile", viper.GetString("swaggerFile"))
	}
	viper.AutomaticEnv()
}

// getSwaggerDocs returns the location of the Swagger JSON file
func getSwaggerDocs() string {
	if viper.GetString("swaggerFile") != "" {
		return viper.GetString("swaggerFile")
	}
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(filename), "swagger.json")
}

// swaggerUIHTML generates the Swagger UI HTML
func swaggerUIHTML() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Calculator API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css">
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js"></script>
    <script>
        window.onload = function() {
            window.ui = SwaggerUIBundle({
                url: "/swagger.json",
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIBundle.SwaggerUIStandalonePreset
                ],
                layout: "BaseLayout",

            });
        };
    </script>
</body>
</html>`
}

func main() {
	SetViperConfig()

	router := mux.NewRouter()

	// Serve Swagger UI
	router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		_, err := w.Write([]byte(swaggerUIHTML()))
		if err != nil {
			slog.Error("Failed to write response", "error", err)
		}
	}).Methods("GET")

	// Serve the Swagger JSON spec
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, getSwaggerDocs())
	}).Methods("GET")

	// Allow CORS
	//handler := cors.Default().Handler(router)

	// Start server
	slog.Info("Swagger UI Server started", "port", viper.GetString("port"))
	slog.Info("Visit http://localhost:" + viper.GetString("port") + "/swagger to see the API documentation")

	err := http.ListenAndServe(":"+viper.GetString("port"), router)
	if err != nil {
		slog.Error("Failed to start swagger server", "error", err)
	}
}
