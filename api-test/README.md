自製的API測試程式
===

## 說明

要增加的API程式放到api的資料夾

對應的API設定檔放在config

以下三者保持一致, 如"get_session"

1. config 檔案名稱 , --> config/get_session.json
2. 下給main的參數名稱 , --> ./tt get_session
3. api/utility 所加的method 名稱 ,  -->
``` go
var apiMap = map[string]func(*greq.Client, helper.APISetting, *sync.WaitGroup){
	"get_session": getSessionAPI,
}
```

## 執行結果

統計的時間單位除了totalTimeSecond是以秒為單位

其他皆是millisecond

``` bash
2019/08/26 11:13:12 Successfully Opened get_session.json
2019/08/26 11:13:12 {APIName:get_session WorkerNum:50 TotalRequestNum:1000 Strategy:normal Timeout:3 OverTimes:500 IsNeedPara:true Para:[map[game_id:105 ip:127.0.0.1 session:45aae27f54ea4d78c3ddf055ae01c3c78efac4fe59e7b153cc3724fba6c2534f]]}
2019/08/26 11:13:12 ======統計紀錄======
2019/08/26 11:13:15 {apiName:get_session strategy:normal workerNum:50 requestNum:1000 successNum:1000 failNum:0 overTimes:500 overTimeNum:0 costMiniTime:0 costMostTime:286 averageTime:96 totalTimeSecond:96}
```

## 執行方式
go build -o tt
./tt 參數名稱
