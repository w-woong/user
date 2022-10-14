package delivery

import (
	"log"
	"net/http"

	"github.com/go-wonk/si"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
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

	var err error
	if err = si.DecodeJson(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		log.Println(err)
		return
	}

	var registeredUser dto.User
	if registeredUser, err = d.userUsc.RegisterUser(r.Context(), user); err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Document: &registeredUser,
	}

	if err = si.EncodeJson(w, &resBody); err != nil {
		log.Println(err)
		return
	}
}

func (d *UserHttpHandler) HandleFindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	user, err := d.userUsc.FindUserByID(ID)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Document: &user,
	}

	if err = si.EncodeJson(w, &resBody); err != nil {
		log.Println(err)
		return
	}
}

func (d *UserHttpHandler) HandleModifyUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	var user dto.User
	reqBody := common.HttpBody{
		Document: &user,
	}

	var err error
	if err = si.DecodeJson(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		log.Println(err)
		return
	}

	if err = d.userUsc.ModifyUser(ID, user); err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err = si.EncodeJson(w, &common.HttpBodyOK); err != nil {
		log.Println(err)
		return
	}
}

func (d *UserHttpHandler) HandleRemoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	var err error
	if err = d.userUsc.RemoveUser(ID); err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err = si.EncodeJson(w, &common.HttpBodyOK); err != nil {
		log.Println(err)
		return
	}
}
