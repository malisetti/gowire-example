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
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "NewLine Sayer",
			args: args{
				sayer: &hello.NewLineSayer{},
				msg:   hello.Message,
			},
			want: hello.Message + "\n",
		},
		{
			name: "Exact Sayer",
			args: args{
				sayer: &hello.ExactSayer{},
				msg:   hello.Message,
			},
			want: hello.Message,
		},
		{
			name: "No Sayer",
			args: args{
				sayer: &hello.NoSayer{},
				msg:   hello.Message,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			_, err := tt.args.sayer.Say(w, tt.args.msg)
			if err != nil {
				t.Fatalf("could not say %s, failed with %v", hello.Message, err)
			}
			if gotW := w.String(); gotW != tt.want {
				t.Errorf("Say() = %v, want %v", gotW, tt.want)
			}
		})
	}
}
