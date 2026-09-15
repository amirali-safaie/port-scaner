package presenter

import (
	"fmt"
	"port-scaner/scaner"
)

type Presenter interface {
	Present(result scaner.Result)
}

type Terminal struct{}

func (t Terminal) Present(result scaner.Result) {
	st := ""
	if result.STATUS == 0 {
		st = "CLOSE"
	} else {
		st = "OPEN"
	}

	fmt.Printf("ip : %s and port : %d is --> %s", result.IP, result.PORT, st)
	fmt.Println()
}
