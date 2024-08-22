# 8/18 今日題目

## 1. 條件語句

**題目**: 撰寫一個 Go 程式，判斷一個整數是否為偶數。如果是偶數，輸出 "Even"，否則輸出 "Odd"。

**你的答案**:

```go
package main

import (
    "fmt",
)

func main() {
  var intValue int = 2

 if intValue%2 == 0 {
  fmt.Println("Even")
 } else {
  fmt.Println("Odd")
 }
}
```

**回饋**: 程式碼邏輯正確，但 `import` 部分使用了逗號 `,`，應該用換行來區分每個引入的包。此外，變數命名應該遵循 Go 的命名慣例，使用駝峰命名法或簡單的全小寫單詞。

## 2. 多重返回值的函式

**題目**: 撰寫一個函式，接受兩個整數並返回它們的和與差。

**你的答案**:

```go
package main

func calc(a int, b int) (int, int) {
 return a + b, a - b
}

func main() {
 // 透過 := 宣告變數 a, b 並指派值
 a, b := 6, 2

 // 呼叫 calc 函式，並取得回傳值
 sum, diff := calc(a, b)

 // 顯示回傳值
 println("sum:", sum)
 println("diff:", diff)
}
```

**回饋**: 程式碼正確地使用多重返回值的函式使用。注意，`println` 是非標準庫函數，建議使用 `fmt.Println` 來輸出結果。

## 3. 具名返回值的函式

**題目**: 撰寫具名返回值的函式，計算矩形的面積與周長。

**你的答案**:

```go
package main

func calcRectangle(length int, width int) (area int , perimeter int) {
 area := length * width
 perimeter := (length * 2) + (width * 2)
 return
}

func main() {
 // 呼叫 calcRectangle 函式，並取得回傳值
 area, perimeter := calc(2, 4)

 // 顯示回傳值
 println("area:", area)
 println("perimeter:", perimeter)
}
```

**回饋**: 程式碼有一個小錯誤。你在函式中使用 `:=` 重新定義了 `area` 和 `perimeter` 變數，這會導致返回值為零。應該使用 `=` 來賦值。此外，`calc` 應該改為 `calcRectangle`。

### 修正後的正確答案

```go
package main

import "fmt"

// calcRectangle 計算矩形的面積與周長，並返回這兩個值
func calcRectangle(length int, width int) (area int, perimeter int) {
    area = length * width        // 使用 = 賦值給具名返回值
    perimeter = (length * 2) + (width * 2) 
    return                       // 直接返回具名返回值
}

func main() {
    // 呼叫 calcRectangle 函式，並取得回傳值
    area, perimeter := calcRectangle(2, 4)

    // 顯示回傳值
    fmt.Println("area:", area)
    fmt.Println("perimeter:", perimeter)
}
```

## 4. 套件、匯入與匯出

**題目**: 定義一個自訂套件，並從主程式中匯入該套件中的函式。

**你的答案**:

```go
package myPackage

func someFunc(args int) int {
  return args * 2
}

package main

import(myPackage)

func main(){
  println(someFunc(2))
}
```

**回饋**: Go 語言中的一個檔案只能有一個 `package` 宣告，因此這樣的程式碼無法執行。此外，`import` 應該使用引號來指定路徑，如 `import "myPackage"`。

## 5. 型別轉換

**題目**: 解釋 Go 中的型別轉換與型別推論，並提供範例程式展示這兩個概念。

**你的答案**:

```go
func main(){
  value := 42 // 型別推論
  var some float64 = float64(value) // 型別轉換
}
```

**回饋**: 範例正確地展示了型別推論與型別轉換的基本用法。解釋也很簡潔明瞭。

## 6. 陣列與切片

**題目**: 建立一個有 5 個整數的陣列，然後將該陣列轉換為切片並進行操作。

**你的答案**:

```go
package main

import "fmt"

func main() {
 ary := []int{1, 2, 3, 4, 5}

 fmt.Println(ary[1:3])
}
```

**回饋**: 你的範例中定義的 `ary` 其實是一個切片而非陣列。如果你要展示從陣列到切片的轉換，應該先定義一個陣列，再進行切片操作。

### 修正後的正確答案

```go
package main

import "fmt"

func main() {
    // 建立一個包含 5 個整數的陣列
    var ary [5]int = [5]int{1, 2, 3, 4, 5}

    // 將陣列轉換為切片
    slice := ary[:]

    // 對切片進行操作，例如取出一部分並列印
    fmt.Println("切片的部分內容:", slice[1:3])
    fmt.Println("切片的完整內容:", slice)
}
```

- 陣列：固定長度，不能使用 append 增加元素。
- 切片：動態長度，可以使用 append 增加元素。

## 7. 結構體

**題目**: 撰寫一個表示學生的結構體，展示如何建立和操作此結構體的實例。

**你的答案**:

```go
package main

import "fmt"

// Student 結構體表示一個學生
type Student struct {
 Name   string // 姓名
 Age    int    // 年齡
 Grades []int  // 成績
}

func calculateAverage(grades []int) float32 {
 sum := 0
 for _, grade := range grades {
  sum += grade
 }
 return float32(sum) / float32(len(grades))
}

func main() {
 stu := Student{
  Name:   "John",
  Age:    25,
  Grades: []int{90, 80, 85},
 }

 fmt.Printf("Name: %s\n", stu.Name)
 fmt.Printf("Age: %d\n", stu.Age)
 fmt.Printf("Grades: %v\n", stu.Grades)

 fmt.Printf("Average: %.2f\n", calculateAverage(stu.Grades))
}
```

**回饋**: 這段程式碼展示了如何使用結構體來表示學生的資訊，並計算平均成績。程式碼清晰且功能完整，表現非常好。

## 8. make() 函式

**題目**: 解釋 `make()` 函式的用途，並撰寫範例程式來展示 `make()` 如何用於建立切片。

**你的答案**:

```go
numbers := make([]int, 5, 10)
fmt.Println(numbers)      // 輸出: [0 0 0 0 0]
fmt.Println(len(numbers)) // 輸出: 5
fmt.Println(cap(numbers)) // 輸出: 10

numbers = append(numbers, 1, 2, 3, 4, 5)
fmt.Println(numbers)      // 輸出: [0 0 0 0 0 1 2 3 4 5]
fmt.Println(len(numbers)) // 輸出: 10
```

**回饋**: 這段程式碼正確展示了如何使用 `make()` 來建立切片。程式碼清晰且符合題目要求。
