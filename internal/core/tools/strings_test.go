package tools

import (
	"reflect"
	"testing"
)

func TestLowerCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "userName"},
			want: "userName",
		},
		{
			args: args{str: "UserName"},
			want: "userName",
		},
		{
			args: args{str: "user_name"},
			want: "userName",
		},
		{
			args: args{str: "User_Name"},
			want: "userName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCamelCase(tt.args.str); got != tt.want {
				t.Errorf("LowerCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "userName"},
			want: "UserName",
		},
		{
			args: args{str: "UserName"},
			want: "UserName",
		},
		{
			args: args{str: "user_name"},
			want: "UserName",
		},
		{
			args: args{str: "User_Name"},
			want: "UserName",
		},
		{
			args: args{str: "userName_Table"},
			want: "UserNameTable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperCamelCase(tt.args.str); got != tt.want {
				t.Errorf("UpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{str: "userName"},
			want: []string{"user", "Name"},
		},
		{
			args: args{str: "UserName"},
			want: []string{"User", "Name"},
		},
		{
			args: args{str: "user_name"},
			want: []string{"user", "name"},
		},
		{
			args: args{str: "User_Name"},
			want: []string{"User", "Name"},
		},
		{
			args: args{str: "userName_Table"},
			want: []string{"user", "Name", "Table"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
