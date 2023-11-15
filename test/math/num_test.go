package math

import (
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/assert"

	"github.com/cookedsteak/gengine/builder"
	"github.com/cookedsteak/gengine/context"
	"github.com/cookedsteak/gengine/engine"
)

type Entity struct {
	Score  int32
	Height interface{}
	Name   string
	Comp   string
	Equal  bool
}

func Test_Num(t *testing.T) {
	var num_rule = `
	rule "rule name" "rule desc"
	begin
	entity.Score = 100 + 12/3
	entity.Height = 1.1234567890123456789
	end
	`
	entity := &Entity{Score: 0}

	dataContext := context.NewDataContext()
	dataContext.Add("entity", entity)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(num_rule)
	end1 := time.Now().UnixNano()

	println(fmt.Sprintf("rules num:%d, load rules cost time:%d ns", len(ruleBuilder.Kc.RuleEntities), end1-start1))

	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()

	start := time.Now().UnixNano()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
	err = eng.Execute(ruleBuilder, true)
	end := time.Now().UnixNano()
	println(entity.Score)
	println(entity.Height.(decimal.Decimal).String())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int32(104), entity.Score, "should equal")
	assert.Equal(t, "1.1234567890123456789", entity.Height.(decimal.Decimal).String(), "should equal")
	println(fmt.Sprintf("execute rule cost %d ns", end-start))

}

func Test_DivZero(t *testing.T) {
	zeroRule := `
	rule "div zero" "rule desc"
	begin
	entity.Height = 10/entity.Score
	end`

	entity := &Entity{Score: 0}
	dataContext := context.NewDataContext()
	dataContext.Add("entity", entity)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(zeroRule)
	end1 := time.Now().UnixNano()

	println(fmt.Sprintf("rules num:%d, load rules cost time:%d ns", len(ruleBuilder.Kc.RuleEntities), end1-start1))

	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()

	start := time.Now().UnixNano()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
	err = eng.Execute(ruleBuilder, true)
	end := time.Now().UnixNano()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, decimal.NewFromFloat(100.0).String(), entity.Height.(decimal.Decimal).String(), "return 100 not error")
	println(fmt.Sprintf("execute rule cost %d ns", end-start))
}

func Test_Decimal(t *testing.T) {
	decimalRule := `
	rule "decimal rule" "rule desc"
	begin
	x = a + 1.1
	b = "gogo" + "gaga"
	c = 1.1234567890123456789
	entity.Height = a + 1.1
	entity.Name = b
	entity.Comp = "one"
	if (c<2) {
		entity.Comp = "two"
	}
	entity.Equal = (d=="3306")
	return (x==2.2234567890123456789) 
	end
	`
	entity := &Entity{Score: 0}
	dataContext := context.NewDataContext()
	dataContext.Add("entity", entity)
	dataContext.Add("println", fmt.Println)
	a, _ := decimal.NewFromString("1.12345678901234567890")
	dataContext.Add("a", a)
	dataContext.Add("d", "3306")

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(decimalRule)
	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()
	err = eng.Execute(ruleBuilder, true)
	resMap, err := eng.GetRulesResultMap()
	println(entity.Height.(decimal.Decimal).String())
	println(entity.Name)
	println(entity.Comp)
	println(entity.Equal)
	fmt.Printf("%v", resMap)
	assert.Equal(t, "gogogaga", entity.Name, "name结果")
	assert.Equal(t, "two", entity.Comp, "comp结果")
	assert.Equal(t, true, resMap["decimal rule"], "结果为true")
}
