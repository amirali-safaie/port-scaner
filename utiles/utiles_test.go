package utiles

import (
	"testing"
)

func TestDec2Bi(t *testing.T) {
	tests := []struct {
		num int
		expected string
	}{
		{10, "00001010"},
		{0, "00000000"},
		{5, "00000101"},
		{15, "00001111"},
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




func TestBaseIp(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want string
	}{
		{
			name: "192.168.0.1/23",
			ip:   "192.168.0.1/23",
			want: "192.168.0.0",
		},
		{
			name: "192.168.1.100/24",
			ip:   "192.168.1.100/24",
			want: "192.168.1.0",
		},
		{
			name: "10.20.30.40/16",
			ip:   "10.20.30.40/16",
			want: "10.20.0.0",
		},
		{
			name: "172.16.50.100/8",
			ip:   "172.16.50.100/8",
			want: "172.0.0.0",
		},
		{
			name: "already base IP",
			ip:   "192.168.10.0/24",
			want: "192.168.10.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BaseIp(tt.ip)

			if err != nil {
				t.Fatalf("baseIp(%q) returned unexpected error: %v", tt.ip, err)
			}

			if got != tt.want {
				t.Errorf("baseIp(%q) = %q, want %q", tt.ip, got, tt.want)
			}
		})
	}
}



func TestEndIp(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want string
	}{
		{
			name: "192.168.0.1/23",
			ip:   "192.168.0.1/23",
			want: "192.168.1.255",
		},
		{
			name: "192.168.1.100/24",
			ip:   "192.168.1.100/24",
			want: "192.168.1.255",
		},
		{
			name: "10.20.30.40/16",
			ip:   "10.20.30.40/16",
			want: "10.20.255.255",
		},
		{
			name: "172.16.50.100/8",
			ip:   "172.16.50.100/8",
			want: "172.255.255.255",
		},
		{
			name: "already end IP",
			ip:   "192.168.10.255/24",
			want: "192.168.10.255",
		},
		{
			name: "10.0.0.1/30",
			ip:   "10.0.0.3/30",
			want: "10.0.0.3",
		},
		{
			name: "10.0.0.1/32",
			ip:   "10.0.0.1/32",
			want: "10.0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EndIp(tt.ip)

			if err != nil {
				t.Fatalf("endIp(%q) returned unexpected error: %v", tt.ip, err)
			}

			if got != tt.want {
				t.Errorf("endIp(%q) = %q, want %q", tt.ip, got, tt.want)
			}
		})
	}
}

