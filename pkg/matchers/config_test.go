package matchers

import (
	"testing"

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
			name: "invalid log format",
			args: args{
				s: "",
			},
			wantErr: true,
		},
		{
			name: "no $ sign",
			args: args{
				s: "remote_addr",
			},
			wantErr: true,
		},
		{
			name: "nonexist required log format",
			args: args{
				s: "$upstream_addr",
			},
			wantErr: true,
		},
		{
			name: "valid log format",
			args: args{
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
			name: "invalid token",
			args: args{
				s: "hola",
			},
			wantErr: true,
		},
		{
			name: "valid slack token",
			args: args{
				s: "xoxp-0123456789-012345678901-y",
			},
			wantErr: false,
		},
		{
			name: "valid telegram token",
			args: args{
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
			name: "invalid hex color",
			args: args{
				s: "\u2713",
			},
			wantErr: true,
		},
		{
			name: "invalid hex syntax",
			args: args{
				s: "#af01",
			},
			wantErr: true,
		},
		{
			name: "valid hex color",
			args: args{
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

func TestIsParseMode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid parse mode",
			args: args{
				s: "Latex",
			},
			wantErr: true,
		},
		{
			name: "outdated markdown",
			args: args{
				s: "MarkdownV1",
			},
			wantErr: true,
		},
		{
			name: "case sensitive invalid html",
			args: args{
				s: "html",
			},
			wantErr: true,
		},
		{
			name: "valid markdown",
			args: args{
				s: "MarkdownV2",
			},
			wantErr: false,
		},
		{
			name: "valid html",
			args: args{
				s: "HTML",
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
			IsParseMode(tt.args.s)
			if (code != 0) != tt.wantErr {
				t.Fatalf("IsParseMode() error code: %v wantErr: %v", code, tt.wantErr)
			}
		})
	}
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
			name: "invalid channel format",
			args: args{
				s: "hola",
			},
			wantErr: true,
		},
		{
			name: "invalid lowercase channel name",
			args: args{
				s: "kitabisa13",
			},
			wantErr: true,
		},
		{
			name: "valid channel format",
			args: args{
				s: "KITABISA13",
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
			name: "invalid chat id",
			args: args{
				s: "hola",
			},
			wantErr: true,
		},
		{
			name: "digit chat id",
			args: args{
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
