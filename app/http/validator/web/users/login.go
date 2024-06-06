/**
*
*
* @author 张帆
* @date 2024/06/05 18:59
**/
package users

import (
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/http/controller/web"
	"GinSkeletonLearn/app/http/validator/core/data_transfer"
	"GinSkeletonLearn/app/utils/response"
	"github.com/gin-gonic/gin"
)

type Login struct {
	BaseField
}

func (l Login) CheckParams(context *gin.Context) {
	err := context.ShouldBind(&l)
	if err != nil {
		response.ValidatorError(context, err)
		return
	}
	addContext := data_transfer.DataAddContext(l, consts.ValidatorPrefix, context)
	if addContext == nil {
		response.ErrorSystem(context, "userLogin表单验证器json化失败", "")
	}
	(&web.Users{}).Login(addContext)
}
