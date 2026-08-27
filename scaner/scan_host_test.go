package scaner

import (
	"fmt"
	"testing"
)


func TestExtractor(t *testing.T) {
	got, err := ListHosts("127.0.0.1/23")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Print(got)
}