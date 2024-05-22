package lib

import (
	"encoding/json"
)

// Convert 类型转换，通过标签json进行中间转换
// from 源数据，
// to 目标数据指针
func Convert(from interface{}, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, to)
	if err != nil {
		return err
	}
	return nil
}
