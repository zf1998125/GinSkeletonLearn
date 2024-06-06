/**
*
*
* @author 张帆
* @date 2024/06/04 13:27
**/
package users

import (
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/global/variable"
	"github.com/gin-gonic/gin"
)

type DestroyBefore struct {
}

func (d *DestroyBefore) Before(context *gin.Context) bool {
	userId := context.GetFloat64(consts.ValidatorPrefix + "id")
	variable.ZapLog.Sugar().Infof("模拟 Users 删除操作， Before 回调,用户ID：%.f\n", userId)
	if userId > 10 {
		return true
	} else {
		return false
	}
}
