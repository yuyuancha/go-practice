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

### Other

- corn-job

    Corn job 排程的展示，透過每兩秒執行一次及每五秒執行一次，打印出執行當下時間，可以了解排程的運作方式。

## 參考

- [Go 簡單例子來理解 sync.Mutex 和 sync.RWMutex](https://clouding.city/go/mutex-rwmutex/)
