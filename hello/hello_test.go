package hello_test

import (
	"bytes"
	"example/hello"
	"testing"
)

func TestSay(t *testing.T) {
	type args struct {
		sayer hello.Sayer
		msg   string
	}
	var buf bytes.Buffer
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "NewLine Sayer",
			args: args{
				sayer: hello.NewSayer(&buf, &hello.NewLine{}),
				msg:   hello.Message,
			},
			want: hello.Message + "\n",
		},
		{
			name: "Exact Sayer",
			args: args{
				sayer: hello.NewSayer(&buf, &hello.Exact{}),
				msg:   hello.Message,
			},
			want: hello.Message,
		},
		{
			name: "No Sayer",
			args: args{
				sayer: hello.NewSayer(&buf, &hello.Zero{}),
				msg:   hello.Message,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.args.sayer.Say(tt.args.msg)
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
