package hello_test

import (
	"example/hello"
	"testing"
)

func TestTransforms(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		t    hello.Transformer
		args args
		want string
	}{
		{
			name: "NewLine",
			t:    hello.NewLine{},
			args: args{
				s: hello.Message,
			},
			want: hello.Message + "\n",
		},
		{
			name: "Exact",
			t:    hello.Exact{},
			args: args{
				s: hello.Message,
			},
			want: hello.Message,
		},
		{
			name: "Zero",
			t:    hello.Zero{},
			args: args{
				s: hello.Message,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Transform(tt.args.s); got != tt.want {
				t.Errorf("%v.Transform() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
