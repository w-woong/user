package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/user/delivery"
)

var (
	userHandler *delivery.UserHttpHandler
)

func SetRoute(router *mux.Router, conf ConfigHttp) {
	router.HandleFunc("/v1/user/register",
		common.AuthBearerHandler(userHandler.HandleRegisterUser, conf.BearerToken),
	).Methods(http.MethodPost)

	router.HandleFunc("/v1/user/{id}",
		common.AuthBearerHandler(userHandler.HandleFindByID, conf.BearerToken),
	).Methods(http.MethodGet)

	// router.HandleFunc("/v1/user/{id}",
	// 	common.AuthJWTHandler(userHandler.HandleChangeUser, conf.Jwt.Secret),
	// ).Methods(http.MethodPut)

	router.HandleFunc("/v1/user/{id}",
		common.AuthJWTHandler(userHandler.HandleRemoveUser, conf.Jwt.Secret),
	).Methods(http.MethodDelete)

}
