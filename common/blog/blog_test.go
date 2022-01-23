package blog

import "testing"

func TestGlogWriter_Write(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		writer  GlogWriter
		args    args
		wantN   int
		wantErr bool
	}{
		{"", GlogWriter{}, args{[]byte(`log`)}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := GlogWriter{}
			gotN, err := writer.Write(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GlogWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("GlogWriter.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestInitLogs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogs()
		})
	}
}

func TestCloseLogs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseLogs()
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{[]interface{}{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.args...)
		})
	}
}

func TestInfoJSON(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{"%s", []interface{}{"string"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoJSON(tt.args.format, tt.args.args...)
		})
	}
}
