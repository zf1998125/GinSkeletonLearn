/**
*
*
* @author 张帆
* @date 2024/06/04 17:09
**/
package snow_flake

import (
	"GinSkeletonLearn/app/global/consts"
	"GinSkeletonLearn/app/global/variable"
	"GinSkeletonLearn/app/utils/snow_flake/snow_flake_interf"
	"sync"
	"time"
)

func CreateSnowFlakeFactory() snow_flake_interf.InterfaceSnowFlake {
	return &snowFlake{
		timestamp: 0,
		machineId: variable.ConfigYml.GetInt64("SnowFlake.SnowFlakeMachineId"),
		sequence:  0,
	}
}

type snowFlake struct {
	sync.Mutex
	timestamp int64
	machineId int64
	sequence  int64
}

func (s *snowFlake) GetId() int64 {
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & consts.SequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	return (now-consts.StartTimeStamp)<<consts.TimestampShift | (s.machineId << consts.MachineIdShift) | (s.sequence)
}
