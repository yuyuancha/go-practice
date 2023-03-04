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

    展示使用資料結構基礎型態——陣列，並且透過 `for loop` 打印陣列內的所有元素。

- linked-list

    展示使用資料結構基礎型態——鏈結串列，並且透過 `function` 遞迴打印出所有節點。

- stack-array

    透過陣列來實作關於 stack(堆疊)的特性——FILO(先進後出)，分別依序放入三個值，隨後取出一個再放入一個值，後取出全部，可以感受到堆疊先進後出的特性。

    > 陣列來實作堆疊時，需要多 `top` 變數來接收目前最後的位置，來得知目前操作放入或取出的陣列位置。

- stack-linked-list

    透過鏈結串列來實作關於 stack(堆疊)的特性——FILO(先進後出)，分別依序放入三個值，隨後取出一個再放入一個值，後取出全部，可以感受到堆疊先進後出的特性。

    > 鏈結串列實作堆疊相對陣列來得簡單，放入時只需要賦予新的值並將 `next` 指向原本的節點即可，取出時也只需要把節點指向 `next`(視同直接丟棄第一個節點)。

### algorithm (演算法)

- sort (排序)

    - selection-sort

        實作選擇排序，透過 `for loop` 走訪一遍比較出最小值，並將最小值與目前走訪的節點做替換位置，時間複雜度為 `O(n^2)`。
    
    - insertion-sort

        實作插入排序，透過比對前面的數，若比之小，則替換位置。比較完最後一個數，即完成排序，時間複雜度為 `O(n^2)`。
    
    - bubble-sort

        實作氣泡排序，透過和下一個數字相比，如果比較大則替換。要是一輪都沒有替換位置，表示已經排序完成，可以提早跳出迴圈，時間複雜度為 `O(n^2)`(最佳狀況即只做一輪就完成，複雜度則為 `O(n)`)。

    - quick-sort

        實作快速排序，先選出中樞數字，之後透過分組的方式進行排序。將比較大的分到右邊、較小的分到左邊，接著分別進行排序，接著合併即完成排序，時間複雜度平均為 `O(nLog(n))`(最佳狀況為 `O(nLog(n))`，就是剛好第一個中樞數字為中位數，將資料切二等分。最糟狀況為 `O(n^2)`，就是中樞數字每次都拿到最大或最小值)。
    
    - merge-sort

        實作合併排序，先將數字左右邊分類，再分別透過遞迴函式再分類，接著開始往回做比較。比較左右邊的數字，較小的就合併回去，並檢查若有一邊已經合併完，則將另一邊合併回去，即可完成排序，時間複雜度為 `O(nLog(n))`。

### Other

- corn-job

    Corn job 排程的展示，透過每兩秒執行一次及每五秒執行一次，打印出執行當下時間，可以了解排程的運作方式。

## 參考

- [Go 簡單例子來理解 sync.Mutex 和 sync.RWMutex](https://clouding.city/go/mutex-rwmutex/)
- [golang开发一个简单的grpc](https://waterflow.link/articles/1665674508275)
- [各種排序法一起跑的影片:Visualization of 24 Sorting Algorithms In 2 Minutes](https://www.youtube.com/watch?v=BeoCbJPuvSE&ab_channel=ViktorBohush)
