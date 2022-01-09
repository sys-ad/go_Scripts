package main 

import "fmt"
        
func main() {
    
    var t, n, k int // t = number of test cases, n = number of students (size of array input array length), k = threshold
    var count = 0
    var a []uint64
    
    fmt.Scan(&t, &n, &k)
    a = make([]uint64, n)
        
    for i := 0; i <= n; i++ {
        fmt.Scanf("%d", &a[i])
         if a[i] <= 0 {
             count++
         }else {
                break
         }
    }
    if count <= k {
        fmt.Println("YES")
    }else {
        fmt.Println("NO")
    }
}
