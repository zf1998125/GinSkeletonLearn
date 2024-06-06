/**
*
*
* @author 张帆
* @date 2024/06/05 19:22
**/
package files

import (
	"GinSkeletonLearn/app/global/my_errors"
	"GinSkeletonLearn/app/global/variable"
	"mime/multipart"
	"net/http"
	"os"
)

// 通过文件名获取文件mime信息
func GetFilesMimeByFileName(filePath string) string {
	open, err1 := os.Open(filePath)
	defer open.Close()
	if err1 != nil {
		variable.ZapLog.Error(my_errors.ErrorsFilesUploadOpenFail + err1.Error())
	}
	// 只需要前 32 个字节就可以了
	buffer := make([]byte, 32)
	if _, err := open.Read(buffer); err != nil {
		variable.ZapLog.Error(my_errors.ErrorsFilesUploadReadFail + err.Error())
		return ""
	}

	return http.DetectContentType(buffer)
}

func GetFilesMimeByFp(fp multipart.File) string {
	buffer := make([]byte, 32)
	if _, err := fp.Read(buffer); err != nil {
		variable.ZapLog.Error(my_errors.ErrorsFilesUploadReadFail + err.Error())
		return ""
	}

	return http.DetectContentType(buffer)
}
