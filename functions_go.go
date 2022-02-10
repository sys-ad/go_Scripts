main package

import (
  "fmt"
)
// examples of anonymous functions. Note a function withou t a name

func main() {
    add := func(m int){
        return m+1
}

    result := add(6)

    println(result)
}

// example of closures. Similar to anonymous functions. More powerful than anonymous functions.


addN := func(m int){
  return func(n int){
      return m+n
  }
}
addFive := addN(5)
result := addN(6)

println(result)
}
