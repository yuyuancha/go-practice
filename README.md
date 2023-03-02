# go-practice

放置 GoLang 程式語言相關的練習。
可以進到各類別的資料夾內，執行各自的 `main.go` 檔的內容。

## 目錄

### base (基礎)

- hello-world

    程式語言第一堂課都會教的在 Terminal 打印出 Hello world!。

### Command line (命令列相關操作)

- flag

    GoLang command line 套件 flag 練習。
    可以透過加上 `-var=value` 的方式，對指定的變數賦值。

    ```
    $ go run main.go -name={your name} -age={your age}
    ```

- progress-bar

    透過 `github.com/schollz/progressbar/v3` 套件，在命令列打印出像是下載東西時的進度條，有多種樣式可以選擇，也可以選擇顯示的文字等。

### Goroutine

- mutex

    透過模擬銀行的存錢及查詢餘額行為，來模擬 GoLang Mutex 互斥鎖的功能。
    使同一時間只會有一個 `goroutine` 對該變數做操作，避免資料發生錯誤。

    * 可以透過 `-race` 參數來檢查是否有 `race condition`。

    ```
    $ go run -race main.go
    ```

- rw-mutex

    和 `mutex` 一樣是透過模擬銀行來示範，差別在於使用 `RWMutex`，針對讀取動作允許多個 `goroutine` 使用該參數。

### gRPC

- grpc-example-01

    透過 GoLang 架設一個 `gRPC` 伺服器的範例。

    使用 `protoc --go_out=. --go-grpc_out=. hello.proto` 來快速建立 GoLang 的 `xxx.proto.go` 檔。
    並且實作 Service 對應方法來使其符合 `interface`。

    如果遇到無法使用 protoc 指令生成 `xxx.proto.go`，需要檢查 `GOPATH` 有沒有放入 `PATH` 內。
    
    ```
    $ vim ~/.bash_profile
    ```

    新增:
    ```
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

    新增完畢需要重新啟用:
    ```
    $ source ~/.bash_profile
    ```

    可以透過打印(echo)方式檢查有沒有成功:
    ```
    $ echo $PATH
    ```

### Data structure (資料結構)

- array

    展示使用資料結構基礎型態——陣列。
    並且透過 `for loop` 打印陣列內的所有元素。

- linked-list

    展示使用資料結構基礎型態——鏈結串列。
    並且透過 `function` 遞迴打印出所有節點。

### Other

- corn-job

    Corn job 排程的展示，透過每兩秒執行一次及每五秒執行一次，打印出執行當下時間，可以了解排程的運作方式。

## 參考

- [Go 簡單例子來理解 sync.Mutex 和 sync.RWMutex](https://clouding.city/go/mutex-rwmutex/)
- [golang开发一个简单的grpc](https://waterflow.link/articles/1665674508275)
