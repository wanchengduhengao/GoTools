package utils

import (
	"encoding/json"
	"os"
	"path"
	"reflect"
	"runtime"
)

type DBSetting struct {
	Engine   string `json:"engine"`
	DBName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type Setting struct {
	Database *DBSetting `json:"database"`
}

func load(setting interface{}, path string) {
	if reflect.TypeOf(setting).Elem().Kind() != reflect.Struct {
		panic("setting mast is a struct")
	}
	fp, err := os.Open(path)
	if err != nil {
		ExceptionLog(err, "Fail to open setting")
		panic(err)
	}
	defer fp.Close()
	decoder := json.NewDecoder(fp)
	err = decoder.Decode(setting)
	if err != nil {
		ExceptionLog(err, "Fail to decode json setting")
		panic(err)
	}
}

func InitSetting(s interface{}, fPath string) interface{} {
	_, currently, _, _ := runtime.Caller(1)
	filename := path.Join(path.Dir(currently), fPath)
	load(&s, filename)
	return s
}
