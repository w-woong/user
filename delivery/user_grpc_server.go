package delivery

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/w-woong/common"
	commonconv "github.com/w-woong/common/conv"
	commondto "github.com/w-woong/common/dto"
	pb "github.com/w-woong/common/dto/protos/user/v1"
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

	registeredUser, err := d.usc.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, common.ErrLoginIDAlreadyExists) && in.GetLoginSource() == "google" {
			registeredUser, err = d.usc.ModifyUser(ctx, user)
			if err != nil {
				logger.Error(err.Error())
				return nil, err
			}
		} else {
			logger.Error(err.Error())
			return nil, err
		}
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

func (d *userGrpcServer) FindByLoginID(ctx context.Context, in *pb.FindByLoginIDRequest) (*pb.UserReply, error) {

	claims, ok := ctx.Value(commondto.IDTokenClaimsKey{}).(commondto.IDTokenClaims)
	if !ok {
		logger.Error("could not find claims")
		return nil, errors.New("could not find claims")
	}

	tokenSource, ok := ctx.Value(commondto.TokenSourceKey{}).(string)
	if !ok {
		logger.Error("could not find claims")
		return nil, errors.New("could not find claims")
	}

	user, err := d.usc.FindByLoginID(ctx, tokenSource, claims.Subject)
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
