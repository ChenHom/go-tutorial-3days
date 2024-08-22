# 8/19 今日題目

1. 循環語句：

> 撰寫一個 Go 程式，使用 for 循環打印出從 1 到 10 的所有數字。

```go
count := 10
 for i := 0; i < count; i++ {
  fmt.Printf("i: %d\n", i)
 }
```

2. 多重條件的 for 循環：

> 撰寫一個程式，使用 for 循環與多個條件，打印出 1 到 20 之間的所有偶數。

```go
count := 20
 for i := 1; i <= count; i++ {
  if i%2 == 0 {
   fmt.Printf("Even i: %d\n", i)
  }
 }
```

3. for-range 語句：

> 請撰寫一個程式，使用 for-range 迴圈遍歷一個包含數字的切片，並打印出每個元素的值與其索引。

```go
for index, value := range []int{1, 2, 3, 4, 5} {
  fmt.Println(index, value)
 }
```

4. 介面（interface）：

> 撰寫一個程式，定義一個介面 Shape，該介面包含兩個方法：Area() 和 Perimeter()。然後，為一個表示矩形的結構體 Rectangle 實現這個介面，並展示如何使用介面來調用這些方法。

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

type Rectangle struct {
 Width  float64
 Height float64
}

func (r Rectangle) Area() float64 {
 return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
 return 2 * (r.Width + r.Height)
}

type Circle struct {
 Radius float64
}

func (c Circle) Area() float64 {
 return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
 return 2 * 3.14 * c.Radius
}

func main() {
 shapes := []Shape{
  Rectangle{10, 20},
  Circle{10},
 }

 for _, shape := range shapes {
  fmt.Printf("Area is %.2f\n", shape.Area())
  fmt.Printf("Perimeter is %.2f\n\n", shape.Perimeter())
 }
}
```

5. 介面與多態性：

> 請撰寫一個程式，創建一個函式 PrintShapeDetails，該函式接受一個 Shape 介面的參數，並打印出該形狀的面積和周長。示範如何使用多態性處理不同的形狀（例如矩形和圓形）。

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

type Rectangle struct {
 Width  float64
 Height float64
}

func (r Rectangle) Area() float64 {
 return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
 return 2 * (r.Width + r.Height)
}

type Circle struct {
 Radius float64
}

func (c Circle) Area() float64 {
 return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
 return 2 * 3.14 * c.Radius
}

func main() {
 shapes := []Shape{
  Rectangle{10, 20},
  Circle{10},
 }

 for _, shape := range shapes {
  fmt.Printf("Area is %.2f\n", shape.Area())
  fmt.Printf("Perimeter is %.2f\n\n", shape.Perimeter())
 }
}
```

6. 錯誤處理：

> 撰寫一個簡單的 Go 程式，示範如何使用 error 介面來進行錯誤處理。設計一個簡單的函式，當傳入的參數為負數時返回錯誤，否則返回計算結果。

```go
```

7. 自定義錯誤：

> 撰寫一個程式，定義一個自定義錯誤，當發生特定情況（例如除以零）時返回該錯誤，並在主函式中處理這個錯誤。

```go
package main

import (
 "errors"
 "fmt"
 "math"
)

// SquareRoot 計算一個數字的平方根，如果輸入為負數則返回錯誤
func SquareRoot(x float64) (float64, error) {
 if x < 0 {
  return 0, errors.New("cannot calculate the square root of a negative number")
 }
 return math.Sqrt(x), nil
}

func main() {
 // 測試傳入正數
 result, err := SquareRoot(9)
 if err != nil {
  fmt.Println("Error:", err)
 } else {
  fmt.Printf("The square root of 9 is: %.2f\n", result)
 }

 // 測試傳入負數
 result, err = SquareRoot(-9)
 if err != nil {
  fmt.Println("Error:", err)
 } else {
  fmt.Printf("The square root of -9 is: %.2f\n", result)
 }
}
```

8. Goroutines：

> 請撰寫一個簡單的 Go 程式，使用 Goroutines 來並行執行兩個函式，並觀察執行結果的順序。

```go
package main

import (
 "fmt"
 "time"
)

func firstFunction() {
 for i := 1; i <= 5; i++ {
  fmt.Printf("First function: %d\n", i)
  time.Sleep(100 * time.Millisecond) // 模擬耗時操作
 }
}

func secondFunction() {
 for i := 1; i <= 5; i++ {
  fmt.Printf("Second function: %d\n", i)
  time.Sleep(150 * time.Millisecond) // 模擬耗時操作
 }
}

func main() {
 // 並行執行 firstFunction 和 secondFunction
 go firstFunction()
 go secondFunction()

 // 等待兩個 Goroutine 執行完成
 time.Sleep(1 * time.Second)
 fmt.Println("Main function finished")
}

```

9. 通道（Channels）：

> 撰寫一個 Go 程式，使用通道來在兩個 Goroutines 之間傳遞數據。展示如何使用通道進行同步。

```go
package main

import (
 "fmt"
 "time"
)

// producer 將數據發送到通道
func producer(ch chan int) {
 for i := 1; i <= 5; i++ {
  fmt.Printf("Producing: %d\n", i)
  ch <- i // 將數據發送到通道
  time.Sleep(100 * time.Millisecond) // 模擬生產耗時
 }
 close(ch) // 生產完成後關閉通道
}

// consumer 從通道接收數據
func consumer(ch chan int) {
 for num := range ch { // 從通道中接收數據
  fmt.Printf("Consuming: %d\n", num)
  time.Sleep(150 * time.Millisecond) // 模擬消耗耗時
 }
}

func main() {
 // 創建一個無緩衝的通道
 ch := make(chan int)

 // 並行執行 producer 和 consumer
 go producer(ch)
 go consumer(ch)

 // 等待 Goroutines 完成
 time.Sleep(1 * time.Second)
 fmt.Println("Main function finished")
}

```
