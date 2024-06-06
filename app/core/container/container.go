/**
*
*
* @author 张帆
* @date 2024/06/04 16:18
**/
package container

import (
	"GinSkeletonLearn/app/global/my_errors"
	"GinSkeletonLearn/app/global/variable"
	"log"
	"strings"
	"sync"
)

var sMap sync.Map

type containers struct {
}

// 创建一个容器工厂 CreateContainersFactory
func CreateContainersFactory() *containers {
	return &containers{}
}

// set 1.以键值对的形式将代码注册到容器中
func (c *containers) Set(key string, value interface{}) (res bool) {
	if _, exists := c.KeyIsExists(key); exists == false {
		sMap.Store(key, value)
		res = true
	} else {
		if variable.ZapLog == nil {
			log.Fatal(my_errors.ErrorsContainerKeyAlreadyExists + ",请解决键名重复问题,相关键：" + key)
		} else {
			variable.ZapLog.Warn(my_errors.ErrorsContainerKeyAlreadyExists + ", 相关键：" + key)
		}
	}
	return
}

// Delete 2. 将键值对进行删除
func (c *containers) Delete(key string) {
	sMap.Delete(key)
}

// Get 3.传递健，从容器中获取
func (c *containers) Get(key string) interface{} {
	if val, exists := c.KeyIsExists(key); exists {
		return val
	}
	return nil
}

// KeyIsExists 4.判断键值是否存在
func (c *containers) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

// FuzzyDelete 5.按照键的前缀模糊删除容器中注册的内容
func (c *containers) FuzzyDelete(keyPre string) {
	sMap.Range(func(key, value interface{}) bool {
		if keyName, ok := key.(string); ok {
			if strings.HasPrefix(keyName, keyPre) {
				sMap.Delete(keyPre)
			}
		}
		return true
	})
}
