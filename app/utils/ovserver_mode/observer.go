/**
*
*
* @author 张帆
* @date 2024/06/06 10:49
**/
package ovserver_mode

// 观察者角色（Observer）接口
type ObserverInterface interface {
	// 接收状态更新消息
	Update(*Subject)
}
