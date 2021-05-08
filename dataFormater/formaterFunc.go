package dataFormater

/**
这里是所有的转换函数, 每个函数接参的格式一样, 由函数内部将收到的参数类型进行转换
如果不符合函数要求,直接返回错误, 由调用方进行处理

所有支持的函数在 init() 中通过 registeFunc 进行注册,
使用方通过 GetFuncByName 将函数名换成函数实体
*/

import (
	"bytes"
	"errors"
	"log"
	"risk/utils"
	"sync"
	"time"
)

// 本地允许的函数集合
var funcMap sync.Map

// DataFunc 数据处理允许的参数格式
type DataFunc func(rd *RiskData, args ...interface{}) (res interface{}, tp string, err error)

func init() {
	registeFunc("dateFormat", FuncTimerFormater)
	registeFunc("combine", FuncCombine)
	registeFunc("subStr", FuncSubStr)
}

func registeFunc(name string, fn DataFunc) {
	funcMap.Store(name, fn)
}

// GetFuncByName 函数名换函数实体
func GetFuncByName(n string) (fn DataFunc, err error) {
	d, ok := funcMap.Load(n)
	if !ok {
		err = errors.New("找不到该名称对应的函数: " + n)
		return
	}
	fn = d.(DataFunc)
	return
}

// FuncTimerFormater 时间格式化
func FuncTimerFormater(rd *RiskData, args ...interface{}) (res interface{}, tp string, err error) {
	log.Println("FuncTimerFormater")
	defer func() {
		exp := recover()
		if exp != nil {
			err = exp.(error)
		}
	}()
	col := args[0].(string)
	format := args[1].(string)
	value, err := rd.GetString(col)
	if err != nil {
		return
	}

	tp = "string"

	timer, _ := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	res = timer.Format(format)
	return
}

// FuncCombine  合并字段
func FuncCombine(rd *RiskData, args ...interface{}) (res interface{}, tp string, err error) {
	log.Println("FuncCombine")
	defer func() {
		exp := recover()
		if exp != nil {
			err = exp.(error)
		}
	}()

	var buffer bytes.Buffer
	var bt []byte
	var v interface{}
	for i := 0; i < len(args); i++ {
		v, _, err = rd.Get(args[i].(string))
		// log.Printf("k: %v, v: %v, e: %v \n", args[i], v, err)
		if err != nil {
			return
		}

		bt, err = utils.GetBytes(v)
		if err != nil {
			return
		}

		buffer.Write(bt)
		if i < len(args)-1 {
			buffer.WriteString(",")
		}
	}

	str := buffer.String()
	// 移除制表符
	str, _ = utils.PureString(str)

	res = str
	tp = "string"

	return
}

// 字符串截取
func FuncSubStr(rd *RiskData, args ...interface{}) (res interface{}, tp string, err error) {
	log.Println("FuncSubStr")
	defer func() {
		exp := recover()
		if exp != nil {
			err = exp.(error)
		}
	}()
	return
}
