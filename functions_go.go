main package

import (
  "fmt"
)
// examples of anonymous functions.

func main() {
    add := func(m int){
        return m+1
}

    result := add(6)

    println(result)
}

// example of closures. Similar to anonymous functions


addN := func(m int){
  return func(n int){
      return m+n
  }
}
