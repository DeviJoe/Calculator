package types

import (
	"testing"
)

func TestNewOperation(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name    string
		args    args
		want    Operation
		wantErr bool
	}{
		{
			name: "Correct symbol plus(+)",
			args: args{
				sign: "+",
			},
			want:    plus,
			wantErr: false,
		},
		{
			name: "Correct symbol minus(-)",
			args: args{
				sign: "-",
			},
			want:    minus,
			wantErr: false,
		},
		{
			name: "Correct symbol multiplication(*)",
			args: args{
				sign: "*",
			},
			want:    multiplication,
			wantErr: false,
		},
		{
			name: "Incorrect symbol",
			args: args{
				sign: "/",
			},
			want:    0,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOperation(tt.args.sign)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewOperation() got = %v, want %v", got, tt.want)
			}
		})
	}
}
