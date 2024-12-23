package dao

import (
	"gitlab.shanhai.int/sre/app-framework/internal/example/config"
	"os"
	"testing"
)

var (
	// ====================
	// >>>请勿删除<<<
	//
	// 用于测试环境的数据层变量
	// ====================
	d *Dao
)

// ====================
// >>>请勿删除<<<
//
// 测试前准备工作
//
// 常用于初始化操作
// ====================
func beforeTest() {
	// ====================
	// 设置测试所用env
	//
	// 通常设为 test
	// ====================
	if err := os.Setenv("env", "example-test"); err != nil {
		return
	}
	// 读取配置
	config.Read("../../../config/config.yaml")
	// 新建测试所用的数据层
	d = New()
}

// ====================
// >>>请勿删除<<<
//
// 测试后清理工作
//
// 常用于删除测试数据
// ====================
func afterTest() {
}

// ====================
// >>>请勿删除<<<
//
// 进行任意测试时，都会最先进行的测试主函数
// ====================
func TestMain(m *testing.M) {
	// 测试前准备工作
	beforeTest()
	// 进行测试
	code := m.Run()
	// 测试后清洗
	afterTest()
	// 退出
	os.Exit(code)
}