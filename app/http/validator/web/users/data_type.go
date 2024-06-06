/**
*
*
* @author 张帆
* @date 2024/06/05 12:57
**/
package users

type BaseField struct {
	UserName string `json:"user_name" form:"user_name" binding:"required,min=10"`
	Pass     string `form:"pass" json:"pass" binding:"required,min=6,max=20"`
}

type Id struct {
	Id float64 `form:"id"  json:"id" binding:"required,min=1"`
}
