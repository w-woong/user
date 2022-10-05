package delivery

import (
	"log"
	"net/http"

	"github.com/go-wonk/si"
	"github.com/gorilla/mux"
	"github.com/w-woong/user/pkg/dto"
	"github.com/w-woong/user/pkg/port"
)

type UserHttpHandler struct {
	userUsc port.UserUsc
}

func NewUserHttpHandler(findUserUsc port.UserUsc) *UserHttpHandler {
	return &UserHttpHandler{
		userUsc: findUserUsc,
	}
}

func (d *UserHttpHandler) HandleFindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	user, err := d.userUsc.FindUserByID(ID)
	if err != nil {
		http.Error(w, "user does not exist", http.StatusInternalServerError)
		return
	}

	if err = si.EncodeJson(w, &user); err != nil {
		log.Println(err)
	}
}

func (d *UserHttpHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	reqBody := dto.HttpBody{
		Document: &user,
	}

	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		http.Error(w, "format of request body is unknown", http.StatusBadRequest)
		return
	}

	if err := d.userUsc.RegisterUser(r.Context(), user); err != nil {
		http.Error(w, "could not register user", http.StatusBadRequest)
		return
	}

}

func (d *UserHttpHandler) HandleModifyUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	var user dto.User
	reqBody := dto.HttpBody{
		Document: &user,
	}

	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		http.Error(w, "format of request body is unknown", http.StatusBadRequest)
		return
	}

	if err := d.userUsc.ModifyUser(ID, user); err != nil {
		http.Error(w, "could not register user", http.StatusBadRequest)
		return
	}
}

func (d *UserHttpHandler) HandleRemoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	if err := d.userUsc.RemoveUser(ID); err != nil {
		http.Error(w, "could not register user", http.StatusBadRequest)
		return
	}
}
