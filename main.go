package main

import (
    "github.com/book-store/cmd/bookstore"
)

func main() {
    bookstore.Run()

    // $ curl -X POST -H "Content-Type:application/json" \
    // -d '{"id": "978-7-111-55842-2", "name": "The Go Programming Language", "authors":["Alan A.A.Donovan", "Brian W. Kergnighan"],"press": "Pearson Education"}' \
    // localhost:8080/book
    // 2022/01/30 16:28:13 recv a POST request from [::1]:64382

    // $ curl -X GET -H "Content-Type:application/json" localhost:8080/book/978-7-111-55842-2
    /**
      {
          "id": "978-7-111-55842-2",
          "name": "The Go Programming Language",
          "authors": [
              "Alan A.A.Donovan",
              "Brian W. Kergnighan"
          ],
          "press": "Pearson Education"
      }
    */
}
