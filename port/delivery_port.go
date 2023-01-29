package port

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/user/entity"
)

type DeliveryRequestTypeRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.DeliveryRequestType) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.DeliveryRequestType, error)
	ReadNoTx(ctx context.Context, id string) (entity.DeliveryRequestType, error)
	Update(ctx context.Context, tx common.TxController, o entity.DeliveryRequestType) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
}

type DeliveryRequestRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.DeliveryRequest) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.DeliveryRequest, error)
	ReadNoTx(ctx context.Context, id string) (entity.DeliveryRequest, error)
	Update(ctx context.Context, tx common.TxController, o entity.DeliveryRequest) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
	DeleteByRequestTypeID(ctx context.Context, tx common.TxController, deliveryRequestTypeID string) (int64, error)
	DeleteByAddressID(ctx context.Context, tx common.TxController, deliveryAddressID string) (int64, error)
}

type DeliveryAddressRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.DeliveryAddress) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.DeliveryAddress, error)
	ReadNoTx(ctx context.Context, id string) (entity.DeliveryAddress, error)
	Update(ctx context.Context, tx common.TxController, o entity.DeliveryAddress) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
	DeleteByUserID(ctx context.Context, tx common.TxController, userID string) (int64, error)
}
