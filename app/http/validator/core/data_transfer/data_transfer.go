/**
*
*
* @author 张帆
* @date 2024/06/05 11:27
**/
package data_transfer

import (
	"GinSkeletonLearn/app/global/variable"
	"GinSkeletonLearn/app/http/validator/core/interf"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func DataAddContext(validatorInterface interf.ValidatorInterface, extraAddDataPrefix string, context *gin.Context) *gin.Context {
	var tempJson interface{}
	if tmpBytes, err1 := json.Marshal(validatorInterface); err1 == nil {
		if err2 := json.Unmarshal(tmpBytes, &tempJson); err2 == nil {
			if value, ok := tempJson.(map[string]interface{}); ok {
				for key, val := range value {
					context.Set(extraAddDataPrefix+key, val)

				}
				// 此外给上下文追加三个键：created_at  、 updated_at  、 deleted_at ，实际根据需要自己选择获取相关键值
				curDateTime := time.Now().Format(variable.DateFormat)
				context.Set(extraAddDataPrefix+"created_at", curDateTime)
				context.Set(extraAddDataPrefix+"updated_at", curDateTime)
				context.Set(extraAddDataPrefix+"deleted_at", curDateTime)
				return context
			}
		}
	}
	return nil
}
