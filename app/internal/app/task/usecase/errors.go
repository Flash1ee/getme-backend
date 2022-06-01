package task_usecase

import "github.com/pkg/errors"

var (
	UserHaveNotThisPlan = errors.New("this user have not plan with this id")
	UserHaveNotThisTask = errors.New("this user have not task with this id")
)
