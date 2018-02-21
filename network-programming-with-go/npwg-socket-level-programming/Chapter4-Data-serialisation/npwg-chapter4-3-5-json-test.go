package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	ret, err := json.Marshal("[{Name: fred, Occupation: programmer},{Name: liping, Occupation: analyst},{Name: sureerat, Occupation: manager}]")
	fmt.Println([][]byte{ret})
}
