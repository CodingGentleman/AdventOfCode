package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	val := "bgvyzdsv"
	for i := 0; true; i++ {
		if strings.HasPrefix(fmt.Sprintf("%x", md5.Sum([]byte(val+strconv.Itoa(i)))), "00000") {
			fmt.Println(i)
			return
		}
	}
}
