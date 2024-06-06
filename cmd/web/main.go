/**
*
*
* @author 张帆
* @date 2024/06/06 11:07
**/
package web

import "GinSkeletonLearn/app/global/variable"

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	router := routers.InitWebRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}
