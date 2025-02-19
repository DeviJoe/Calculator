package types

import (
	"sync"
	"testing"
)

func TestVariable_SetValue(t *testing.T) {
	type fields struct {
		name           string
		value          int64
		isPrintable    bool
		isValueSetting bool
	}
	type args struct {
		value int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SetValue to New Variable",
			fields: fields{
				name:           "x",
				value:          0,
				isPrintable:    false,
				isValueSetting: false,
			},
			args: args{
				value: 20,
			},
			wantErr: false,
		},
		{
			name: "SetValue to existing Variable",
			fields: fields{
				name:           "x",
				value:          35,
				isPrintable:    false,
				isValueSetting: true,
			},
			args: args{
				value: 20,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Variable{
				name:           tt.fields.name,
				value:          tt.fields.value,
				isPrintable:    tt.fields.isPrintable,
				isValueSetting: tt.fields.isValueSetting,
				mu:             &sync.Mutex{},
			}
			if err := v.SetValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SetValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
