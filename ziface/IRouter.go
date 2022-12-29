package ziface

type IRouter interface {
	// PreHandle 处理Conn业务之前的钩子方法Hook
	PreHandle(request IRequest)
	// Handle 处理Conn业务的主方法Handle
	Handle(request IRequest)
	// PostHandle 处理Conn业务之后的方法Hook
	PostHandle(request IRequest)
}
