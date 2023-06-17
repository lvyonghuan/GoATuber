package function

//函数的存储工具

import (
	"GoTuber/NLP/service/gpt"
	"reflect"
	"runtime"
)

// Function 存储开发者所编写的函数，以便通过字符调用
var Function = make(map[string]func([]string) string)

// GetFunctionResult 根据字符串调用函数，返回字符串
func GetFunctionResult(functionName string, parameter []string) string {
	return executeFunction(get(functionName), parameter)
}

// 根据函数名称获取函数
func get(funcName string) func([]string) string {
	return Function[funcName]
}

// 执行函数
func executeFunction(fun func([]string) string, parameter []string) string {
	return fun(parameter)
}

// 添加函数
func addFunc(fun func([]string) string) {
	funcName := getFunctionName(fun)
	Function[funcName] = fun
}

// 添加函数的json信息
func addFuncJson(fun interface{}) {
	gpt.FunctionJson = append(gpt.FunctionJson, fun)
}

// 获取函数
func getFunctionName(fn func([]string) string) string {
	pc := reflect.ValueOf(fn).Pointer()
	function := runtime.FuncForPC(pc)
	return function.Name()
}
