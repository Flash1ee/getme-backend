package skill_usecase

type logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}
