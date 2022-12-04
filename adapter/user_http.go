package adapter

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-wonk/si/sicore"
	"github.com/go-wonk/si/sihttp"
	"github.com/w-woong/common"
	"github.com/w-woong/user/dto"
)

type userHttp struct {
	client      *sihttp.Client
	loginSource string
	baseUrl     string
	host        string
	bearerToken string
}

func NewUserHttp(client *http.Client, loginSource string, baseUrl string, bearerToken string) *userHttp {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"

	c := sihttp.NewClient(client, sihttp.WithBaseUrl(baseUrl),
		sihttp.WithRequestOpt(sihttp.WithBearerToken(bearerToken)),
		sihttp.WithWriterOpt(sicore.SetJsonEncoder()),
		sihttp.WithReaderOpt(sicore.SetJsonDecoder()),
		sihttp.WithDefaultHeaders(headers))

	a := &userHttp{
		client:      c,
		loginSource: loginSource,
		baseUrl:     baseUrl,
		bearerToken: bearerToken,
	}
	if u, err := url.Parse(baseUrl); err == nil {
		a.host = u.Host
	}
	return a
}

func (a *userHttp) RegisterUser(ctx context.Context, user dto.User) (dto.User, error) {

	req := common.HttpBody{
		Count:    1,
		Document: &user,
	}

	resUser := dto.User{}
	res := common.HttpBody{
		Document: &resUser,
	}
	err := a.client.RequestPostDecodeContext(ctx, "/v1/user/"+a.loginSource, nil, &req, &res)
	if err != nil {
		return dto.NilUser, err
	}

	if res.Status != http.StatusOK {
		return resUser, errors.New(res.Message)
	}

	return resUser, nil
}
