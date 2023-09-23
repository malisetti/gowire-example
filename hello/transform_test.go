package hello_test

import (
	"example/hello"
	"example/internal/config"
	"strings"
	"testing"
)

func TestTransforms(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		transform hello.TransformType
		want      string
	}{
		{
			name: "NewLine",
			args: args{
				s: hello.Message,
			},
			transform: hello.NewLineTransform,
			want:      hello.Message + "\n",
		},
		{
			name: "Exact",
			args: args{
				s: hello.Message,
			},
			transform: hello.ExactTransform,
			want:      hello.Message,
		},
		{
			name: "Zero",
			args: args{
				s: hello.Message,
			},
			transform: hello.ZeroTransform,
			want:      "",
		},
		{
			name: "Upper",
			args: args{
				s: hello.Message,
			},
			transform: hello.UpperTransform,
			want:      strings.ToUpper(hello.Message),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cfg := config.GetCfg()
			cfg.TransformerProviderType = tt.transform
			tr := hello.NewTransform(cfg)
			if got := tr.Transform(tt.args.s); got != tt.want {
				t.Errorf("%v.Transform() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
