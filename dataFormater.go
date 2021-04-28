package risk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

func Text() {
	fmt.Println("hello")
}

type RiskData struct {
	rawData    rawType
	config     sync.Map
	intType    map[string]int
	stringType map[string]string
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
	err := json.Unmarshal([]byte(configJson), &cf)

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
	cf = t.(ConfigJsonTpl)
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

// LoadData 加载数据
func (rd *RiskData) LoadData(raw string) (err error) {
	tpl := rawType{}
	err = JsonDecode(raw, &tpl, true)
	if err != nil {
		return err
	}
	rd.rawData = make(rawType, 1)
	// rd.rawData = tpl

	// log.Printf("%#+v \n", tpl)
	err = rd.swithcType(tpl)
	if err != nil {
		return err
	}

	return nil
}

// 将传入的数据从 interface{} 转成指定类型
// 需要函数处理的新字段不在这里处理
func (rd *RiskData) swithcType(data map[string]interface{}) (err error) {
	if rd.config == nil {
		return errors.New("未加载字段配置")
	}

	rd.intType = make(map[string]int, 1)
	rd.stringType = make(map[string]string, 1)

LOOP:
	for k, kConfig := range rd.config {
		// 配置中带 func 的稍后处置
		if kConfig.FuncName != "" || len(kConfig.FuncName) > 0 {
			continue LOOP
		}

		// 通过配置判断外界传入数据是否合法, 否则给过滤掉
		rawValue, exist := data[k]
		if !exist {
			continue LOOP
		}
		rd.rawData[k] = rawValue

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
					return
				}
			}
			rd.intType[k] = intValue
			break
		case "string":
			strValue, ok := rawValue.(string)
			if ok {
				rd.stringType[k] = strValue
			}
			break
		default:
			continue LOOP
		}
	}

	return
}

func (rd *RiskData) Get(key string) (value interface{}, valueType string, err error) {
	value, ok := rd.rawData[key]
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
	value, ok := rd.stringType[key]
	if !ok {
		err = errors.New("获取不到 string 类型数据")
		return
	}
	return
}

func (rd *RiskData) GetInt(key string) (value int, err error) {
	value, ok := rd.intType[key]
	if !ok {
		err = errors.New("获取不到 int 类型数据")
		return
	}
	return
}

func (rd *RiskData) GetKeyType(key string) (valueType string, err error) {
	cf, ok := rd.config[key]
	if !ok {
		err = errors.New("获取不到类型配置")
		return
	}
	valueType = cf.DataType
	return
}
