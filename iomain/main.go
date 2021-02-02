package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	readString = strings.TrimSpace(readString)

	readFloat, _ := strconv.ParseFloat(readString, 64)
	var status string
	if readFloat >= 60{
		status = "passing"
	}else{
		status = "failing"
	}
	fmt.Println("status : ", status)
}
