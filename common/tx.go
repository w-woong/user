package common

type TxBeginner interface {
	Begin() (TxController, error)
}

type TxController interface {
	Commit() error
	Rollback() error
}
