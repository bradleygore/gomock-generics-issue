package iface

type WorkResult[T any] struct {
	Result T
	Error  error
}

//go:generate mockgen -destination=./mocks/Worker.go -package=mock_iface gomock-generics-issue/iface Worker
type Worker[T any] interface {
	DoWork(...interface{}) WorkResult[T]
}
