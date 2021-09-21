// help.go
package champlib

import "fmt"

var helpMsg = [3][2]string{
	{"help", "Print this message and exit"},
	{"debug", "Enable logging (more info)"},
	{"file", "[argument] Specify where to write log output"},
}

func PrintHelp() {
	for i := 0; i < len(helpMsg); i++ {
		fmt.Printf("--%s\t", helpMsg[i][0])
		fmt.Println(helpMsg[i][1])
	}
}
