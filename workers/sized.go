package workers

import (
	"gomock-generics-issue/iface"
	"gomock-generics-issue/workdetail"
)

// tried to get it to work with this type for the T, but couldn't, not sure why...

// type BigOrLittleWorkDetail interface {
// 	*workdetail.BigDetail | *workdetail.LittleDetail
// }

// type SizedWorker interface {
// 	iface.Worker[BigOrLittleWorkDetail]
// }

//--------------------------------------------

//go:generate mockgen -destination=./mocks/BigWorker.go -package=mock_iface gomock-generics-issue/workers BigWorker

/*
	NOTE: The above go:generate will fail to run and the output ./mocks/BigWorker.go file will be blank.
	If you run the `mockgen ...` cmd in terminal, you will see the full output of the generated code, which contains this:

	// DoWork mocks base method.
	func (m *MockBigWorker) DoWork(arg0 ...interface{}) iface.WorkResult[*gomock-generics-issue/workdetail.BigDetail] {
        m.ctrl.T.Helper()
        varargs := []interface{}{}
        for _, a := range arg0 {
                varargs = append(varargs, a)
        }
        ret := m.ctrl.Call(m, "DoWork", varargs...)
        ret0, _ := ret[0].(iface.WorkResult[*gomock-generics-issue/workdetail.BigDetail])
        return ret0
	}

	note the `iface.WorkResult[*gomock-generics-issue/workdetail.BigDetail]` portion - this is problematic
	but also not how it behaves differently when the generic parameter passed-in is in same package as usage (see CustomWorker mock output)
	vs when the generic parameter passed-in is in a totally different package.
*/
type BigWorker interface {
	iface.Worker[*workdetail.BigDetail]
}

//go:generate mockgen -destination=./mocks/LittleWorker.go -package=mock_iface gomock-generics-issue/workers LittleWorker
/*
	NOTE: This would have same issue as BigWorker interface above
*/
type LittleWorker interface {
	iface.Worker[*workdetail.LittleDetail]
}
