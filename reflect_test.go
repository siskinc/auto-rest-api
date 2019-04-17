package autorestapi

import "testing"

type demo struct {
}

func Test_getStructName(t *testing.T) {
	type args struct {
		t interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantName string
	}{
		{
			"test1",
			args{
				&demo{},
			},
			"demo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName := getStructName(tt.args.t); gotName != tt.wantName {
				t.Errorf("getStructName() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
