package service

import (
	"gitlab.shanhai.int/sre/app-framework"
	"gitlab.shanhai.int/sre/app-framework/internal/example/config"
	"gitlab.shanhai.int/sre/app-framework/internal/example/dao"
	"gitlab.shanhai.int/sre/library/net/httpclient"
)

var (
	// ====================
	// >>>请勿删除<<<
	//
	// 全局服务
	// ====================
	SVC *Service
)

// ====================
// >>>请勿删除<<<
//
// 自定义服务
// ====================
type Service struct {
	// ====================
	// >>>请勿删除<<<
	//
	// 基础服务
	// ====================
	*framework.DefaultService

	// ====================
	// >>>请勿删除<<<
	//
	// 数据层
	// ====================
	dao *dao.Dao

	// ====================
	// 根据实际情况，选择性保留
	// ====================
	// Http客户端
	httpClient *httpclient.Client
}

// ====================
// >>>请勿删除<<<
//
// 新建服务
// ====================
func New() *Service {
	// ====================
	// >>>请勿删除<<<
	//
	// 新建数据层
	// ====================
	d := dao.New()

	SVC = &Service{
		// ====================
		// >>>请勿删除<<<
		// ====================
		DefaultService: framework.GetDefaultService(config.Conf.Config, d),
		dao:            d,

		// ====================
		// 根据实际情况，选择性保留
		// ====================
		httpClient: httpclient.NewHttpClient(config.Conf.HttpClient),
	}

	return SVC
}
