package auth_usecase

//
//import (
//	"reflect"
//	"testing"
//
//	"github.com/golang/mock/gomock"
//	"github.com/pkg/errors"
//
//	"getme-backend/internal/app/auth/dto"
//	"getme-backend/internal/app/user/entities"
//	mock_user_repository "getme-backend/internal/app/user/repository/mock"
//	"getme-backend/internal/app/user/usecase/user_usecase/mock"
//)
//
//func TestUserUsecase_Auth(t *testing.T) {
//	type args struct {
//		user *dto.AuthUsecase
//	}
//	tests := []struct {
//		name        string
//		repository  func(ctrl *gomock.Controller) *mock_user_repository.MockRepository
//		authChecker func(ctrl *gomock.Controller) *mock.MockauthChecker
//		args        args
//		want        *dto.AuthUsecase
//		wantErr     bool
//	}{
//		{
//			name: "ok auth",
//			repository: func(ctrl *gomock.Controller) *mock_user_repository.MockRepository {
//				m := mock_user_repository.NewMockRepository(ctrl)
//				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&entities.User{
//					Nickname:  "lomodar",
//					FirstName: "Nick",
//					LastName:  "Jackson",
//				}, nil)
//
//				return m
//			},
//			authChecker: func(ctrl *gomock.Controller) *mock.MockauthChecker {
//				m := mock.NewMockauthChecker(ctrl)
//				m.EXPECT().Check(gomock.Any()).Return(true)
//
//				return m
//			},
//			args: args{
//				user: &dto.AuthUsecase{
//					ID:         1,
//					TelegramID: 2,
//					FirstName:  "Nick",
//					LastName:   "Jackson",
//					Username:   "lomodar",
//				},
//			},
//			want: &dto.AuthUsecase{
//				Username:  "lomodar",
//				FirstName: "Nick",
//				LastName:  "Jackson",
//			},
//			wantErr: false,
//		},
//		{
//			name: "false check auth",
//			repository: func(ctrl *gomock.Controller) *mock_user_repository.MockRepository {
//				m := mock_user_repository.NewMockRepository(ctrl)
//
//				return m
//			},
//			authChecker: func(ctrl *gomock.Controller) *mock.MockauthChecker {
//				m := mock.NewMockauthChecker(ctrl)
//				m.EXPECT().Check(gomock.Any()).Return(false)
//
//				return m
//			},
//			args: args{
//				user: &dto.AuthUsecase{
//					ID:         1,
//					TelegramID: 2,
//					FirstName:  "Nick",
//					LastName:   "Jackson",
//					Username:   "lomodar",
//				},
//			},
//			want:    nil,
//			wantErr: true,
//		},
//		{
//			name: "error - empty user",
//			repository: func(ctrl *gomock.Controller) *mock_user_repository.MockRepository {
//				m := mock_user_repository.NewMockRepository(ctrl)
//
//				return m
//			},
//			authChecker: func(ctrl *gomock.Controller) *mock.MockauthChecker {
//				m := mock.NewMockauthChecker(ctrl)
//
//				return m
//			},
//			args: args{
//				user: nil,
//			},
//			want:    nil,
//			wantErr: true,
//		},
//		{
//			name: "error - repository error",
//			repository: func(ctrl *gomock.Controller) *mock_user_repository.MockRepository {
//				m := mock_user_repository.NewMockRepository(ctrl)
//				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
//
//				return m
//			},
//			authChecker: func(ctrl *gomock.Controller) *mock.MockauthChecker {
//				m := mock.NewMockauthChecker(ctrl)
//				m.EXPECT().Check(gomock.Any()).Return(true)
//
//				return m
//			},
//			args: args{
//				user: &dto.AuthUsecase{
//					ID:         1,
//					TelegramID: 2,
//					FirstName:  "Nick",
//					LastName:   "Jackson",
//					Username:   "lomodar",
//				},
//			},
//			want:    nil,
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//
//			u := NewUserUsecase(tt.repository(ctrl), tt.authChecker(ctrl))
//
//			got, err := u.AuthTelegram(tt.args.user)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("AuthTelegram() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("AuthTelegram() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
