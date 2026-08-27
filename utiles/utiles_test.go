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

func TestAdd2Ipv4(t *testing.T) {
	tests := []struct {
		ip       string
		expected string
		wantErr  bool
	}{
		{"0.0.0.0", "0.0.0.1", false},
		{"192.168.1.1", "192.168.1.2", false},
		{"192.168.1.255", "192.168.2.0", false},
		{"192.168.255.255", "192.169.0.0", false},
		{"255.255.255.255", "", true},
	}

	for _, test := range tests {
		got, err := Add2Ipv4(test.ip)
		if (err != nil) != test.wantErr {
			t.Errorf("Add2Ipv4(%s) error = %v; wantErr %v", test.ip, err, test.wantErr)
		}
		if !test.wantErr && got != test.expected {
			t.Errorf("Add2Ipv4(%s) = %s; expected %s", test.ip, got, test.expected)
		}
	}
}