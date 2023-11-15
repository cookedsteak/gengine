package test

import (
	"fmt"
	"testing"

	"github.com/cookedsteak/gengine/builder"
	"github.com/cookedsteak/gengine/context"
	"github.com/cookedsteak/gengine/engine"
)

func Test_input_array(t *testing.T) {

	var rule = `
rule "test" "test" 
begin
	InputAndResult.GetSlice("12","23")
	InputAndResult.AddStringArray("arr",paramArray,11,12,13)
	//if isNil(InputAndResult.Result["tenant_id"]) {
    //   InputAndResult.AddStringArray("data",InputAndResult.GetSlice("123","234"))
    //}
end
`

	dataContext := context.NewDataContext()
	//init rule engine
	InputAndResult := &InputAndResult{
		Result:  make(map[string][]string),
		Options: make(map[string]interface{}),
	}
	dataContext.Add("InputAndResult", InputAndResult)
	dataContext.Add("paramArray", []string{"a", "b", "c"})

	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	e1 := ruleBuilder.BuildRuleFromString(rule)
	if e1 != nil {
		panic(e1)
	}

	eng := engine.NewGengine()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
	e2 := eng.Execute(ruleBuilder, true)
	if e2 != nil {
		panic(e2)
	}

	//输出在规则中输入的内容
	println(fmt.Sprintf("%+v", InputAndResult.Result["data"]))
	println(fmt.Sprintf("%+v", InputAndResult.Result["arr"]))

}

type InputAndResult struct {
	Result  map[string][]string
	Options map[string]interface{}
}

func (input *InputAndResult) AddStringArray(key string, value []string, opt ...int) {
	input.Result[key] = value
	input.Options["opt"] = opt
}

func (input *InputAndResult) GetSlice(vs ...string) []string {
	var s []string
	for _, v := range vs {
		s = append(s, v)
	}
	return s
}
