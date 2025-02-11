package router

import (
	. "github.com/1340691923/ElasticView/api"
	"github.com/1340691923/ElasticView/pkg/api_config"

	. "github.com/gofiber/fiber/v2"
)

// ES GM用户 路由
func runGmUser(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/gm_user"
	gmUser := app.Group(AbsolutePath)
	{
		userController := &UserController{}
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "查询用户详细信息",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "info",
		}, gmUser.(*Group), true, userController.UserInfo)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "GM角色列表",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "roles",
		}, gmUser.(*Group), true, RoleController{}.RolesAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "修改GM角色",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "role/update",
		}, gmUser.(*Group), true, RoleController{}.RolesUpdateAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "新增GM角色",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "role/add",
		}, gmUser.(*Group), true, RoleController{}.RolesAddAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除GM角色",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "role/delete",
		}, gmUser.(*Group), true, RoleController{}.RolesDelAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "查询用户列表",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "userlist",
		}, gmUser.(*Group), true, userController.UserListAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "角色下拉选",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "roleOption",
		}, gmUser.(*Group), false, RoleController{}.RoleOptionAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "通过ID获取用户信息",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "getUserById",
		}, gmUser.(*Group), true, userController.GetUserByIdAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "修改用户信息",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "UpdateUser",
		}, gmUser.(*Group), true, userController.UserUpdateAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "新增用户信息",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "InsertUser",
		}, gmUser.(*Group), true, userController.UserAddAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除用户",
			Method:       api_config.MethodAny,
			AbsolutePath: AbsolutePath,
			RelativePath: "DelUser",
		}, gmUser.(*Group), true, userController.DeleteUserAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取接口路由信息",
			Method:       api_config.MethodGet,
			AbsolutePath: AbsolutePath,
			RelativePath: "UrlConfig",
		}, gmUser.(*Group), false, RbacController{}.UrlConfig)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "退出登录",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "logout",
		}, gmUser.(*Group), false, userController.LogoutAction)

	}
}
