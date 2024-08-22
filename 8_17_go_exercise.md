# 8/17 今日題目

 1. 基本語法：

 > 請撰寫一個簡單的 Go 程式，輸出 "Hello, Go!" 字串。

```go
package main

import (
 "fmt"
)

func main() {
 fmt.Println("Hello, World!")
}
```

 2. 變數宣告：

> 在 Go 中，有多種宣告變數的方式。請舉出並示範三種不同的變數宣告方式，並說明每種方式的用途。

```go
// 1. 變數宣告(variable declaration)
var name string
name = "John Doe"

// 2. 直接定義(direct define)
var name string = "John Doe"

// 3. 簡短宣告(short declaration)
name := "John Doe"
```

 3. 基本資料型別：

> 請宣告一個 int 型別的變數，並將值 42 賦予該變數。然後，宣告一個 float64 型別的變數，並將值 3.14 賦予該變數。最後，將這兩個變數的值相加並輸出結果。

```go
package main

import ("fmt")

func main() {
  var intNumber = 42
  var float64Number = 3.14

  fmt.Println(float64Number + float64(intNumber))
}

```

 4. 變數與常數：

 > 在 Go 中，常數是不可變的。請宣告一個常數 Pi，其值為 3.14159，並說明常數與變數的差異

 ```go
// 常數宣告時就必須初始化且無法再變更裡面的數值
// 變數宣告時不一定要初始化也可以變更數值
 const Pi = 3.14159

 ```

 5. 資料型別轉換：

 > 請撰寫一個程式，將一個 int 型別的變數轉換為 float64 型別，並輸出轉換後的結果。

 ```go
 package main

import ("fmt")

func main() {
  var intNumber = 42

  fmt.Println(float64(intNumber))
}
 ```

 6. 複數型別：

 > Go 支援複數資料型別。請宣告一個 complex64 型別的變數，並將值 1 + 2i 賦予該變數。然後，輸出這個複數的實部和虛部。

 ```go
 package main

import (
    "fmt",
    "math/cmplx"
)

func main() {
  var a complex = 1 + 2i

  fmt.Println("實部:", real(a))
  fmt.Println("虛部:", imag(a))
}
 ```

 7. 字元型別：

 > 請宣告一個 rune 型別的變數，並將 Unicode 字元 U+0041（字母 A）賦予該變數，然後輸出該變數。

 ```go
 package main

import (
    "fmt",
)

func main() {
  var r rune = 'A'

  fmt.Println(r)
}
 ```

 8. 布林型別：

 > 請撰寫一個程式，宣告一個布林變數 isGoFun，並將其設為 true。根據該變數的值，輸出 "Go is fun!" 或 "Go is not fun."。

 ```go
 package main

import (
    "fmt",
)

func main() {
  var isGoFun = true

  if isGoFun {
    fmt.Println("Go is fun!")
  } else {
    fmt.Println("Go is not fun.")
  }
}
 ```

 9. 指標的基礎：

 > 請解釋在 Go 中如何使用指標，並撰寫一個簡單的程式來練習指標的基本使用。

 ```go
 package main

import (
    "fmt",
)

func updateValue(p *int) {
    *p = 100
}

func main() {
  var x int = 500

  updateValue(&x)
  fmt.Println(x)
}
 ```
