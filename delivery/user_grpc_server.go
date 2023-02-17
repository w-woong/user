package delivery

import (
	"context"
	"errors"
	"net/http"
	"time"

	commonconv "github.com/w-woong/common/conv"
	commondto "github.com/w-woong/common/dto"
	pb "github.com/w-woong/common/dto/protos/user/v2"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/user/port"
)

type userGrpcServer struct {
	pb.UserServiceServer

	timeout time.Duration
	usc     port.UserUsc
}

func NewUserGrpcServer(timeout time.Duration, usc port.UserUsc) *userGrpcServer {
	return &userGrpcServer{
		timeout: timeout,
		usc:     usc,
	}
}

func (d *userGrpcServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.UserReply, error) {
	user, err := commonconv.ToUserDtoFromProto(in.GetDocument())
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Debug(user.String())

	registeredUser, err := d.usc.RegisterUser(ctx, user)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	userProtoRes, err := commonconv.ToUserProtoFromDto(registeredUser)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	res := pb.UserReply{
		Status:   http.StatusOK,
		Document: userProtoRes,
	}

	return &res, nil
}

func (d *userGrpcServer) FindByIDToken(ctx context.Context, in *pb.FindByIDTokenRequest) (*pb.UserReply, error) {

	claims, ok := ctx.Value(commondto.IDTokenClaimsContextKey{}).(commondto.IDTokenClaims)
	if !ok {
		logger.Error("could not find claims")
		return nil, errors.New("could not find claims")
	}

	user, err := d.usc.FindByLoginID(ctx, claims.TokenSource, claims.Subject)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	userProtoRes, err := commonconv.ToUserProtoFromDto(user)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	res := pb.UserReply{
		Status:   http.StatusOK,
		Document: userProtoRes,
	}

	return &res, nil
}
