package api

import (
	"api-test/helper"
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	greq "github.com/syhlion/greq"
)

// getSessionAPI 測試取得GetSession API
func getSessionAPI(c *greq.Client, setting helper.APISetting, wg *sync.WaitGroup) {

	paraRaw := setting.Para
	IsNeedPara := setting.IsNeedPara
	url := setting.APIURL

	if IsNeedPara && (len(paraRaw) == 0) {
		log.Println("設定檔API參數設定有問題")
		return
	}

	// 隨機取一個參數case
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(paraRaw))

	// io buffer
	para := bytes.NewBuffer(helper.HttpJSONRawBuild(paraRaw[x]))

	// 發API
	tS := time.Now()
	_, s, err := c.PostRaw(url, para)
	tE := time.Now().Sub(tS)

	// 結果紀錄
	var resultStatistic SingleAPIResult

	// 成功或失敗紀錄
	if err != nil {
		log.Println(err)
		resultStatistic.errMessage = err.Error()
		if s != http.StatusOK {
			resultStatistic.isSuccess = false
		}
	} else {
		resultStatistic.isSuccess = true
		// 前後花費時間(成功才算進去)
		resultStatistic.costTime = helper.NanoToMilli(tE.Nanoseconds())
	}

	// API結果送到統計去
	joinStatistic(resultStatistic)

	wg.Done()
}
