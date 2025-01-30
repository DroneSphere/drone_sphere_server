package adapter

// IEngine 是一个引擎接口，定义了引擎的初始化、启动和停止方法。
type IEngine interface {
	// Init 方法用于初始化引擎。
	Init() error
	// Start 方法用于启动引擎。
	Start() error
	// Stop 方法用于停止引擎。
	Stop() error
}
