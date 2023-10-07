package idgen

import (
	"hash/crc32"
	"os"
	"sync"
	"time"

	"github.com/chenyan/wheels/network"
)

const (
	workerIDBits     = uint(5)                          // 机器id所占位数
	workerIDMask     = int64(-1 ^ (-1 << workerIDBits)) // 机器id掩码
	dataCenterIDBits = uint(5)                          // 数据中心id所占位数
	sequenceBits     = uint(12)                         // 序列所占位数

	maxWorkerID     = int64(-1 ^ (-1 << workerIDBits))     // 支持的最大机器id数量
	maxDataCenterID = int64(-1 ^ (-1 << dataCenterIDBits)) // 支持的最大数据中心id数量
	maxSequence     = int64(-1 ^ (-1 << sequenceBits))     // 支持的最大序列id数量

	timeShift         = uint(22)                    // 时间戳向左移动22位
	dataCenterIDShift = sequenceBits + workerIDBits // 数据中心id向左移动17位(12+5)
	workerIDShift     = sequenceBits                // 机器id向左移动12位
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerID     int64
	dataCenterID int64
	sequence     int64
}

func NewSnowflake(workerID, dataCenterID int64) *Snowflake {
	return &Snowflake{
		timestamp:    0,
		workerID:     workerID,
		dataCenterID: dataCenterID,
		sequence:     0,
	}
}

func (sf *Snowflake) GetID() int64 {
	sf.Lock()
	defer sf.Unlock()

	now := time.Now().UnixNano() / 1e6 // 转换为毫秒
	if sf.timestamp == now {
		sf.sequence = (sf.sequence + 1) % maxSequence
		if sf.sequence == 0 {
			for now <= sf.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.timestamp = now

	return ((now) << timeShift) | (sf.dataCenterID << dataCenterIDShift) | (sf.workerID << workerIDShift) | (sf.sequence)
}

func GenWorkerID() (int64, error) {
	localIP, err := network.GetOutboundIP()
	if err != nil {
		return 0, err
	}
	pid := os.Getpid()
	return int64((crc32.ChecksumIEEE([]byte(localIP)) + uint32(pid)) & uint32(workerIDMask)), nil
}
