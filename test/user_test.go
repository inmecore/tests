package test

import (
	"encoding/json"
	"github.com/agiledragon/gomonkey/v2"
	"testing"
	"time"
)

func TestUserLogin(t *testing.T) {
	patch := gomonkey.ApplyFunc(time.Now, func() time.Time { return now })
	defer patch.Reset()

	type in struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}
	type want struct {
		Id      int    `json:"id"`
		Token   string `json:"token"`
		Message string `json:"message"`
	}
	tests := []struct {
		name   string
		method string
		path   string
		status int
		in     in
		want   want
	}{
		{
			name:   "Login - Success",
			path:   "/api/user/login",
			method: "POST",
			status: 200,
			in:     in{Username: user.Username, Password: user.Password, Code: code},
			want:   want{Id: user.Id, Token: token},
		},
		{
			name:   "Login - Code Illegal",
			path:   "/api/user/login",
			method: "POST",
			status: 400,
			in:     in{Username: user.Username, Password: user.Password, Code: ""},
			want:   want{Message: "Code format verify fail"},
		},
		{
			name:   "Login - Code Illegal",
			path:   "/api/user/login",
			method: "POST",
			status: 400,
			in:     in{Username: user.Username, Password: user.Password, Code: "12345"},
			want:   want{Message: "Code format verify fail"},
		},
		{
			name:   "Login - Invalid Code",
			path:   "/api/user/login",
			method: "POST",
			status: 400,
			in:     in{Username: user.Username, Password: user.Password, Code: "112233"},
			want:   want{Message: "Invalid Code"},
		},
		{
			name:   "Login - Username Not Exist",
			path:   "/api/user/login",
			method: "POST",
			status: 400,
			in:     in{Username: "admin", Password: user.Password, Code: code},
			want:   want{Message: "Invalid username or password"},
		},
		{
			name:   "Login - Password Error",
			path:   "/api/user/login",
			method: "POST",
			status: 400,
			in:     in{Username: user.Username, Password: "", Code: code},
			want:   want{Message: "Password value can not empty"},
		},
		{
			name:   "Login - Password Error",
			path:   "/api/user/login",
			method: "POST",
			status: 400,
			in:     in{Username: user.Username, Password: "123456", Code: code},
			want:   want{Message: "Invalid username or password"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup(false)
			before(c, tt.method, tt.path, tt.in)
			r.ServeHTTP(w, c.Request)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data want
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, tt.want, data)
		})
	}
}

func TestUserInfo(t *testing.T) {
	patch := gomonkey.ApplyFunc(time.Now, func() time.Time { return now })
	defer patch.Reset()

	type in struct {
	}
	type want struct {
		Id          int        `json:"id"`
		Username    string     `json:"username"`
		Phone       string     `json:"phone"`
		Password    string     `json:"password"`
		LastLoginAt *time.Time `json:"lastLoginAt"`
		Message     string     `json:"message"`
	}
	tests := []struct {
		name    string
		method  string
		path    string
		status  int
		logined bool
		in      in
		want    want
	}{
		{
			name:    "Login - Success",
			path:    "/api/user",
			method:  "GET",
			status:  200,
			logined: true,
			in:      in{},
			want: want{
				Id:          1,
				Username:    "test",
				Phone:       "13123456789",
				Password:    "",
				LastLoginAt: &now,
			},
		},
		{
			name:    "Login - Unauthorized",
			path:    "/api/user",
			method:  "GET",
			status:  401,
			logined: false,
			in:      in{},
			want: want{
				Message: "Unauthorized",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup(tt.logined)
			before(c, tt.method, tt.path, tt.in)

			r.ServeHTTP(w, c.Request)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data want
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, tt.want, data)
		})
	}
}
