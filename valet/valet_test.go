package valet_test

import (
	"bytes"
	"example/hello"
	"example/internal"
	"example/internal/config"
	"example/valet"
	"testing"
)

func TestValetPark_Park(t *testing.T) {
	type args struct {
		v valet.Vehicle
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		trasform hello.TransformType
	}{
		{
			name: "Test Valet1",
			args: args{
				v: valet.Car{},
			},
			trasform: hello.ZeroTransform,
			wantErr:  false,
		},
		{
			name: "Test Valet2",
			args: args{
				v: valet.Car{},
			},
			trasform: hello.NewLineTransform,
			wantErr:  false,
		},
		{
			name: "Test Valet3",
			args: args{
				v: valet.Car{},
			},
			trasform: hello.ExactTransform,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer
			cfg := config.GetCfg()
			vp := internal.InitializeValetParker(&buf, cfg, hello.Message)
			if err := vp.Park(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ValetPark.Park() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
