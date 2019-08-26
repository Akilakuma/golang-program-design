package api

import (
	"api-test/helper"
	"log"
	"sync"
	"time"

	greq "github.com/syhlion/greq"
	requestwork "github.com/syhlion/requestwork.v2"
)

var apiMap = map[string]func(*greq.Client, helper.APISetting, *sync.WaitGroup){
	"get_session": getSessionAPI,
}

var wg sync.WaitGroup

// TriggerAPI API連線測試開打
func TriggerAPI(configName string) {

	// 讀取設定檔
	c := helper.LoadAPIConfig(configName)
	// 將設定檔資訊寫入統計結果
	SetRoundStatistic(c)
	log.Printf("%+v", c)
	log.Println("======統計紀錄======")

	// 建立workNum個worker
	worker := requestwork.New(c.WorkerNum)

	// client 連線client
	APIConn := greq.New(worker, time.Duration(c.Timeout)*time.Second, false)
	fn, ok := apiMap[c.APIName]
	if !ok {
		log.Println("api name:" + c.APIName + " 不存在已知設定")
		return
	}

	tS := time.Now()
	// 執行method
	for i := 0; i < c.TotalRequestNum; i++ {
		wg.Add(1)
		go fn(APIConn, c, &wg)

		if i%c.WorkerNum == 0 {
			if c.Strategy == "normal" {
				wg.Wait()
			}
		}
	}

	wg.Wait()
	tE := time.Now().Sub(tS)

	// 發完所有request後，做最後統計
	statistic := CalculateStatistic()
	statistic.totalTimeSecond = tE.Seconds()
	log.Printf("%+v", statistic)
}
