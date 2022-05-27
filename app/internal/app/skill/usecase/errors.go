package skill_usecase

import "github.com/pkg/errors"

var (
	SkillNotExists = errors.New("skill not exists in storage")
)
