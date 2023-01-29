package delivery_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/user/delivery"
	"github.com/w-woong/user/port/mocks"
)

var (
	bd      = time.Date(2002, 1, 2, 0, 0, 0, 0, time.Local)
	userDto = commondto.User{
		ID:      "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
		LoginID: "wonk",
		CredentialPassword: commondto.CredentialPassword{
			ID:     "333cbf79-ca5f-42dc-8ca0-29441209a36a",
			UserID: "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
			Value:  "asdfasdfasdf",
		},
		Personal: commondto.Personal{
			ID:          "433cbf79-ca5f-42dc-8ca0-29441209a36a",
			UserID:      "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
			FirstName:   "wonk",
			LastName:    "sun",
			BirthYear:   2002,
			BirthMonth:  1,
			BirthDay:    2,
			BirthDate:   &bd,
			Gender:      "M",
			Nationality: "KOR",
		},
	}
)

func TestHandleFindUserByID(t *testing.T) {
	urlPath := "/v1/user/{id}"
	ctrl := gomock.NewController(t)
	usc := mocks.NewMockUserUsc(ctrl)
	usc.EXPECT().FindUser(gomock.Any(), userDto.ID).
		Return(userDto, nil).AnyTimes()

	router := mux.NewRouter()
	handler := delivery.NewUserHttpHandler(15*time.Second, usc)
	router.HandleFunc(urlPath, handler.HandleFindUser).Methods(http.MethodGet)

	// request PosVersionHttpHandler
	req, err := http.NewRequest(http.MethodGet, "/v1/user/22bcbf79-ca5f-42dc-8ca0-29441209a36a", nil)
	assert.Nil(t, err)

	// response 받을 레코더 초기화
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	resBody, err := io.ReadAll(rr.Body)
	assert.Nil(t, err)
	fmt.Println(string(resBody))
	assert.Equal(t, http.StatusOK, rr.Code)
}
