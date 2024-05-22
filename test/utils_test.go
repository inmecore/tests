package test

import (
	"testing"
	"tests/lib"
)

func TestQuery(t *testing.T) {
	type in struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Lock bool   `json:"lock"`
	}
	type want string
	var w want = "id=1&name=test&lock=false"
	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "Encode And Decode",
			in:   in{Id: 1, Name: "test", Lock: false},
			want: w,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := query(tt.in)
			assert.Equal(t, string(tt.want), q)
		})
	}
}

func TestConvert(t *testing.T) {
	type base struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Lock bool   `json:"lock"`
	}
	type source struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Lock bool   `json:"lock"`
	}
	type in base
	type want source
	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "Encode And Decode",
			in:   in{Id: 1, Name: "test", Lock: false},
			want: want{Id: 1, Name: "test", Lock: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i want
			err := lib.Convert(tt.in, &i)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, i)
		})
	}
}
