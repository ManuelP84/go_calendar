package bus

//go:generate mockery --name TaskBus
type TaskBus interface {
	Publish(mssge string) error
}
