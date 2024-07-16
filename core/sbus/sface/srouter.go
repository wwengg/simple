package sface

type SRouter interface {
	PreHandle(task Stask)  //Hook method before processing conn business(在处理conn业务之前的钩子方法)
	Handle(task Stask)     //Method for processing conn business(处理conn业务的方法)
	PostHandle(task Stask) //Hook method after processing conn business(处理conn业务之后的钩子方法)
}
