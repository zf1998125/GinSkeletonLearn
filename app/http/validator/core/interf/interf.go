/**
*
*
* @author 张帆
* @date 2024/06/05 11:29
**/
package interf

import "github.com/gin-gonic/gin"

// 验证器接口，每个验证器必须实现该接口，请勿修改
type ValidatorInterface interface {
	CheckParams(context *gin.Context)
}
