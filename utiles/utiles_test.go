package utiles
import (
	"testing"
)

func TestDec2Bi(t *testing.T) {
	tests := []struct {
		num int
		expected string
	}{
		{10, "1010"},
		{0, ""},
		{5, "101"},
		{15, "1111"},
	}

	for _, test := range tests {
		got, err := Dec2bi(test.num)
		if err != nil{
			t.Errorf("faild!")
		}
		if got != test.expected {
			t.Errorf("Dec2bi(%d) = %s; expected %s", test.num, got, test.expected)
		}
	}

}
