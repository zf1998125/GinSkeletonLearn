/**
*
*
* @author 张帆
* @date 2024/06/05 14:06
**/
package md5_encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func MD5(param string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// 先base64 然后MD5
func Base64Md5(param string) string {
	return MD5(base64.StdEncoding.EncodeToString([]byte(param)))
}
