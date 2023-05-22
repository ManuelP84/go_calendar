package bus

type TaskBus interface {
	Publish(mssge string) error
}
