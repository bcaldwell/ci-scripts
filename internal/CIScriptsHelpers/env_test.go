package CIScriptsHelpers

import "testing"

func TestEnvCheck(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EnvCheck(tt.args.key, tt.args.value)
		})
	}
}

func TestRequiredEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RequiredEnv(tt.args.key)
		})
	}
}

func TestEnvFetch(t *testing.T) {
	type args struct {
		key          string
		defaultValue []string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantOk    bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := EnvFetch(tt.args.key, tt.args.defaultValue...)
			if gotValue != tt.wantValue {
				t.Errorf("EnvFetch() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("EnvFetch() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
