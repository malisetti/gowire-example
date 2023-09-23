package valet_test

import (
	"example/hello"
	"example/valet"
	"testing"
)

func TestValetPark_Park(t *testing.T) {
	type fields struct {
		helloSayer hello.Sayer
	}
	type args struct {
		v valet.Vehicle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Valet1",
			fields: fields{
				helloSayer: hello.NewSayer(hello.Exact{}),
			},
			args: args{
				v: valet.Car{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vp := valet.NewValetParker(tt.fields.helloSayer)
			if err := vp.Park(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ValetPark.Park() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
