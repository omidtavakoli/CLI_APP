package main

import (
	"errors"
	"os"
	"testing"
)

func TestCLI(t *testing.T) {
	//	todo: run some cli tests here
}

func TestFunctionality(t *testing.T) {
	os.Args = []string{"greet", "--name", "Jeremy"}

	tests := []struct {
		description    string
		name           string
		req            string
		expectedErrors error
		want           User
	}{
		{
			description:    "test",
			name:           "success",
			expectedErrors: errors.New("sth wrong"),
			want: User{
				ID:       0,
				UserName: "Omid",
			},
		},
		{
			description:    "test1",
			name:           "success",
			expectedErrors: errors.New("sth wrong"),
			want: User{
				ID:       0,
				UserName: "Omid1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			err := Calculation()
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}
