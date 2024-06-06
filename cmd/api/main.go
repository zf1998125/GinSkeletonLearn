/**
*
*
* @author 张帆
* @date 2024/06/06 11:07
**/
package api

import "GinSkeletonLearn/app/global/variable"

// 这里可以存放门户类网站入口
func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
