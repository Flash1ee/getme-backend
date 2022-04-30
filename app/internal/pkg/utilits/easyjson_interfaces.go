package utilits

import "github.com/mailru/easyjson"

type MarshUnmarsh interface {
	easyjson.Marshaler
	easyjson.Unmarshaler
}
