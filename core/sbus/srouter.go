package sbus

type SRouter interface {
	PreHandle(task STask)    //Hook method before processing conn business(在处理conn业务之前的钩子方法)
	Handle(task STask) error //Method for processing conn business(处理conn业务的方法)
	PostHandle(task STask)   //Hook method after processing conn business(处理conn业务之后的钩子方法)
}

type BaseRouter struct{}

// Here, all of BaseRouter's methods are empty, because some routers may not want to have PreHandle or PostHandle.
// Therefore, inheriting all routers from BaseRouter has the advantage that PreHandle and PostHandle do not need to be
// implemented to instantiate a router.
// (这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化)

// PreHandle -
func (br *BaseRouter) PreHandle(task STask) {
}

// Handle -
func (br *BaseRouter) Handle(task STask) error {
	return nil
}

// PostHandle -
func (br *BaseRouter) PostHandle(task STask) {
}
