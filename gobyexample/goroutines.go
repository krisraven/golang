package main
// when we run this program we see the outputof the two blocking calls first then the interleaved output of the two goroutines
import(
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
// function call f(s)
func main() {
	f("direct") // this would run the function call synchronously

	go f("goroutine") // to invoke this function in a goroutine use go f(s). This will execute concurrently with the calling one

	// it is also possible to start a goroutine for an annonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // our two function calls are running asynchronously in separate goroutines now
	fmt.Println("done")
}