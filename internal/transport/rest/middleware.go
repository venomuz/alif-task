package rest

import (
	"github.com/rs/cors"
	"github.com/venomuz/alif-task/pkg/middleware"
	"net/http"
)

func GinCorsMiddleware() middleware.Options {
	o := cors.Options{
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
		},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Content-Length",
			"image/png",
			"accept",
			"Accept-Encoding",
			"origin",
			"Cache-Control",
			"X-Requested-With",
			"application/json",
		},
		OptionsPassthrough: false,
		ExposedHeaders: []string{
			"application/json",
			"multipart/form-data",
			"Authorization",
			"application/pdf",
			"video/mp4",
			"Content-Type",
			"image/png",
			"image/jpg",
		},
		Debug:                true,
		OptionsSuccessStatus: 200,
	}

	return o
}
