package matchers

import (
	"flag"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"ktbs.dev/teler/pkg/errors"
)

func TestIsLogformat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid log format",
			args: args {
				s: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid log format: need $ sign",
			args: args {
				s: "remote_addr",
			},
			wantErr: true,
		},
		{
			name: "Log format is valid",
			args: args {
				s: "$request_method",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var code int
			errors.Abort = func(c int) {
				code = c
			}
			IsLogformat(tt.args.s)
			if (code != 0) != tt.wantErr {
				t.Fatalf("IsLogFormat() error code: %v wantErr: %v", code, tt.wantErr)
			}
		})
	}
}

func TestIsToken(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid token",
			args: args {
				s: "hola",
			},
			wantErr: true,
		},
		{
			name: "Token is valid",
			args: args {
				s: "xoxp-0123456789-012345678901-y",
			},
			wantErr: false,
		},
		{
			name: "Token is valid",
			args: args {
				s: "012345678:abcdefghijklmnopqrstuvwxyz012345678",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var code int
			errors.Abort = func(c int) {
				code = c
			}
			IsToken(tt.args.s)
			if (code != 0) != tt.wantErr {
				t.Fatalf("IsToken() error code: %v wantErr: %v", code, tt.wantErr)
			}
		})
	}
}

func TestIsHexcolor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid hex color code",
			args: args {
				s: "\u2713",
			},
			wantErr: true,
		},
		{
			name: "Invalid hex color code",
			args: args {
				s: "#af01",
			},
			wantErr: true,
		},
		{
			name: "Hex color code is valid",
			args: args {
				s: "#F00000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var code int
			errors.Abort = func(c int) {
				code = c
			}
			IsHexcolor(tt.args.s)
			if (code != 0) != tt.wantErr {
				t.Fatalf("IsHexColor() error code: %v wantErr: %v", code, tt.wantErr)
			}
		})
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
}

func randomChannelID() string {
	numStr := "0123456789"
	minLn := 9
	maxLn := 13
	ln := minLn + rand.Intn(maxLn-minLn)
	lenNum := len(numStr)
	var res strings.Builder
	for i := 1; i <= ln; i++ {
		res.WriteByte(numStr[rand.Intn(lenNum)])
	}
	return res.String()
}

func TestIsChannel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid Slack channel ID",
			args: args {
				s: "hola",
			},
			wantErr: true,
		},
		{
			name: "Slack channel ID is valid",
			args: args {
				s: randomChannelID(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var code int
			errors.Abort = func(c int) {
				code = c
			}
			IsChannel(tt.args.s)
			if (code != 0) != tt.wantErr {
				t.Fatalf("IsChannel() error code: %v wantErr: %v", code, tt.wantErr)
			}
		})
	}
}

func TestIsChatID(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid Telegram chat_id",
			args: args {
				s: "hola",
			},
			wantErr: true,
		},
		{
			name: "Telegram chat_id is valid",
			args: args {
				s: "12345678",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var code int
			errors.Abort = func(c int) {
				code = c
			}
			IsChatID(tt.args.s)
			if (code != 0) != tt.wantErr {
				t.Fatalf("IsChatID() error code: %v wantErr: %v", code, tt.wantErr)
			}
		})
	}
}
