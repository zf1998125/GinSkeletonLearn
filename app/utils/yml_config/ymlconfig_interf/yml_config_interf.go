/**
*
*
* @author 张帆
* @date 2024/06/04 15:55
**/
package ymlconfig_interf

import "time"

type YmlConfigInterf interface {
	ConfigFileChangeListen()
	Clone(fileName string) YmlConfigInterf
	Get(keyName string) interface{}
	GetString(keyName string) string
	GetInt(keyName string) int
	GetBool(keyName string) bool
	GetInt32(keyName string) int32
	GetInt64(keyName string) int64
	GetFloat64(keyName string) float64
	GetDuration(keyName string) time.Duration
	GetStringSlice(keyName string) []string
}
