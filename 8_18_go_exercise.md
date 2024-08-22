# 8/18 今日題目

1. 條件語句：

 > 請撰寫一個 Go 程式，判斷一個整數是否為偶數。如果是偶數，輸出 "Even"，否則輸出 "Odd"。

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

2. 多重返回值的函式：

 > 在 Go 中，函式可以返回多個值。請撰寫一個函式，接受兩個整數並返回它們的和與差。

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

3. 具名返回值的函式：

 > 請撰寫一個具名返回值的函式，計算一個矩形的面積與周長，並返回這兩個值。

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

4. 套件、匯入與匯出：

 > 請撰寫一個簡單的 Go 程式，定義一個自訂套件，並從主程式中匯入該套件中的函式。

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

5. 型別轉換：

 > 請解釋 Go 中的型別轉換與型別推論，並提供一個範例程式展示這兩個概念。

```go
func main(){
  value := 42 // 型別推論
  var some float64 = float64(value) // 型別轉換
}
```

6. 陣列與切片：

 > 請撰寫一個程式，建立一個有 5 個整數的陣列，然後將該陣列轉換為切片並進行操作。

```go
package main

import "fmt"

func main() {
 ary := []int{1, 2, 3, 4, 5}

 fmt.Println(ary[1:3])
}
```

7. 結構體：

 > 請撰寫一個表示學生的結構體，其中包含姓名、年齡和成績，並展示如何建立和操作此結構體的實例。

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

8. make() 函式：

請解釋 make() 函式的用途，並撰寫一個範例程式來展示 make() 如何用於建立切片。

```go
  numbers := make([]int, 5, 10)
 fmt.Println(numbers)      // 輸出: [0 0 0 0 0]
 fmt.Println(len(numbers)) // 輸出: 5
 fmt.Println(cap(numbers)) // 輸出: 10

 numbers = append(numbers, 1, 2, 3, 4, 5)
 fmt.Println(numbers)      // 輸出: [0 0 0 0 0 1 2 3 4 5]
 fmt.Println(len(numbers)) // 輸出: 10
```
