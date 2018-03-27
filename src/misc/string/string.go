package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "host:port"
	arr := strings.Split(str, ":")
	fmt.Println(arr)
}
