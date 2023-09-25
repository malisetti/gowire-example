package handlers_test

import (
	"example/handlers"
	"example/hello"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransformServer(t *testing.T) {
	req1, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/?transform=%d&message=%s", hello.ZeroTransform, hello.Message), nil)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "req1",
			args: args{
				w: httptest.NewRecorder(),
				r: req1,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := handlers.InitializeTransformHandler()
			th.ServeHTTP(tt.args.w, tt.args.r)
			got := tt.args.w.Body.String()
			if got != tt.want {
				t.Errorf("response body is wrong, got %q want %q", got, tt.want)
			}
		})
	}
}
