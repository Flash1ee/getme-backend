package middleware

import (
	"context"
	"net/http"
	"strings"

	dto_token "getme-backend/internal/app/token/dto"
	token_usecase "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/pkg/utilits"

	"github.com/sirupsen/logrus"
)

type JWTMiddleware struct {
	usecase token_usecase.Usecase
	utilits.LogObject
}

func NewJWTMiddleware(log *logrus.Logger, uc token_usecase.Usecase) *JWTMiddleware {
	return &JWTMiddleware{
		LogObject: utilits.NewLogObject(log),
		usecase:   uc,
	}
}

func (mw *JWTMiddleware) CheckJWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearer := strings.Split(authHeader, "Bearer")
		if len(bearer) < 2 {
			mw.Log(r).Infof("jwt-auth header is empty")
			next.ServeHTTP(w, r)
		}
		//userId, okUser := r.Context().Value("user_id").(int64)
		//if !okUser {
		//	mw.Log(r).Info("can not get user_id form context")
		//	mw.Log(r).Infof("user_id = %v", userId)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		if len(authHeader) == 0 {
			mw.Log(r).Infof("jwt-auth header is empty")
			//mw.Log(r).Infof("userId: %v", userId)
			//w.WriteHeader(http.StatusForbidden)
			next.ServeHTTP(w, r)
		}
		tokenSources := dto_token.TokenSourcesUsecase{
			IdentifierData: r.RemoteAddr,
		}
		tokenDTO := dto_token.TokenUsecase{
			Token: bearer[1],
		}
		if err := mw.usecase.Check(tokenSources, tokenDTO); err != nil {
			mw.Log(r).Infof("jwt-auth middleware. JWT token expired or not valid. err: %v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		okAuth := context.WithValue(context.Background(), "JWT", true)
		r = r.WithContext(okAuth)
		next.ServeHTTP(w, r)
	})
}

// CheckNotAuthorized Errors:
//		Status 418 "user already authorized"
func (mw *JWTMiddleware) CheckNotAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		bearer := strings.Split(token, "Bearer")
		if len(bearer) < 2 {
			mw.Log(r).Infof("jwt-auth header is empty")
			next.ServeHTTP(w, r)
			return
		}
		tokenSources := dto_token.TokenSourcesUsecase{
			IdentifierData: r.RemoteAddr,
		}
		tokenDTO := dto_token.TokenUsecase{
			Token: bearer[1],
		}
		if err := mw.usecase.Check(tokenSources, tokenDTO); err != nil {
			mw.Log(r).Debug("User not Authorized")
			next.ServeHTTP(w, r)
			return
		}
		mw.Log(r).Debug("User Authorized")
		w.WriteHeader(http.StatusTeapot)
	})
}

//func (mw *JWTMiddleware) CheckCsrfToken(next http.Handler) http.Handler {
//	return http.HandlerFunc(mw.CheckJWTAuth(next.ServeHTTP))
//}
