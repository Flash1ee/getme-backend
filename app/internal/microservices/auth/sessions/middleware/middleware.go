package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"getme-backend/internal/microservices/auth/delivery/grpc/client"
	"getme-backend/internal/microservices/auth/sessions/usecase"

	"getme-backend/internal/pkg/utilits"
)

type SessionMiddleware struct {
	SessionClient client.AuthCheckerClient
	utilits.LogObject
}

func NewSessionMiddleware(authClient client.AuthCheckerClient, log *logrus.Logger) *SessionMiddleware {
	return &SessionMiddleware{
		SessionClient: authClient,
		LogObject:     utilits.NewLogObject(log),
	}
}

func (m *SessionMiddleware) updateCookie(w http.ResponseWriter, cook *http.Cookie) {
	cook.Expires = time.Now().Add(usecase.ExpiredCookiesTime)
	cook.Path = "/"
	cook.HttpOnly = true
	http.SetCookie(w, cook)
}

func (m *SessionMiddleware) clearCookie(w http.ResponseWriter, cook *http.Cookie) {
	cook.Expires = time.Now().AddDate(0, 0, -1)
	cook.Path = "/"
	cook.HttpOnly = true
	http.SetCookie(w, cook)
}

// CheckFunc Errors:
//		Status 401 "not authorized user"
func (m *SessionMiddleware) CheckFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			m.Log(r).Warnf("in parsing cookie: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		uniqID := sessionID.Value
		if res, err := m.SessionClient.Check(context.Background(), uniqID); err != nil {
			m.Log(r).Warnf("Error in checking session: %v", err)
			m.clearCookie(w, sessionID)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			m.Log(r).Debugf("Get session for user: %d", res.UserID)
			r = r.WithContext(context.WithValue(r.Context(), "user_id", res.UserID))
			r = r.WithContext(context.WithValue(r.Context(), "session_id", res.UniqID))
			m.updateCookie(w, sessionID)
		}
		next.ServeHTTP(w, r)
	})
}

// Check Errors:
//		Status 401 "not authorized user"
func (m *SessionMiddleware) Check(next http.Handler) http.Handler {
	return m.CheckFunc(next)
}

// CheckNotAuthorized Errors:
//		Status 418 "user already authorized"
func (m *SessionMiddleware) CheckNotAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			m.Log(r).Debug("User not Authorized")
			next.ServeHTTP(w, r)
			return
		}

		uniqID := sessionID.Value
		if res, err := m.SessionClient.Check(context.Background(), uniqID); err != nil {
			m.Log(r).Debug("User not Authorized")
			m.clearCookie(w, sessionID)
			next.ServeHTTP(w, r)
			return
		} else {
			m.Log(r).Warnf("UserAuthorized: %d", res.UserID)
			m.updateCookie(w, sessionID)
		}
		w.WriteHeader(http.StatusTeapot)
	})
}

// AddUserIdFunc Errors:
//		Nothing return only add user_id and session_id to context
func (m *SessionMiddleware) AddUserIdFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err == nil {
			uniqID := sessionID.Value
			if res, err := m.SessionClient.Check(context.Background(), uniqID); err == nil {
				m.Log(r).Debugf("Get session for user: %d", res.UserID)
				r = r.WithContext(context.WithValue(r.Context(), "user_id", res.UserID))
				r = r.WithContext(context.WithValue(r.Context(), "session_id", res.UniqID))
				m.updateCookie(w, sessionID)
			} else {
				m.clearCookie(w, sessionID)
			}
		}
		next.ServeHTTP(w, r)
	})
}

//// CheckNotAuthorizedFunc Errors:
////		Status 418 "user already authorized"
//func (m *SessionMiddleware) CheckNotAuthorizedFunc(next hf.HandlerFunc) hf.HandlerFunc {
//	return func(ctx echo.Context) {
//		sessionID, err := ctx.Request().Cookie("session_id")
//		if err != nil {
//			m.Log(ctx.Request()).Debug("User not Authorized")
//			next.ServeHTTP(ctx)
//			return
//		}
//
//		uniqID := sessionID.Value
//		if res, err := m.SessionClient.Check(context.Background(), uniqID); err != nil {
//			m.Log(ctx.Request()).Debug("User not Authorized")
//			m.clearCookie(ctx.Response(), sessionID)
//			next.ServeHTTP(ctx)
//			return
//		} else {
//			m.Log(ctx.Request()).Warnf("UserAuthorized: %d", res.UserID)
//			m.updateCookie(ctx.Response(), sessionID)
//		}
//		ctx.Response().WriteHeader(http.StatusTeapot)
//	}
//}

// AddUserId Errors:
//		Nothing return only add user_id and session_id to context
func (m *SessionMiddleware) AddUserId(next http.Handler) http.Handler {
	return m.AddUserIdFunc(next)
}
