package workers

import "gomock-generics-issue/iface"

//go:generate mockgen -destination=./mocks/PrimitiveWorker.go -package=mock_iface gomock-generics-issue/workers PrimitiveWorker
/*
	The mock is generated without any issue here 👍
*/
type PrimitiveWorker interface {
	iface.Worker[int]
	BePrimitive()
}
