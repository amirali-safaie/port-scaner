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
		got, err := Dec2Bi(test.num)
		if err != nil{
			t.Errorf("faild!")
		}
		if got != test.expected {
			t.Errorf("Dec2bi(%d) = %s; expected %s", test.num, got, test.expected)
		}
	}

}


func TestBi2Dec(t *testing.T){
	got, err := Bi2Dec("101")
	expect := 5
	if err != nil {
		t.Errorf("faild to run")
	}
	if got != expect {
		t.Errorf("expected %d but got %d",expect, got)
	}
}