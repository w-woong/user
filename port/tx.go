package port

//go:generate mockgen -destination=./mocks/mock_tx.go -package=mocks -mock_names=TxBeginner=MockTxBeginner,TxController=MockTxController -source=./tx.go . TxBeginner,TxController

type TxBeginner interface {
	Begin() (TxController, error)
}

type TxController interface {
	Commit() error
	Rollback() error
}
