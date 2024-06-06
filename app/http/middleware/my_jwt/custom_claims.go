/**
*
*
* @author 张帆
* @date 2024/06/05 17:15
**/
package my_jwt

import "github.com/dgrijalva/jwt-go"

// 自定义jwt的声明字段信息+标准字段，参考地址：https://blog.csdn.net/codeSquare/article/details/99288718
type CustomClaims struct {
	jwt.StandardClaims
	UserId int64  `json:"user_id"`
	Name   string `json:"user_name"`
	Phone  string `json:"phone"`
}
