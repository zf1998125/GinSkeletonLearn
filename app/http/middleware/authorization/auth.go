/**
*
*
* @author 张帆
* @date 2024/06/05 20:36
**/
package authorization

import (
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/global/variable"
	token2 "GinSkeletonLearn/app/service/users/token"
	"GinSkeletonLearn/app/utils/response"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

func CheckTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerParams := HeaderParams{}

		if err := c.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(c, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			tokenIsEffective := token2.CreateUserTokenFactory().IsEffective(token[1])
			if tokenIsEffective {
				if customToken, err := token2.CreateUserTokenFactory().ParseToken(token[1]); err == nil {
					key := variable.ConfigYml.GetString("Token.BindContextKeyName")
					// token验证通过，同时绑定在请求上下文
					c.Set(key, customToken)
				}
				c.Next()
			} else {
				response.ErrorTokenAuthFail(c)
			}
		} else {
			response.ErrorTokenBaseInfo(c)
		}
	}
}

// CheckTokenAuthWithRefresh 检查token完整性、有效性并且自动刷新中间件
func CheckTokenAuthWithRefresh() gin.HandlerFunc {
	return func(context *gin.Context) {

		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			tokenIsEffective := token2.CreateUserTokenFactory().IsEffective(token[1])
			// 判断token是否有效
			if tokenIsEffective {
				if customToken, err := token2.CreateUserTokenFactory().ParseToken(token[1]); err == nil {
					key := variable.ConfigYml.GetString("Token.BindContextKeyName")
					// token验证通过，同时绑定在请求上下文
					context.Set(key, customToken)
					// 在自动刷新token的中间件中，将请求的认证键、值，原路返回，与后续刷新逻辑格式保持一致
					context.Header("Refresh-Token", "")
					context.Header("Access-Control-Expose-Headers", "Refresh-Token")
				}
				context.Next()
			} else {
				// 判断token是否满足刷新条件
				if token2.CreateUserTokenFactory().TokenIsMeetRefreshCondition(token[1]) {
					// 刷新token
					if newToken, ok := token2.CreateUserTokenFactory().RefreshToken(token[1], context.ClientIP()); ok {
						if customToken, err := token2.CreateUserTokenFactory().ParseToken(newToken); err == nil {
							key := variable.ConfigYml.GetString("Token.BindContextKeyName")
							// token刷新成功，同时绑定在请求上下文
							context.Set(key, customToken)
						}
						// 新token放入header返回
						context.Header("Refresh-Token", newToken)
						context.Header("Access-Control-Expose-Headers", "Refresh-Token")
						context.Next()
					} else {
						response.ErrorTokenRefreshFail(context)
					}
				} else {
					response.ErrorTokenRefreshFail(context)
				}
			}
		} else {
			response.ErrorTokenBaseInfo(context)
		}
	}
}

// RefreshTokenConditionCheck 刷新token条件检查中间件，针对已经过期的token，要求是token格式以及携带的信息满足配置参数即可
func RefreshTokenConditionCheck() gin.HandlerFunc {
	return func(context *gin.Context) {

		headerParams := HeaderParams{}
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			// 判断token是否满足刷新条件
			if token2.CreateUserTokenFactory().TokenIsMeetRefreshCondition(token[1]) {
				context.Next()
			} else {
				response.ErrorTokenRefreshFail(context)
			}
		} else {
			response.ErrorTokenBaseInfo(context)
		}
	}
}

// CheckCasbinAuth casbin检查用户对应的角色权限是否允许访问接口
func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		requstUrl := c.Request.URL.Path
		method := c.Request.Method

		// 模拟请求参数转换后的角色（roleId=2）
		// 主线版本没有深度集成casbin的使用逻辑
		// GinSkeleton-Admin 系统则深度集成了casbin接口权限管控
		// 详细实现参考地址：https://gitee.com/daitougege/gin-skeleton-admin-backend/blob/master/app/http/middleware/authorization/auth.go
		role := "2" // 这里模拟某个用户的roleId=2

		// 这里将用户的id解析为所拥有的的角色，判断是否具有某个权限即可
		isPass, err := variable.Enforcer.Enforce(role, requstUrl, method)
		if err != nil {
			response.ErrorCasbinAuthFail(c, err.Error())
			return
		} else if !isPass {
			response.ErrorCasbinAuthFail(c, "")
			return
		} else {
			c.Next()
		}
	}
}

// CheckCaptchaAuth 验证码中间件
func CheckCaptchaAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		captchaIdKey := variable.ConfigYml.GetString("Captcha.captchaId")
		captchaValueKey := variable.ConfigYml.GetString("Captcha.captchaValue")
		captchaId := c.PostForm(captchaIdKey)
		value := c.PostForm(captchaValueKey)
		if captchaId == "" || value == "" {
			response.Fail(c, consts.CaptchaCheckParamsInvalidCode, consts.CaptchaCheckParamsInvalidMsg, "")
			return
		}
		if captcha.VerifyString(captchaId, value) {
			c.Next()
		} else {
			response.Fail(c, consts.CaptchaCheckFailCode, consts.CaptchaCheckFailMsg, "")
		}
	}
}
