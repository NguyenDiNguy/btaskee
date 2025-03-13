package main

import (
	"btaskee/libs/logger"
	"btaskee/libs/viper"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	urlHTTP := ":" + viper.GlobalConfig.PortHTTP

	fmt.Println("...Start Swagger...")

	logger.Init()
	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Continue with the request
			next.ServeHTTP(w, r)
		})
	})

	docs.SwaggerInfo.Host = viper.GlobalConfig.Endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(viper.GlobalConfig.Endpoint+"/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	http.ListenAndServe(urlHTTP, r)
}
