package test

import (
	"testing"
	"tests/lib"
)

func TestEncode(t *testing.T) {
	type in struct {
		Pass string
		Key  string
	}
	type want struct {
		Hash string
	}
	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "Encode And Decode",
			in:   in{Pass: lib.DefaultPasswordVal, Key: lib.KeyEncode},
			want: want{Hash: "eYqxuSjMenM="},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := lib.Encode(tt.in.Pass, tt.in.Key)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.Hash, hash)
			pass, err := lib.Decode(hash, tt.in.Key)
			assert.Nil(t, err)
			assert.Equal(t, tt.in.Pass, pass)
		})
	}
}

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Password 123456",
			in:   lib.DefaultPasswordVal,
			want: lib.DefaultPasswordHash,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := lib.Password(tt.in)
			assert.Equal(t, tt.want, hash)
		})
	}
}
