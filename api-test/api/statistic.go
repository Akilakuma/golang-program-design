package api

import (
	"api-test/helper"
	"sync"
)

// SingleAPIResult 單次結果紀錄
type SingleAPIResult struct {
	isSuccess  bool
	costTime   int64
	errMessage string
}

// Recorder 紀錄器
type Recorder struct {
	costTimes  []int64
	successNum int
	failNum    int
	miniTime   int64
	maxTime    int64
}

// RoundStatistic 命令執行總統計
type RoundStatistic struct {
	apiName         string  // API 名稱
	strategy        string  // 策略
	workerNum       int     // gorutine 數量
	requestNum      int     // 總處理需求數量
	successNum      int     // 成功數量
	failNum         int     // 失敗數量
	overTimes       int64   // 設定的超過時間
	overTimeNum     int     // 超過設定預期時間數量
	costMiniTime    int64   // 最小花費時間
	costMostTime    int64   // 最大花費時間
	averageTime     int64   // 平均時間
	totalTimeSecond float64 // 所有request做完花費時間
}

var (
	roundStatistic RoundStatistic
	recorder       Recorder
	l              sync.RWMutex
)

// SetRoundStatistic 設定statistic 細節
func SetRoundStatistic(s helper.APISetting) {
	roundStatistic.apiName = s.APIName
	roundStatistic.strategy = s.Strategy
	roundStatistic.workerNum = s.WorkerNum
	roundStatistic.requestNum = s.TotalRequestNum
	roundStatistic.overTimes = s.OverTimes
}

func joinStatistic(r SingleAPIResult) {
	l.Lock()
	recorder.costTimes = append(recorder.costTimes, r.costTime)
	if r.isSuccess {
		recorder.successNum = recorder.successNum + 1
	} else {
		recorder.failNum = recorder.failNum + 1
	}
	if recorder.miniTime == 0 {
		recorder.miniTime = r.costTime
	} else {
		if r.costTime < recorder.miniTime {
			recorder.miniTime = r.costTime
		}
	}
	if r.costTime > recorder.maxTime {
		recorder.maxTime = r.costTime
	}

	// log.Printf("%+v", len(recorder.costTime))
	l.Unlock()
}

// CalculateStatistic 統計整個結
func CalculateStatistic() RoundStatistic {
	roundStatistic.costMiniTime = recorder.miniTime
	roundStatistic.costMostTime = recorder.maxTime
	roundStatistic.failNum = recorder.failNum
	roundStatistic.successNum = recorder.successNum

	var (
		totalTime   int64
		overTimeNum int
	)
	for _, v := range recorder.costTimes {
		totalTime = totalTime + v

		if v > roundStatistic.overTimes {

			overTimeNum = overTimeNum + 1
		}
	}

	roundStatistic.overTimeNum = overTimeNum
	if len(recorder.costTimes) > 0 {
		roundStatistic.averageTime = totalTime / int64(len(recorder.costTimes))
	}
	return roundStatistic
}
