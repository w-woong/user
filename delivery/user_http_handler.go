package delivery

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/go-wonk/si"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/user/port"
)

type UserHttpHandler struct {
	timeout time.Duration
	userUsc port.UserUsc
}

func NewUserHttpHandler(timeout time.Duration, findUserUsc port.UserUsc) *UserHttpHandler {
	return &UserHttpHandler{
		timeout: timeout,
		userUsc: findUserUsc,
	}
}

func (d *UserHttpHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user commondto.User
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

	var registeredUser commondto.User
	if registeredUser, err = d.userUsc.RegisterUser(ctx, user); err != nil {
		// TODO: handle "google"
		if errors.Is(err, common.ErrLoginIDAlreadyExists) && user.LoginSource == "google" {
			registeredUser, err = d.userUsc.ModifyUser(ctx, user)
			if err != nil {
				common.HttpError(w, http.StatusInternalServerError)
				logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
				return
			}
		} else {
			common.HttpError(w, http.StatusInternalServerError)
			logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
			return
		}
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

// func (d *UserHttpHandler) HandleRegisterGoogleUser(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	var user dto.User
// 	reqBody := common.HttpBody{
// 		Document: &user,
// 	}

// 	var copiedReqBody *bytes.Buffer
// 	var err error
// 	if copiedReqBody, err = si.DecodeJsonCopied(&reqBody, r.Body); err != nil {
// 		common.HttpError(w, http.StatusBadRequest)
// 		logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
// 		return
// 	}

// 	var registeredUser dto.User
// 	registeredUser, err = d.userUsc.RegisterGoogleUser(ctx, user)
// 	if err != nil {
// 		if errors.Is(err, common.ErrLoginIDAlreadyExists) {
// 			registeredUser, err = d.userUsc.ModifyGoogleUser(ctx, user)
// 			if err != nil {
// 				common.HttpError(w, http.StatusInternalServerError)
// 				logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
// 				return
// 			}
// 		} else {
// 			common.HttpError(w, http.StatusInternalServerError)
// 			logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
// 			return
// 		}
// 	}

// 	resBody := common.HttpBody{
// 		Status:   http.StatusOK,
// 		Document: &registeredUser,
// 	}

// 	var copiedResBody *bytes.Buffer
// 	if copiedResBody, err = si.EncodeJsonCopied(w, &resBody); err != nil {
// 		logger.Error(err.Error(), logger.UrlField(r.URL.String()),
// 			logger.ReqBodyField(copiedReqBody.Bytes()),
// 			logger.ResBodyField(copiedResBody.Bytes()))
// 		return
// 	}
// }

func (d *UserHttpHandler) HandleFindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	user, err := d.userUsc.FindUser(r.Context(), ID)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
		Document: &user,
	}

	if err = si.EncodeJson(w, &resBody); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *UserHttpHandler) HandleFindByLoginID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims, ok := ctx.Value(commondto.IDTokenClaimsKey{}).(commondto.IDTokenClaims)
	if !ok {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("could not find claims", logger.UrlField(r.URL.String()))
		return
	}

	tokenSource, ok := ctx.Value(commondto.TokenSourceKey{}).(string)
	if !ok {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("could not find claims", logger.UrlField(r.URL.String()))
		return
	}

	user, err := d.userUsc.FindByLoginID(ctx, tokenSource, claims.Subject)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
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
	if err = d.userUsc.RemoveUser(r.Context(), ID); err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if err = si.EncodeJson(w, &common.HttpBodyOK); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}
