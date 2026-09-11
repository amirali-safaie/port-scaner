package extractor

import (
	"slices"
	"testing"
)

func TestExtractor(t *testing.T) {
	testChann := make(chan string)
	go func() {
		err := ListHosts("127.0.0.0/30", testChann)
		if err != nil {
			t.Error(err)
		}
	}()

	want := []string{
		"127.0.0.1",
		"127.0.0.2",
	}

	got := []string{}
	for ip := range testChann { // this will block until listhosts close the channel
		got = append(got, ip)
	}

	if !slices.Equal(want, got) {
		t.Error("error in list host")
	}

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
			ips := make(chan string)

			go func() {
				err := extractSubHosts(tt.ip, ips)
				if err != nil {
					t.Error(err)
				}
				close(ips)
			}()

			var got []string

			for ip := range ips {
				got = append(got, ip)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
