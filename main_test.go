package main

import "testing"

func TestGetUTFLength(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Valid UTF-8",
			args: args{
				input: []byte("Hello, world!"),
			},
			want:    13,
			wantErr: false,
		},
		{
			name: "Invalid UTF-8",
			args: args{
				input: []byte{0xFF, 0xFE, 0x00, 0x00, 0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x77, 0x6F, 0x72, 0x6C, 0x64, 0x21},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Empty string",
			args: args{
				input: []byte(""),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUTFLength(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUTFLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUTFLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
