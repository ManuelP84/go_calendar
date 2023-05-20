package bus

type TaskBus interface {
	CloseConnection()
	CloseChannel()
	DeclareExchange() error
	Publish(mssge string) error
}
