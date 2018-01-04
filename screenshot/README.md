## 自動でスクリーンショット

### 方法
1. ``` go run main.go ``` を起動する
2. ``` ngrok http 3564 ```を起動して出て来たurlを ```screen/main.go ``` のtarget_urlに貼り付ける
3. ``` go build -o screen screen/main.go ```
4. できたバイナリを実行するだけ

### 結果

``` tmp/```にスクリーンショットが配置される

