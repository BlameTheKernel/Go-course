package main
 
import (
   "fmt"
)
 
func main() {
  str := "This is a string"
  fmt.Println(str)
 
  // get first 10 chars
  first10 :=  str[0:9]
  fmt.Println(first10)
}
