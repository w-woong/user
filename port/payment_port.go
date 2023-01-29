package port

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/user/entity"
)

type PaymentTypeRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.PaymentType) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.PaymentType, error)
	ReadNoTx(ctx context.Context, id string) (entity.PaymentType, error)
	Update(ctx context.Context, tx common.TxController, o entity.PaymentType) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
}

type PaymentMethodRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.PaymentMethod) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.PaymentMethod, error)
	ReadNoTx(ctx context.Context, id string) (entity.PaymentMethod, error)
	Update(ctx context.Context, tx common.TxController, o entity.PaymentMethod) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
}
