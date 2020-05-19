## Composite Data Types
### Arrays 
- A sequence of elements of a single type. __Length :__ No. of Elements.
- Using Arrays
  ```
    var arr[10] int // Declaring an Array
    arr[0] = 10 // setting value for arr[0] to 10
    fmt.Println(arr[0])
    fmt.Println(arr) 
  ```
### Slices
- Slice is a special data type in Go! It can be created by slicing an array or using a Slice Constructor. 
- Slicing an Array 
  ```
    arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
    s1 := arr[1:3] // Slicing arr 
    fmt.Println(s1)
  ```
- Slice Constructor 
  ```
    sli:= []int{1,2,3,4,5}
	  fmt.Println(sli)
  ```
### HashTables
- HashTables in Go is same as Hash implementation in other programming languages.
- HashTables are internally implemented through Maps in Go.

### Maps
- Maps in Go link Key,Value pairs together.
- Implementing HashTables through Maps
  ```
  id := map[string]int{"joe": 123}
  fmt.Print(id["Alan"])
  ```

