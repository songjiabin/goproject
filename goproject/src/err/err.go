package err

import (
	"fmt"
	"os"
)

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func Test()  {
	fmt.Println("Test a good person")
}






