package main

import (
	"api-test/api"
	"log"
	"os"
)

func main() {

	// 需要傳入設定檔名稱
	if len(os.Args) > 1 {

		api.TriggerAPI(os.Args[1])
	} else {
		log.Println("須要傳入參數:設定檔名稱 (不需副檔名)")
	}

}
