package main 

import "fmt"
//import "os"
//import "bufio"

        
func main() {
    
    // in := bufio.NewReader(os.Stdin)    
    var t, n, k int // t = number of test cases, n = number of students (size of array input array length), k = threshold
    var count = 0
    var a []uint64
    a = make([]uint64, n)
    
    fmt.Scan(&t, &n, &k)
    
    for i := 0; i <= n; i++ {
        fmt.Scan(&a[i])
        if a[i] <= 0 {
            count++
        }else {
            break
        }
    }
    if count <= k {
        println("YES")
    }else {
        println("NO")
    }
}
