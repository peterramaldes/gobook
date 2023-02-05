## 2.3.2 Ponteiros

```golang
var x int // declaration
&x        // `x` address 
*x = p    // p pointer to x or p has the x address


x := 1
p := &x         // p, type of *int, pointer to x
fmt.Println(*p) // "1"
*p = 2          // same as x = 2
fmt.Println(x)  // "2"
```
