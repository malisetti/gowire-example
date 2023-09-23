package hello_test

import (
	"bytes"
	"example/hello"
	"example/internal"
	"example/internal/config"
	"testing"
)

func TestSay(t *testing.T) {
	type args struct {
		msg string
	}
	var buf bytes.Buffer
	cfg := config.GetCfg()
	tests := []struct {
		name      string
		transform hello.TransformType
		args      args
		want      string
	}{
		{
			name: "NewLine Sayer",
			args: args{
				msg: hello.Message,
			},
			transform: hello.NewLineTransform,
			want:      hello.Message + "\n",
		},
		{
			name: "Exact Sayer",
			args: args{
				msg: hello.Message,
			},
			transform: hello.ExactTransform,
			want:      hello.Message,
		},
		{
			name: "No Sayer",
			args: args{
				msg: hello.Message,
			},
			transform: hello.ZeroTransform,
			want:      "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.TransformerProviderType = tt.transform
			sayer := internal.InitializeSayer(&buf, cfg)
			_, err := sayer.Say(tt.args.msg)
			if err != nil {
				t.Fatalf("could not say \"%s\", failed with %v", hello.Message, err)
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("Say() = %v, want %v", got, tt.want)
			}
			buf.Reset()
		})
	}
}
