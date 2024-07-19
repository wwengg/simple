package sbus

type SRouter interface {
	PreHandle(task STask)  //Hook method before processing conn business(在处理conn业务之前的钩子方法)
	Handle(task STask)     //Method for processing conn business(处理conn业务的方法)
	PostHandle(task STask) //Hook method after processing conn business(处理conn业务之后的钩子方法)
}
