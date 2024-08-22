# 8. make() 函式

```go
 numbers := make([]int, 5, 10)
 fmt.Println(numbers)      // 輸出: [0 0 0 0 0]
 fmt.Println(len(numbers)) // 輸出: 5
 fmt.Println(cap(numbers)) // 輸出: 10

 numbers = append(numbers, 1, 2, 3, 4, 5)
 fmt.Println(numbers)      // 輸出: [0 0 0 0 0 1 2 3 4 5]
 fmt.Println(len(numbers)) // 輸出: 10

 numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
 fmt.Println(numbers)      // 輸出: [1 2 3 4 5 6 7 8 9 0]
 fmt.Println(len(numbers)) // 輸出: 10
 fmt.Println(cap(numbers)) // 輸出: 10

 numbers = append(numbers, 1, 2, 3, 4, 5)
 fmt.Println(numbers)      // 輸出: [1 2 3 4 5 6 7 8 9 0 1 2 3 4 5]
 fmt.Println(len(numbers)) // 輸出: 15
 fmt.Println(cap(numbers)) // 輸出: 20

 // 宣告一個有 10 個元素的切片
 base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

 // 創建一個從 base 切片的索引 2 到索引 4 (不含) 的子切片
 subSlice := base[2:5] // 這會包含元素 3, 4, 5

 fmt.Printf("subSlice: Length = %d, Capacity = %d\n", len(subSlice), cap(subSlice))
 fmt.Println("subSlice elements:", subSlice)

 // 宣告一個有 10 個元素的切片
 base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

 // 創建一個長度和容量相同的子切片
 subSlice := base[2:5:5] // 指定容量為 5，這樣長度和容量都為 3

 fmt.Printf("subSlice: Length = %d, Capacity = %d\n", len(subSlice), cap(subSlice))
 fmt.Println("subSlice elements:", subSlice)
```
