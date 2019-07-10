動態指定struct key，設定value
===

## 說明
搭配lua，根本作為黑科技使用

## 想法
1. 對於全新的專案開發，前後格式搭配好，基本上不需要動態的指定key，再去assign value。
1. 會如此使用是因為專案需求，對接的格式無法變更，並且有硬性的欄位對應。
1. 第一步，lua替代golang一個又一個專屬的struct，自行組裝好json string傳遞給golang。
1. 第二步，lua在傳遞json string的時候，也順便將使用對象的key告訴給golang。
1. 第三步，golang拿到key string的時候，利用reflect的方式將json string塞到原本的struct。

