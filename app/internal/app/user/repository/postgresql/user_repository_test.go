package repository_postgresql

import (
	"context"
	"regexp"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/zhashkevych/go-sqlxmock"

	"getme-backend/internal/app/user/entities"
	user_repository "getme-backend/internal/app/user/repository"
)

func TestUserRepository_Create(t *testing.T) {
	data := getRepositoryData(t)

	type args struct {
		in *entities.User
	}
	type output struct {
		out *entities.User
	}
	tests := map[string]struct {
		prepare  func() (user_repository.Repository, sqlmock.Sqlmock)
		args     args
		out      output
		wantErr  assert.ErrorAssertionFunc
		validate func(args, output)
	}{
		"success_create": {
			prepare: func() (user_repository.Repository, sqlmock.Sqlmock) {
				db, mockDB, _ := sqlmock.Newx()
				st := NewUserRepository(db)

				arg := data

				mockDB.ExpectQuery(regexp.QuoteMeta(queryCreateUser)).
					WithArgs(arg.TelegramID, arg.FirstName, arg.LastName, arg.Nickname, arg.Avatar).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(arg.ID)).
					WillReturnError(nil)

				return st, mockDB
			},
			args: args{
				in: data,
			},
			out: output{
				out: data,
			},
			wantErr: assert.NoError,
			validate: func(args args, output output) {
				assert.Equal(t, args.in.ID, output.out.ID)
			},
		},
		"error_create": {
			prepare: func() (user_repository.Repository, sqlmock.Sqlmock) {
				db, mockDB, _ := sqlmock.Newx()
				st := NewUserRepository(db)

				arg := data

				mockDB.ExpectQuery(regexp.QuoteMeta(queryCreateUser)).
					WithArgs(arg.TelegramID, arg.FirstName, arg.LastName, arg.Nickname, arg.Avatar).
					WillReturnError(errors.New("some error"))

				return st, mockDB
			},
			args: args{
				in: data,
			},
			out: output{
				out: nil,
			},
			wantErr:  assert.Error,
			validate: func(args args, output output) {},
		},
		"success_ID_in_db_no_equal_from_request": {
			prepare: func() (user_repository.Repository, sqlmock.Sqlmock) {
				db, mockDB, _ := sqlmock.Newx()
				st := NewUserRepository(db)

				arg := data

				mockDB.ExpectQuery(regexp.QuoteMeta(queryCreateUser)).
					WithArgs(arg.TelegramID, arg.FirstName, arg.LastName, arg.Nickname, arg.Avatar).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(999)).
					WillReturnError(nil)

				return st, mockDB
			},
			args: args{
				in: data,
			},
			out: output{
				out: &entities.User{
					ID:           999,
					TelegramID:   data.TelegramID,
					FirstName:    data.FirstName,
					LastName:     data.LastName,
					Nickname:     data.Nickname,
					About:        data.About,
					Avatar:       data.Avatar,
					Email:        data.Email,
					IsSearchable: data.IsSearchable,
					CreatedAt:    data.CreatedAt,
					UpdatedAt:    data.UpdatedAt,
				},
			},
			wantErr:  assert.NoError,
			validate: func(args args, output output) {},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			repo, mockDB := tt.prepare()

			var err error
			tt.out.out, err = repo.Create(context.Background(), tt.args.in)

			assert.NoError(t, mockDB.ExpectationsWereMet())
			tt.wantErr(t, err)
			tt.validate(tt.args, tt.out)
		})
	}
}
