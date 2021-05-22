package broker

type Handler interface {
	Handle([]byte) error
}
