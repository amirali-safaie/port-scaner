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




func TestSubnetMask(t *testing.T) {
	tests := []struct {
		name    string
		ip      string
		want    string
		wantErr bool
	}{
		{
			name:    "CIDR /23",
			ip:      "192.168.0.1/23",
			want:    "11111111.11111111.11111110.00000000",
			wantErr: false,
		},
		{
			name:    "CIDR /24",
			ip:      "192.168.1.10/24",
			want:    "11111111.11111111.11111111.00000000",
			wantErr: false,
		},
		{
			name:    "CIDR /8",
			ip:      "10.0.0.1/8",
			want:    "11111111.00000000.00000000.00000000",
			wantErr: false,
		},
		{
			name:    "missing CIDR prefix",
			ip:      "192.168.1.1",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subnetMask(tt.ip)

			if (err != nil) != tt.wantErr {
				t.Fatalf("SubnetMask() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("ip = %s SubnetMask() = %q, want %q",tt.ip, got, tt.want)
			}
		})
	}
}


// "11111111.11111111.11111111.11111110.00000000", 
// want "11111111.11111111.11111110.00000000"
