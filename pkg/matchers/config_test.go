package matchers

import (
	"testing"
)

func TestIsLogformat(t *testing.T) {
	fnTest := "IsLogFormat"
	fnName := "Test" + fnTest
	if IsEnvSet(fnName) {
		IsLogformat(GetTestArgEnv(fnName))
		return
	}

	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid log format",
			args: args{
				text: "",
			},
			wantErr: true,
		},
		{
			name: "no $ sign",
			args: args{
				text: "remote_addr",
			},
			wantErr: true,
		},
		{
			name: "valid log format",
			args: args{
				text: "$remote_addr",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitExecCommand(fnName, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Fatalf(fnTest+"() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestIsToken(t *testing.T) {
	fnTest := "IsToken"
	fnName := "Test" + fnTest
	if IsEnvSet(fnName) {
		IsToken(GetTestArgEnv(fnName))
		return
	}

	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO:
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitExecCommand(fnName, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Fatalf(fnTest+"() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
