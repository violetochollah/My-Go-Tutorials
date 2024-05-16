package main
import "fmt"


type shaper interface {
Area() float32
}
type square struct  {
side float32
}

func (sq *square) Area() float32 {
return sq.side *sq.side
}

func main() {
sq1 :=new(square)
sq1.side = 5

//var areaIntf shaper
//areaIntf = sq1
//shorter,without separate declaration
//areaIntf := shaper(sq1)
//or even:
areaIntf := sq1
fmt.Printf("the square has: %f\n",areaIntf.Area())
}
