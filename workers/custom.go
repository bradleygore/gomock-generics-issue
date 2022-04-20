package workers

import "gomock-generics-issue/iface"

type CustomWorkDetail struct {
	IsFunTimes bool
}

//go:generate mockgen -destination=./mocks/CustomWorker.go -package=mock_iface gomock-generics-issue/workers CustomWorker
/*
	NOTE: This one is using a type defined in _this_ pkg. mockgen generates a mock but it has a problem:

	// DoWork mocks base method.
	func (m *MockCustomWorker) DoWork(arg0 ...interface{}) iface.WorkResult[iface.CustomWorkDetail] {
		m.ctrl.T.Helper()
		varargs := []interface{}{}
		for _, a := range arg0 {
			varargs = append(varargs, a)
		}
		ret := m.ctrl.Call(m, "DoWork", varargs...)
		ret0, _ := ret[0].(iface.WorkResult[iface.CustomWorkDetail])
		return ret0
	}

	notice `iface.WorkResult[iface.CustomWorkDetail]` - while `WorkResult` type is in `iface` pkg, making `iface.WorkResult` correct,
	what is not correct is the generic parameter as `CustomWorkDetail` is not in the `iface` pkg.
*/
type CustomWorker interface {
	iface.Worker[CustomWorkDetail]
	HaveGoodTimes()
}
