/**
*
*
* @author 张帆
* @date 2024/06/05 22:02
**/
package cur_userinfo

import (
	"GinSkeletonLearn/app/global/variable"
	"GinSkeletonLearn/app/http/middleware/my_jwt"
	"github.com/gin-gonic/gin"
)

// GetCurrentUserId 获取当前用户的id
// @context 请求上下文
func GetCurrentUserId(context *gin.Context) (int64, bool) {
	tokenKey := variable.ConfigYml.GetString("Token.BindContextKeyName")
	currentUser, exist := context.MustGet(tokenKey).(my_jwt.CustomClaims)
	return currentUser.UserId, exist
}
