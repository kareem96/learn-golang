package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware{
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// success
		middleware.Handler.ServeHTTP(writer, request)
		
	}else{
		// error
		writer.Header().Set("Content-Type", "application/json")

		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "Unauthorized",
			}

		helper.WriteToResponseBody(writer, webResponse)
	}
}