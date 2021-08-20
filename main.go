package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MRBednar/GoAdvent2020/s3connect"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")

	text, _ := reader.ReadString('\n')

	cleanText := strings.TrimSuffix(text, "\n")

	input := s3connect.Connect(cleanText)

	fmt.Println(input)
}
