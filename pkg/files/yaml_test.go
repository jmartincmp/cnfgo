package files

import "testing"
import "github.com/go-test/deep"

func TestYAMLUnmarshal(t *testing.T) {
	type args struct {
		source        []byte
		configuration interface{}
	}
	tests := []struct {
		name    string
		args    args
		result  *TestStruct
		wantErr bool
	}{
		{
			"Test valid format",
			args{
				[]byte("int: 1\nstring: \"x\"\nnested:\n  int: 2\n  string: \"xx\"\n"),
				&TestStruct{},
			},
			&TestStruct{Int: 1, String: "x", Nested: NestedStruct{Int: 2, String: "xx"}},
			false,
		},
		{
			"Test invalid format",
			args{
				[]byte("int::1"),
				&TestStruct{},
			},
			&TestStruct{Int: 1, String: "x", Nested: NestedStruct{Int: 2, String: "xx"}},
			true,
		},
		{
			"Test empty input",
			args{
				[]byte(""),
				&TestStruct{},
			},
			&TestStruct{Int: 0, String: "", Nested: NestedStruct{Int: 0, String: ""}},
			true,
		},
		{
			"Test invalid stuct pointer",
			args{
				[]byte(""),
				TestStruct{},
			},
			&TestStruct{Int: 0, String: "", Nested: NestedStruct{Int: 0, String: ""}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yf := NewYAMLFormatUnmarshaler()
			if err := yf.Unmarshal(tt.args.source, tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("YAMLFormat.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if diff := deep.Equal(tt.args.configuration, tt.result); diff != nil {
				t.Error(diff)
			}
		})
	}
}
