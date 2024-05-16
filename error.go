//in this code i implement panic,recover&defer package
package main
import (
"fmt"

)
func badCall()  {
panic("bad end")
}
func test() {
defer func() {
if e :=recover(); e !=nil {
fmt.Printf("panicking %s\r\n", e);
}
}()
badCall()
fmt.  Printf("After bad Call\r\n");
}
func main() {
  fmt.Printf("Calling test\r\n");
  test()
  fmt.Printf("Test completed\r\n");
}
