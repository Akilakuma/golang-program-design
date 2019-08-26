package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// APISetting 測試API的相關設定
type APISetting struct {
	APIName         string                   `json:"api_name"`              // API 名稱
	APIURL          string                   `json:"api_url"`               // API URL
	WorkerNum       int                      `json:"worker_num"`            // 同時打gorutine的數量
	TotalRequestNum int                      `json:"total_request_num"`     // 總request數量
	Strategy        string                   `json:"strategy"`              // 策略：嚴格或一般
	Timeout         int                      `json:"timeout"`               // 強制截斷的秒數
	OverTimes       int64                    `json:"over_time_millisecond"` // 超過多少時間記錄下來
	IsNeedPara      bool                     `json:"is_need_para"`          // API 是否須要傳入參數
	Para            []map[string]interface{} `json:"para"`                  // 丟到API的參數
}

// LoadAPIConfig 讀取config檔案設定
func LoadAPIConfig(configName string) APISetting {

	var apiSetting APISetting

	jsonFile, err := os.Open("config/" + configName + ".json")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Opened " + configName + ".json")

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &apiSetting)

		// log.Println("%+v", apiSetting)
	}

	return apiSetting
}
