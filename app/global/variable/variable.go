/**
*
*
* @author 张帆
* @date 2024/06/04 15:22
**/
package variable

import (
	"GinSkeletonLearn/app/global/my_errors"
	"GinSkeletonLearn/app/utils/snow_flake/snow_flake_interf"
	"GinSkeletonLearn/app/utils/yml_config/ymlconfig_interf"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var (
	BasePath           string
	EventDestroyPrefix = "Destroy_"
	ConfigKeyPrefix    = "Config_"
	DateFormat         = "2006-01-02 15:04:05"

	//全局日志指针
	ZapLog *zap.Logger
	//全局配置指针

	ConfigYml       ymlconfig_interf.YmlConfigInterf // 全局配置文件指针
	ConfigGormV2Yml ymlconfig_interf.YmlConfigInterf // 全局配置文件指针

	//gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	GormDbMysql      *gorm.DB // 全局gorm的客户端连接
	GormDbSqlserver  *gorm.DB // 全局gorm的客户端连接
	GormDbPostgreSql *gorm.DB // 全局gorm的客户端连接

	SnowFlake snow_flake_interf.InterfaceSnowFlake

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = `{"code":200,"msg":"ws连接成功","data":""}`
	WebsocketServerPingMsg    = "Server->Ping->Client"

	//casbin 全局操作指针
	Enforcer *casbin.SyncedEnforcer

	//  用户自行定义其他全局变量 ↓
)

func init() {
	//1.初始化程序根目录
	if curPath, err := os.Getwd(); err == nil {
		//路径进行处理，兼容单元测试程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
