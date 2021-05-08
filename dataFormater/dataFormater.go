package dataFormater

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
)

func Text() {
	fmt.Println("hello")
}

type RiskData struct {
	rawData    sync.Map // interface{}
	config     sync.Map // ConfigJsonTpl
	intType    sync.Map // int
	stringType sync.Map // string
}

// 原始数据类型
type rawType map[string]interface{}
type IntType int
type StringType string

// 标记获取数据类型
const TypeInt = "int"
const TypeString = "string"

// ConfigJsonTpl 字段配置 json 模板
type ConfigJsonTpl struct {
	KeyName  string        `json:"key_name"`
	FuncName string        `json:"func_name"`
	DataType string        `json:"type"`
	Args     []interface{} `json:"args"`
}

// LoadConfig 加载配置
func (rd *RiskData) LoadConfig(configJson string) error {
	cf := []ConfigJsonTpl{}
	// err := json.Unmarshal([]byte(configJson), &cf)
	err := JsonDecode(configJson, &cf, true)

	if err != nil {
		return err
	}

	for i := 0; i < len(cf); i++ {
		rd.config.Store(cf[i].KeyName, cf[i])
	}
	return nil
}

// 通过 key 获取配置
func (rd *RiskData) GetConf(k string) (cf ConfigJsonTpl, ok bool) {
	t, ok := rd.config.Load(k)
	if ok {
		cf = t.(ConfigJsonTpl)
	}
	return
}

// JsonDecode 通用的 json 解析方法
func JsonDecode(raw string, tpl interface{}, useNumber bool) error {
	if useNumber {
		d := json.NewDecoder(bytes.NewReader([]byte(raw)))
		d.UseNumber()
		return d.Decode(tpl)
	}

	return json.Unmarshal([]byte(raw), &tpl)
}

// LoadData 加载数据, 并将数据按配置好的类型进行转换
func (rd *RiskData) LoadData(raw string) (err error) {
	tpl := rawType{}
	err = JsonDecode(raw, &tpl, true)
	if err != nil {
		return err
	}

	// 按字段配置将数据分配到各类型的 map 中
	err = rd.swithcType(tpl)
	if err != nil {
		return err
	}
	return nil
}

// 将传入的数据从 interface{} 转成指定类型
// 需要函数处理的新字段不在这里处理
func (rd *RiskData) swithcType(data map[string]interface{}) (err error) {

	rd.config.Range(func(k, v interface{}) bool {
		// log.Println("iterate:", k, v)
		confKey := k.(string)
		kConfig := v.(ConfigJsonTpl)

		// 配置中带 func 的稍后处置
		if kConfig.FuncName != "" || len(kConfig.FuncName) > 0 {
			return true
		}

		// 通过配置判断外界传入数据是否允许, 非配置中的给过滤掉
		rawValue, exist := data[confKey]
		if !exist {
			return true
		}
		rd.rawData.Store(confKey, rawValue)

		switch kConfig.DataType {
		case "int":
			str, ok := rawValue.(string)
			var intValue int
			if ok {
				intValue, err = strconv.Atoi(str)
			} else {
				str := rawValue.(json.Number).String()
				intValue, err = strconv.Atoi(str)
				if err != nil {
					// log.Printf("%#+v, %#+v", rawValue, kConfig)
					// panic(err)
					return true
				}
			}
			rd.intType.Store(confKey, intValue)
			break
		case "string":
			strValue, ok := rawValue.(string)
			if ok {
				rd.stringType.Store(confKey, strValue)
			}
		}
		return true
	})

	return
}

// handleFuncData 按按配置中的函数生成新字段
func (rd *RiskData) HandleFuncData(keyArr []string) error {
	if len(keyArr) == 0 {
		return nil
	}

	for _, v := range keyArr {
		conf, ok := rd.GetConf(v)
		log.Printf("k: %s, ok: %v, v: %#+v \n", v, ok, conf)
		if !ok {
			continue
		}
		err := rd.handleDataByConf(conf)
		if err != nil {
			// utils.dd(err)
		}
	}
	return nil
}

// 将数据按配置内容进行处理
func (rd *RiskData) handleDataByConf(conf ConfigJsonTpl) error {
	fn, err := GetFuncByName(conf.FuncName)
	if err != nil {
		return err
	}

	res, _, err := fn(rd, conf.Args...)
	if err != nil {
		return err
	}

	rd.rawData.Store(conf.KeyName, res)
	tp := reflect.TypeOf(res)
	switch tp.Kind() {
	case reflect.Int:
		rd.intType.Store(conf.KeyName, res)
	case reflect.String:
		rd.stringType.Store(conf.KeyName, res)
	}

	return nil
}

// Get 获取原始数据以及类型
func (rd *RiskData) Get(key string) (value interface{}, valueType string, err error) {
	value, ok := rd.rawData.Load(key)
	if !ok {
		err = errors.New("获取不到数据")
		return
	}
	rawConfig, ok := rd.GetConf(key)
	if !ok {
		valueType = "string"
	}
	valueType = rawConfig.DataType
	return
}

func (rd *RiskData) GetString(key string) (value string, err error) {
	v, ok := rd.stringType.Load(key)
	if !ok {
		err = errors.New("获取不到 string 类型数据")
		return
	}
	value = v.(string)
	return
}

func (rd *RiskData) GetInt(key string) (value int, err error) {
	v, ok := rd.intType.Load(key)
	if !ok {
		err = errors.New("获取不到 int 类型数据")
		return
	}
	value = v.(int)
	return
}

func (rd *RiskData) GetKeyType(key string) (valueType string, err error) {
	cf, ok := rd.config.Load(key)
	if !ok {
		err = errors.New("获取不到类型配置")
		return
	}
	valueType = cf.(ConfigJsonTpl).DataType
	return
}
