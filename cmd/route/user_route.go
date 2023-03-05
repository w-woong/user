package route

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/middlewares"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/user/delivery"
	"github.com/w-woong/user/port"
)

func UserRoute(router *mux.Router, conf common.ConfigHttp, cookie commonport.Cookie,
	idTokenParser commonport.IDTokenParser, usc port.UserUsc) *delivery.UserHttpHandler {

	handler := delivery.NewUserHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	router.HandleFunc("/v1/user/account",
		middlewares.AuthIDToken(handler.HandleFindByLoginID, cookie, idTokenParser),
	).Methods(http.MethodGet)

	router.HandleFunc("/v1/user",
		middlewares.AuthBearerHandler(handler.HandleRegisterUser, conf.BearerToken),
	).Methods(http.MethodPost)

	router.HandleFunc("/v1/user/{id}",
		middlewares.AuthIDToken(handler.HandleFindUser, cookie, idTokenParser),
	).Methods(http.MethodGet)

	router.HandleFunc("/v1/user/{id}",
		middlewares.AuthJWTHandler(handler.HandleRemoveUser, conf.Jwt.Secret),
	).Methods(http.MethodDelete)

	return handler
}
