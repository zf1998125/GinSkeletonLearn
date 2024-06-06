/**
*
*
* @author 张帆
* @date 2024/06/05 19:04
**/
package users

import (
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/http/controller/web"
	"GinSkeletonLearn/app/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type RefreshToken struct {
	Authorization string `json:"token" header:"Authorization" binding:"required,min=20"`
}

func (r RefreshToken) CheckParams(context *gin.Context) {
	err := context.ShouldBind(&r)
	if err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}
	token := strings.Split(r.Authorization, " ")
	if len(token) == 2 {
		context.Set(consts.ValidatorPrefix+"token", token[1])
		(&web.Users{}).RefreshToken(context)
	} else {
		err := gin.H{
			"tips": "Token不合法，token请放置在header头部分，按照按=>键提交，例如：Authorization：Bearer 你的实际token....",
		}
		response.Fail(context, consts.JwtTokenFormatErrCode, consts.JwtTokenFormatErrMsg, err)
	}

}
