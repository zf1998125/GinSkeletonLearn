/**
*
*
* @author 张帆
* @date 2024/06/04 18:33
**/
package destroy

import (
	"GinSkeletonLearn/app/core/event_manage"
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/global/variable"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//用于系统信号的监听
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) //监听可能退出的信号
		received := <-c                                                                           //接受管道中的值
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("信号值", received.String()))
		event_manage.CreateEventManageFactory().FuzzyCall(variable.EventDestroyPrefix)
		close(c)
		os.Exit(1)
	}()
}
