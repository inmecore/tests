package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"tests/lib"
	"tests/model"
	"tests/router"
)

const (
	loginx = "x-login"
	code   = "123456"
	token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJCdWZmZXJUaW1lIjo4NjQwMCwiVWlkIjoxLCJpc3MiOiJUZXN0ZXIiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTcwNDc2OTIwMCwibmJmIjoxNzA0MTY0Mzk5fQ.dZzYHbJkKiMtNjiBVgOxDhxW8b-HhsbciUDd9KQoipg"
)

var (
	r      = router.New()
	assert = Assert{}
	now    = lib.Now
	user   = model.User{
		Id:          1,
		Username:    "test",
		Phone:       "13123456789",
		Password:    "eYqxuSjMenM=",
		LastLoginAt: &now,
	}
)

func setup(logined bool) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set(loginx, logined)
	return c, w
}

func before(c *gin.Context, method string, path string, data any) {
	request(c, method, path, data)
	if logined, _ := c.Get(loginx); logined == true {
		login(c)
	}
}

func request(c *gin.Context, method string, path string, data any) {
	_data, _ := json.Marshal(data)
	path = params(path, data)
	q := query(data)

	c.Request = httptest.NewRequest(method, path+"?"+q, bytes.NewBuffer(_data))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)
}

// query struct to http query
// key1=1&key2=true&key3=test&key4=%key5=...
func query(data any) string {
	var q string
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		value := v.Field(i)
		q += fmt.Sprintf("%s=%v&", key, value)
	}
	q, _ = strings.CutSuffix(q, "&")
	return q
}

// params struct to http param
// /api/user/:id to /api/user/1
func params(path string, data any) string {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	re := regexp.MustCompile(`:(\w+)`)
	matches := re.FindAllStringSubmatch(path, -1)
	for _, match := range matches {
		for i := 0; i < t.NumField(); i++ {
			key := t.Field(i).Tag.Get("json")
			value := v.Field(i).String()
			if key == match[0][1:] {
				old := fmt.Sprintf("/:%s", key)
				news := fmt.Sprintf("/%s", value)
				path = strings.Replace(path, old, news, 1)
				break
			}
		}
	}
	return path
}

func login(c *gin.Context) {
	c.Request.Header.Add(lib.KeyXUId, strconv.Itoa(user.Id))
	c.Request.Header.Add(lib.KeyXToken, token)
}
