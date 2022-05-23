package telegram

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

type UserAuthTestSuite struct {
	handler.SuiteHandler
	handler *UserAuthHandler
}

func (s *UserAuthTestSuite) SetupSuite() {
	s.SuiteHandler.SetupSuite()
	s.handler = NewUserAuthHandler(s.Logger, s.MockUserUsecase, s.MockServiceSessionClient, s.MockTokenUsecase)
}

func TestUserAuthTestSuite(t *testing.T) {
	suite.Run(t, new(UserAuthTestSuite))
}

func (s *UserAuthTestSuite) TestUserAuthHandler_GET() {
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
			//func (v Values) Set(key, value string) {
			//	v[key] = []string{value}
			//}
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/auth"+"?"+tt.queryParams.Encode(), nil)

			h := NewUserAuthHandler(s.Logger, s.MockUserUsecase, s.MockServiceSessionClient, s.MockTokenUsecase)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			tt.wantErr(t, h.GET(c), fmt.Sprintf("GET(%v)", c))
		})
	}
}
