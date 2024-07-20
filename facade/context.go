package facade

type Context[T1 any, T2 any] struct {
	Request T1
	Result T2
}

func (u *Context[T1, T2]) SetRequest(request T1) {
    u.Request = request
}

func (u *Context[T1, T2]) SetResult(result T2) {
    u.Result = result
}