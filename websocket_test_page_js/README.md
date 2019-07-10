測試golang裡面的websocket，作為client 的 web page (只有html＆js)
===

## 使用重點
1. 丟進websocket message的是一串json type string。
1. server 以解析message裡面commandObj.command作為client所下的命令。
1. 所以下拉式選單，塞的是我們server自定義的command。
1. 因為server重啟，這邊web還要重新整理頁面太麻煩了，所以弄一個重連的選項
1. 我不會寫vue或react，只會寫jqeury，整理過勉強還能使用，維護和修改也方便。

