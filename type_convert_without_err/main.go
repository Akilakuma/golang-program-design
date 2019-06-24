package main

import (
	"fmt"
	"strconv"
)

func main() {
	paraList := []string{"abc", "1", "99"}

	for _, v := range paraList {
		err := dosomething(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("everything is fine")
		}
	}

	// result:

	// dosomethingstring轉成Int 重要參數發生錯誤,傳入: abc
	// everything is fine
	// everything is fine
}

// dosomething 業務邏輯
func dosomething(para string) (methodErr interface{}) {

	// 有嚴重錯誤的panic，會在這裡接到，並且將錯誤往上拋
	defer func() {
		if err := recover(); err != nil {
			methodErr = err
		}
	}()

	convertStrToInt(para, true, "dosomething")
	return nil
}

// convertStrToInt string轉成Int
// needPanic:true 緊要資料有問題觸發panic，needPanic:false 不要緊資料，有錯誤回傳預設值無妨
// funcName 印出log之類可以考慮使用
func convertStrToInt(s string, needPanic bool, funcName string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		if needPanic {
			panic(funcName + "string轉成Int 重要參數發生錯誤,傳入: " + s)
		}
	}
	return r
}

func convert2(s interface{}, needPanic bool, funcName string) interface{} {

	// convert ...
	return nil
}

func convert3(s interface{}, needPanic bool, funcName string) interface{} {

	// convert ...
	return nil
}
