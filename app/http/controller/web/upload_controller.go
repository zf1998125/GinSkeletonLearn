/**
*
*
* @author 张帆
* @date 2024/06/05 16:54
**/
package web

import (
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/global/variable"
	"GinSkeletonLearn/app/service/upload_file"
	"GinSkeletonLearn/app/utils/response"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

//	文件上传是一个独立模块，给任何业务返回文件上传后的存储路径即可。
//
// 开始上传
func (u *Upload) StartUpload(context *gin.Context) {
	savePath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")
	if r, finnalSavePath := upload_file.Upload(context, savePath); r == true {
		response.Success(context, consts.CurdStatusOkMsg, finnalSavePath)
	} else {
		response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
	}
}
