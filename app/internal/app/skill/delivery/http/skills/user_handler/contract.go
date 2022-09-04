package skills_user_handler

type Cache interface {
	Get(key string)
	Set(key string)
}
