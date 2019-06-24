型態轉換不回傳噴錯
===

我們心目中的嚴重錯誤才算錯，有錯就噴panic！

## 問題需求
1. 有時候，回傳值若不止一個，其實非常困擾，例如在需要將大量的參數寫入struct裡面的時機。
1. 有些狀態值有錯誤並不要緊，我們不這麼在乎。
1. 若某些重要值發生錯誤，根本不需要繼續往下做，預期整塊都是壞光光，不如幫你爆panic，順便逼迫你在一開始就思考，該值做型態轉轉換時，是不是真的很重要，提早做決定。



## 情境如下
``` go

numInt, conErr := strconv.Atoi(num)
if conErr != nil {
    // 錯誤處理
}

sumCheckInt, conErr2 := strconv.Atoi(sumCheck)
if conErr2 != nil {
    // 錯誤處理
}

setTimesInt, conErr2 := strconv.Atoi(setTimes)
if conErr2 != nil {
    // 錯誤處理
}

userInfo := &UserData{
    docNum : numInt,
    sumCheck : sumCheckInt,
    setTimes : setTimesInt,
}
```

## 設計與想法
1. 只要拿一個回傳值。
1. 用input參數決定重不重要，要不要爆panic。
1. 傳入function 名稱，寫log紀錄也可以一併處理。
1. 不限制string轉int，其他不同型態轉換也可比照。
1. 只要有用這個型態轉換method，一定要recover做準備。
