package main

import (
	framework "gitlab.shanhai.int/sre/app-framework"
	"gitlab.shanhai.int/sre/app-framework/internal/example/config"
	"gitlab.shanhai.int/sre/app-framework/internal/example/server/rpc"
	"gitlab.shanhai.int/sre/app-framework/internal/example/service"
)

// ====================
// >>>请勿删除<<<
//
// rpc服务
// ====================
func main() {
	// 启动服务
	framework.Run(
		// ========================
		// >>>请勿删除<<<
		//
		// 读取配置文件
		//
		//		configPath为配置文件的所属位置，例如 ./config/config.yaml
		// ========================
		config.Read("./config/config-rpc.yaml").Config,

		// ====================
		// >>>请勿删除<<<
		//
		// 新建服务
		// ====================
		service.New(),

		// 启动rpc服务器
		rpc.GetServer(),
	)
}
