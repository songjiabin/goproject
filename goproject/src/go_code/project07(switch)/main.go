package main
import (
	"fmt"
	"runtime"
	"time"
)

// swithc 
func main()  {
	// os:=runtime.GOOS
	switch os:=runtime.GOOS; os {
		case "windows":
			fmt.Println("window")
		case "linux":
			fmt.Println("linux")
		default:
			fmt.Printf("%s",os) 
	}


	num:=0
	switch num {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	}


	//没有条件的switch 就和switch true 一样。

	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}




}