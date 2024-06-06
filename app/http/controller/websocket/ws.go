/**
*
*
* @author 张帆
* @date 2024/06/04 20:16
**/
package websocket

import (
	"GinSkeletonLearn/app/service/websocket"
	"github.com/gin-gonic/gin"
)

type Ws struct{}

// onOpen主要解决握手+协议升级
func (ws *Ws) OnOpen(ctx *gin.Context) (*websocket.Ws, bool) {
	return (&websocket.Ws{}).OnOpen(ctx)
}

// OnMessage 处理业务消息
func (ws *Ws) OnMessage(serviceWs *websocket.Ws, ctx *gin.Context) {
	serviceWs.OnMessage(ctx)
}
