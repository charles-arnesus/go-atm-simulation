package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func login() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Account Number:")
	fmt.Print("->")

	accountNumber, _ := reader.ReadString('\n')
	accountNumber = strings.Replace(accountNumber, "\n", "", -1)

}
