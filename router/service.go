package router

import (
	"sample/openapi"
	"sample/session"
)

// NewAPI APIのコンストラクタ
func NewAPI(sess session.Session) (*openapi.Api, error) {
	return &openapi.Api{
		Middleware: Middleware{},
		AuthApi: Auth{},
		FavoriteApi: Favorite{},
		FollowApi: Follow{},
		MessageApi: Message{},
		PinApi: Pin{},
		UserApi: User{},
	}, nil
}