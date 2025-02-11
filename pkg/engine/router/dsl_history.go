package router

import (
	. "github.com/1340691923/ElasticView/api"
	"github.com/1340691923/ElasticView/pkg/api_config"
	. "github.com/gofiber/fiber/v2"
)

// DslHistory 路由
func runDslHistory(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/dslHistory"
	dslHistory := app.Group(AbsolutePath)
	{
		dslHistoryController := &DslHistoryController{}
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "清空DSL查询历史记录",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "CleanAction",
		}, dslHistory.(*Group), true, dslHistoryController.CleanAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "查看DSL查询历史记录",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "ListAction",
		}, dslHistory.(*Group), true, dslHistoryController.ListAction)

	}
}
