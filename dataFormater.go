package risk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Text() {
	fmt.Println("hello")
}

type RiskData struct {
	rawData    rawType
	config     map[string]ConfigJsonTpl
	intType    map[string]int
	stringType map[string]string
}

// 原始数据类型
type rawType map[string]interface{}
type IntType int
type StringType string

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

	rd.config = make(map[string]ConfigJsonTpl, len(cf))
	for i := 0; i < len(cf); i++ {
		rd.config[cf[i].KeyName] = cf[i]
	}
	return nil
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

	rd.rawData = make(rawType, len(tpl))
	rd.rawData = tpl

	log.Printf("%#+v \n", tpl)
	err = rd.swithcType()
	if err != nil {
		return err
	}

	// log.Println(rd)

	return nil
}

func dd(params ...interface{}) {
	log.Println(params...)
	os.Exit(0)
}

func (rd *RiskData) swithcType() (err error) {
	if rd.config == nil {
		return errors.New("未加载字段配置")
	}

	rd.intType = make(map[string]int, 1)
	rd.stringType = make(map[string]string, 1)

LOOP:
	for k, kConfig := range rd.config {
		rawValue, exist := rd.rawData[k]

		if !exist {
			continue LOOP
		}

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
					continue LOOP
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
