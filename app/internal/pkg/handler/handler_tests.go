package handler

import (
	"io"

	"github.com/mailru/easyjson"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	mock_token_usecase "getme-backend/internal/app/token/usecase/mock"
	mock_user_usecase "getme-backend/internal/app/user/usecase/mock"
	mock_session_service "getme-backend/internal/microservices/auth/delivery/grpc/client/mock"
)

type TestTable struct {
	Name              string
	Data              easyjson.Marshaler
	ExpectedMockTimes int
	ExpectedCode      int
}

type SuiteHandler struct {
	suite.Suite
	Mock                     *gomock.Controller
	MockUserUsecase          *mock_user_usecase.MockUsecase
	MockTokenUsecase         *mock_token_usecase.MockUsecase
	MockServiceSessionClient *mock_session_service.MockAuthCheckerClient
	Tb                       TestTable
	Logger                   *logrus.Logger
}

func (s *SuiteHandler) SetupSuite() {
	s.Mock = gomock.NewController(s.T())
	s.MockUserUsecase = mock_user_usecase.NewMockUsecase(s.Mock)
	s.MockTokenUsecase = mock_token_usecase.NewMockUsecase(s.Mock)
	s.MockServiceSessionClient = mock_session_service.NewMockAuthCheckerClient(s.Mock)
	s.Tb = TestTable{}
	s.Logger = logrus.New()
	s.Logger.SetOutput(io.Discard)
}

func (s *SuiteHandler) TearDownSuite() {
	s.Mock.Finish()
}
