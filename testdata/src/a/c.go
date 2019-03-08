package a // want ".+.go is not formated by gofmt"

import (
	"time"
	"fmt"
)

func g() {
	fmt.Println(time.Now())
}
