package delivery

import (
	"bytes"
	"net/http"

	"github.com/go-wonk/si"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/user/dto"
	"github.com/w-woong/user/port"
)

type UserHttpHandler struct {
	userUsc port.UserUsc
}

func NewUserHttpHandler(findUserUsc port.UserUsc) *UserHttpHandler {
	return &UserHttpHandler{
		userUsc: findUserUsc,
	}
}

func (d *UserHttpHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	reqBody := common.HttpBody{
		Document: &user,
	}

	var copiedReqBody *bytes.Buffer
	var err error
	if copiedReqBody, err = si.DecodeJsonCopied(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
		return
	}

	var registeredUser dto.User
	if registeredUser, err = d.userUsc.RegisterUser(r.Context(), user); err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Document: &registeredUser,
	}

	var copiedResBody *bytes.Buffer
	if copiedResBody, err = si.EncodeJsonCopied(w, &resBody); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()),
			logger.ReqBodyField(copiedReqBody.Bytes()),
			logger.ResBodyField(copiedResBody.Bytes()))
		return
	}
}

func (d *UserHttpHandler) HandleFindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	user, err := d.userUsc.FindUserByID(ID)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Document: &user,
	}

	if err = si.EncodeJson(w, &resBody); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *UserHttpHandler) HandleRemoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	var err error
	if err = d.userUsc.RemoveUser(ID); err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if err = si.EncodeJson(w, &common.HttpBodyOK); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}
