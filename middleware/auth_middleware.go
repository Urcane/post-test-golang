package middleware

import (
	"context"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/web"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		unauthorized(writer)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		unauthorized(writer)
		return
	}

	tokenString := parts[1]
	claims, err := helper.ValidateJWT(tokenString)
	if err != nil {
		unauthorized(writer)
		return
	}

	// Add the claims to the request context
	ctx := context.WithValue(request.Context(), "claims", claims)
	request = request.WithContext(ctx)

	middleware.Handler.ServeHTTP(writer, request)
}

func unauthorized(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

//func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	if "RAHASIA" == request.Header.Get("X-API-Key") {
//		// ok
//		middleware.Handler.ServeHTTP(writer, request)
//	} else {
//		// error
//		writer.Header().Set("Content-Type", "application/json")
//		writer.WriteHeader(http.StatusUnauthorized)
//
//		webResponse := web.WebResponse{
//			Code:   http.StatusUnauthorized,
//			Status: "UNAUTHORIZED",
//		}
//
//		helper.WriteToResponseBody(writer, webResponse)
//	}
//}
