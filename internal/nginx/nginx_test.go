package nginx

import (
	"testing"
)

func TestGetNginxCommand(t *testing.T) {
	tests := []struct {
		cmd             string
		nginxBinaryPath string
		expected        string
	}{
		{"reload", "/usr/sbin/nginx", "/usr/sbin/nginx -s reload"},
		{"stop", "/usr/sbin/nginx-debug", "/usr/sbin/nginx-debug -s stop"},
	}
	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			nginx := NewNginxController("/etc/nginx", tt.nginxBinaryPath, true)

			if got := nginx.getNginxCommand(tt.cmd); got != tt.expected {
				t.Errorf("getNginxCommand returned \n%v, but expected \n%v", got, tt.expected)
			}
		})
	}
}
