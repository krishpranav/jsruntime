package lib

import (
	"bufio"
	"fmt"
	"os"
)

func Repl() {
	fmt.Printf("JSRuntime %s", VERSION)
	fmt.Println("exit using ctrl-c or exit()")

	for true {
		fmt.Print(">> ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		stringToEval := ""
		fmt.Println(Eval(text, &stringToEval))
		stringToEval += ";undefined;"
	}
}
