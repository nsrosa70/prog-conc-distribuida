package main
import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {

	// mouse events: mleft, mright, wheelDown, wheelUp, wheelLeft, wheelRight.
	for {
		mleft := robotgo.AddEvent("mright")
		if mleft {
			fmt.Println("you press... ", "mouse right button")
		}
	}
}
