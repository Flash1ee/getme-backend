package telegram_auth_handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"getme-backend/internal/pkg/handler"
)

type AuthTestSuite struct {
	handler.SuiteHandler
	handler *AuthHandler
}

func (s *AuthTestSuite) SetupSuite() {
	s.SuiteHandler.SetupSuite()
	s.handler = NewAuthHandler(s.Logger, s.MockServiceSessionClient, s.MockTokenUsecase)
}

func TestUserAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) TestAuthTestSuite_GET() {
	tests := []struct {
		name        string
		queryParams url.Values
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			name:        "validation error, empty params",
			queryParams: url.Values{},
			wantErr:     assert.Error,
		},

		{
			name: "validation error, no token param",
			queryParams: url.Values{
				"token": []string{},
				"oauth": []string{"telegram"},
			},
			wantErr: assert.Error,
		},
		{
			name: "validation error, oauth method not supported",
			queryParams: url.Values{
				"token": []string{},
				"oauth": []string{"vk"},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/auth"+"?"+tt.queryParams.Encode(), nil)

			h := NewAuthHandler(s.Logger, s.MockServiceSessionClient, s.MockTokenUsecase)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			tt.wantErr(t, h.GET(c), fmt.Sprintf("GET(%v)", c))
		})
	}
}
