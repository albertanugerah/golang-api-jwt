package helper

import (
	"AP/golang-api-jwt/entity"
	"go/types"
	"reflect"
	"testing"
)

func TestBuildResponse(t *testing.T) {
	type args struct {
		status  bool
		message string
		error types.Nil
		data    entity.Profile
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name : "response",
			args: args{
				status:  true,
				message: "OK",
				data:  entity.Profile{
					ID:      1,
					Name:    "Albert",
					Address: "Kreo",
					City:    "Tangerang",
					UserID:  1,
					User:    entity.User{
						ID:       1,
						Email:    "albertanugerah@hotamil.com",
						Token:    "safasfdsafas",
					},
				},
			},
			want: BuildResponse(true,"OK",entity.Profile{
				ID:      1,
				Name:    "Albert",
				Address: "Kreo",
				City:    "Tangerang",
				UserID:  1,
				User:    entity.User{
					ID:       1,
					Email:    "albertanugerah@hotamil.com",
					Token:    "safasfdsafas",
				},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildResponse(tt.args.status, tt.args.message, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildErrorResponse(t *testing.T) {
	type args struct {
		message string
		err     string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "error response",
			args: args{
				message: "gagal",
				err:     "error1\nerror2\n",
				data:   nil,
			},
			want: BuildErrorResponse("gagal","error1\nerror2\n",nil),
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildErrorResponse(tt.args.message, tt.args.err, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}