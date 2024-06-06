/**
*
*
* @author 张帆
* @date 2024/06/05 11:40
**/
package factory

import (
	"GinSkeletonLearn/app/core/container"
	"GinSkeletonLearn/app/global/my_errors"
	"GinSkeletonLearn/app/global/variable"
	"GinSkeletonLearn/app/http/validator/core/interf"
	"github.com/gin-gonic/gin"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(ctx *gin.Context) {
	if value := container.CreateContainersFactory().Get(key); value != nil {
		if value, ok := value.(interf.ValidatorInterface); ok {
			return value.CheckParams
		}
	}
	variable.ZapLog.Error(my_errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
