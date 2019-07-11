package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {
	log.Println("simple case result is :")
	simpleCase()
	fmt.Println()
	log.Println("complex case result is :")
	complexCase()
}

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

// simpleCase 基本的case，要塞的目標就在第一層而已
func simpleCase() {

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

// complexCase 如果要塞的目標在有幾個深度的struct，set field就需要有一些比較特別的處理
func complexCase() {

	result := `{"name":"little johnny","other":{"max_fg":9,"opt":[{"double":1,"times":20},{"double":1,"times":15},{"double":1,"times":12},{"double":1,"times":9},{"double":1,"times":6}],"select_trigger_count":3,"select_trigger_way":"reel","switch":true}}`

	var middleInterface interface{}
	json.Unmarshal([]byte(result), &middleInterface)
	// log.Println(middleInterface)

	var c ComplexStruct

	c.Name = "meow meow"
	c.ID = 12355679

	f := setSTInternalField(&c, "Option")
	f.Set(reflect.ValueOf(middleInterface))
	log.Println(c)

	// result:
	// {meow meow 12355679 { {map[name:little johnny other:map[max_fg:9 opt:[map[double:1 times:20] map[double:1 times:15] map[double:1 times:12] map[double:1 times:9] map[double:1 times:6]] select_trigger_count:3 select_trigger_way:reel switch:true]] 0}}}
}

type ComplexStruct struct {
	Name string
	ID   int
	Res  struct {
		ErrMsg string
		Data   struct {
			Option interface{}
			Flag   int
		}
	}
}

func setSTInternalField(b *ComplexStruct, field string) reflect.Value {
	// 特別注意這裡，如果不是用&，則會出現錯誤：panic: reflect: reflect.Value.Set using unaddressable value
	// 這也蠻好理解，如果不是用address的話，如何去改到目標strcut的某個field資料內容
	r := reflect.ValueOf(&b.Res.Data)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}
