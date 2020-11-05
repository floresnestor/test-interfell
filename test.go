package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

const spc = " "
const quote = "\""
const msgrt = "Return:"
const msgInput = "Enter text:"

func main() {
	var reversedStr string = reverse(capBeforeNewline());
	fmt.Println(msgrt + quote + reversedStr + quote)
}

func reverse(str string) string {
	s := strings.Fields(str)
	var rta string = ""
	for i := (len(s)-1); i >= 0; i-- {
		rta = rta + s[i] + spc
	}
	last := len(rta) - 1;
	if last >= 0 {
        rta = rta[:last]
    }
	return(strings.Trim(rta, "\t \n"))
}

func capBeforeNewline() string {
	capture := bufio.NewReader(os.Stdin)
	fmt.Print(msgInput)
	text, _ := capture.ReadString('\n')
	return text
}
