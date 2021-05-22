package middleware

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/logger"
	"github.com/lorezi/golang-bank-app/ports"
	"github.com/lorezi/golang-bank-app/utils"
)

type AuthMiddleware struct {
	Repo ports.AuthRepository
}

func (a AuthMiddleware) Authentication() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				logger.Error("No authentication header provided")

				err := errs.AppError{Code: http.StatusForbidden, Status: "authentication failure", Message: "No authentication header provided"}
				utils.Response(w, err.Code, err.ShowError())
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer") {

				logger.Error("Bearer missing in authentication header")
				err := errs.AppError{Code: http.StatusForbidden, Status: "authentication failure", Message: "Add 'Bearer' to the token 👍🏾👍🏾👍🏾"}
				utils.Response(w, err.Code, err.ShowError())
				return

			}

			if strings.HasPrefix(authHeader, "Bearer") {

				token := strings.TrimPrefix(authHeader, "Bearer ")
				isAuthorized := a.Repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)
				if isAuthorized {
					// Call the next handler, which can be another middleware in the chain, or the final handler.
					next.ServeHTTP(w, r)
					return
				}
				err := errs.AppError{Code: http.StatusForbidden, Status: "authentication failure", Message: "invalid token"}
				utils.Response(w, err.Code, err.ShowError())
				return
			}

		})
	}
}
