package router

import (
	. "github.com/1340691923/ElasticView/api"
	"github.com/1340691923/ElasticView/pkg/api_config"
	. "github.com/gofiber/fiber/v2"
)

// ES文档 路由
func runEsDoc(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/es_doc"
	esDoc := app.Group(AbsolutePath)
	{
		esDocController := &EsDocController{}
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除文档",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "DeleteRowByIDAction",
		}, esDoc.(*Group), true, esDocController.DeleteRowByIDAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "修改文档",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "UpdateByIDAction",
		}, esDoc.(*Group), true, esDocController.UpdateByIDAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "新增文档",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "InsertAction",
		}, esDoc.(*Group), true, esDocController.InsertAction)

	}
}
