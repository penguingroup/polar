package utils

//依照snowflake算法得来guid, 做了些许更改，每毫秒仅能产生256个号，能够保证19年不重号

import (
	"errors"
	"math"
	"sync"
	"time"
)

const (
	// 2018-0-0 0:0:0
	poch = 1495875155000
	// 本地ip地址，16位
	WorkerIDBits = uint64(16)
	// 发号序号 8位
	SenquenceBits  = uint64(8)
	WorkerIDShift  = SenquenceBits
	TimeStampShift = SenquenceBits + WorkerIDBits
	SequenceMask   = int64(-1) ^ (int64(-1) << SenquenceBits)
	MaxWorker      = int64(-1) ^ (int64(-1) << WorkerIDBits)
)

//GUID GUID定义
type GUID struct {
	sync.Mutex
	//Sequence 序列号
	Sequence int64
	//lastTimestamp 上一次时间戳
	lastTimeStamp int64
	//lastID 上一次生成的id
	lastID int64
	//WorkID
	WorkID int64
}

var g *GUID

//NewGUID 获取一个GUID对象
func init() {
	g = new(GUID)
	g.Lock()
	defer g.Unlock()
	if workid, err := workBitPrivateIP(); err != nil {
		g.WorkID = 0
	} else {
		g.WorkID = workid
	}
}

//milliseconds 获得当前毫秒时间
func (g *GUID) milliseconds() int64 {
	return time.Now().UnixNano() / 1e6
}

//NextID 获取一个GUID
func NextID() (int64, error) {
	var ts int64
	var err error
	g.Lock()
	defer g.Unlock()
	ts = g.milliseconds()
	if ts == g.lastTimeStamp {
		g.Sequence = (g.Sequence + 1) & SequenceMask
		if g.Sequence == 0 {
			ts = g.timeStamp(ts)
		}
	} else {
		g.Sequence = 0
	}
	if ts < g.lastTimeStamp {
		err = errors.New("时钟过期")
		return 0, err
	}
	g.lastTimeStamp = ts

	ts = (ts-poch)<<TimeStampShift | g.WorkID<<WorkerIDShift | g.Sequence
	return ts, nil
}

//timeStamp 获取一个可用时间基数
func (g *GUID) timeStamp(lastTimeStamp int64) int64 {
	ts := g.milliseconds()
	for {
		if ts < lastTimeStamp {
			ts = g.milliseconds()
		} else {
			break
		}
	}
	return ts
}

func workBitPrivateIP() (int64, error) {
	ip, err := PrivateIPv4()
	if err != nil {
		return 0, err
	}
	return int64(ip[2])<<8 | int64(ip[3]), nil
}

func Round(f float64, n int) float64 {
	pow10 := math.Pow10(n)
	return math.Trunc((f+0.5/pow10)*pow10) / pow10
}
