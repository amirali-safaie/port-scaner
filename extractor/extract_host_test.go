package extractor

import (
	"fmt"
	"testing"
)

func TestExtractor(t *testing.T) {
	got, err := ListHosts("127.0.0.1/23")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(got)
}

func TestExtractSubHosts(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want []string
	}{
		{
			name: "192.168.1.0/29",
			ip:   "192.168.1.0/29",
			want: []string{
				"192.168.1.1",
				"192.168.1.2",
				"192.168.1.3",
				"192.168.1.4",
				"192.168.1.5",
				"192.168.1.6",
			},
		},
		{
			name: "192.168.1.0/30",
			ip:   "192.168.1.0/30",
			want: []string{
				"192.168.1.1",
				"192.168.1.2",
			},
		},
		{
			name: "10.0.0.0/29",
			ip:   "10.0.0.0/29",
			want: []string{
				"10.0.0.1",
				"10.0.0.2",
				"10.0.0.3",
				"10.0.0.4",
				"10.0.0.5",
				"10.0.0.6",
			},
		},
		{
			name: "192.168.10.0/30",
			ip:   "192.168.10.0/30",
			want: []string{
				"192.168.10.1",
				"192.168.10.2",
			},
		},
		{
			name: "small subnet /31",
			ip:   "192.168.1.0/31",
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractSubHosts(tt.ip)

			if err != nil {
				t.Fatalf(
					"extractSubHosts(%q) returned unexpected error: %v",
					tt.ip,
					err,
				)
			}

			if len(got) != len(tt.want) {
				t.Fatalf(
					"extractSubHosts(%q) returned %d hosts, want %d",
					tt.ip,
					len(got),
					len(tt.want),
				)
			}

			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf(
						"extractSubHosts(%q)[%d] = %q, want %q",
						tt.ip,
						i,
						got[i],
						tt.want[i],
					)
				}
			}
		})
	}
}
