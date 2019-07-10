package main

import (
	"encoding/json"
	"log"
	"reflect"
)

// BigStruct 假設這是要回傳的主要struct
// 注意這裡的field Name開頭要大寫，對於reflect來說，非公開的屬性會有存取不到的問題
type BigStruct struct {
	Name   string
	ID     int
	Option interface{}
}

func setStructField(b *BigStruct, field string) reflect.Value {
	r := reflect.ValueOf(b)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

func main() {

	b := BigStruct{
		Name: "繃啾好可愛",
		ID:   999,
	}

	// 假設這是lua過來的json string
	result := `{"name":"little johnny","other":{"max_value":9,"option":[{"ratio":1,"times":20},{"ratio":1,"times":15},{"ratio":1,"times":12},{"ratio":1,"times":9},{"ratio":1,"times":6}],"trigger_count":3,"trigger_way":"reel","switch":true}}`

	// 因為lua組裝什麼json，內容是很彈性的，無法也不要事先準備struct type去unmarshal，一切都靠lua組好，便不再更改裡面的內容
	var middleInterface interface{}
	json.Unmarshal([]byte(result), &middleInterface)
	// log.Println(middleInterface)

	f := setStructField(&b, "Option")
	f.Set(reflect.ValueOf(middleInterface))
	// 看一下，內容Option已經被改變了
	log.Println(b)

	// 回傳給前端是json格式
	c, _ := json.Marshal(b)
	log.Println("前端收到的json string")
	log.Println(string(c))
	// result:

	// {
	//   "Name": "繃啾好可愛",
	//   "ID": 999,
	//   "Option": {
	//     "name": "little johnny",
	//     "other": {
	//       "max_value": 9,
	//       "option": [
	//         {
	//           "ratio": 1,
	//           "times": 20
	//         },
	//         {
	//           "ratio": 1,
	//           "times": 15
	//         },
	//         {
	//           "ratio": 1,
	//           "times": 12
	//         },
	//         {
	//           "ratio": 1,
	//           "times": 9
	//         },
	//         {
	//           "ratio": 1,
	//           "times": 6
	//         }
	//       ],
	//       "switch": true,
	//       "trigger_count": 3,
	//       "trigger_way": "reel"
	//     }
	//   }
	// }
}
