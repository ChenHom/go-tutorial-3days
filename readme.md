# Go 練習

- **8_17_go_exercise.md**: 8月17日的練習說明文件。
- **8_18_feedback.md**: 8月18日的反饋文件。
- **8_18_go_exercise_addition.md**: 8月18日的額外練習說明文件。
- **8_18_go_exercise.md**: 8月18日的練習說明文件。
- **8_19_go_exercise.md**: 8月19日的練習說明文件。
- **go.mod**: Go 模組文件，包含模組名稱和依賴項。
- **main.go**: 主程式文件，包含應用程式的入口點。

## 安裝與使用

### 前置條件

- 安裝 [Go](https://golang.org/dl/)

### 安裝步驟

1. 克隆專案到本地端：

    ```sh
    git clone <專案的 Git URL>
    cd <專案目錄>
    ```

2. 安裝依賴：

    ```sh
    go mod tidy
    ```

### 執行專案

在專案目錄下執行以下命令來啟動應用程式：

```sh
go run main.go
