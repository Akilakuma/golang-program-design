package helper

import (
	"strconv"
	"strings"
	"time"
)

// HttpJSONRawBuild 將map組成json格式
func HttpJSONRawBuild(a map[string]interface{}) []byte {
	var retString = "{"
	for k, v := range a {

		switch t := v.(type) {
		case string:
			sub := (`"` + k + `"` + `:"` + v.(string) + `"`)
			retString += (sub + ",")
		case int:
			v = strconv.Itoa(t)
			sub := (`"` + k + `"` + `:` + v.(string))

			retString += (sub + ",")
		case []string:
			// string array裡面的東西有沒有雙引號會不會有問題，待驗證
			sub := strings.Join(v.([]string), ",")
			sub = strings.TrimSuffix(sub, ",")
			retString += "[" + sub + "]"
		}
	}
	// 最後的,移除掉
	retString = strings.TrimSuffix(retString, ",")
	retString += "}"

	// log.Println(retString)

	return []byte(retString)
}

// NanoToMilli nanoseconds 轉換成 millisecond
func NanoToMilli(nano int64) int64 {
	return nano / int64(time.Millisecond)
}
