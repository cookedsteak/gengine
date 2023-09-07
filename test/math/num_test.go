package math

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cookedsteak/gengine/builder"
	"github.com/cookedsteak/gengine/context"
	"github.com/cookedsteak/gengine/engine"
)

type Entity struct {
	Score  int32
	Height float64
}

const num_rule = `

rule "rule name" "rule desc"
begin
entity.Score = 100 + 12/3
entity.Height = 1.12345678901234567890
end
`

func exec_num() {

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
	println(entity.Height)
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("execute rule cost %d ns", end-start))

}

func Test_num(t *testing.T) {
	exec_num()
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
	assert.Equal(t, 100.0, entity.Height, "return 100 not error")
	println(fmt.Sprintf("execute rule cost %d ns", end-start))
}

func Test_Decimal(t *testing.T) {
	decimalRule := `
	rule "decimal rule" "rule desc"
	begin
	a = 1.12345678901234567890+1.1
	entity.Height = a
	return a
	end
	`
	entity := &Entity{Score: 0}
	dataContext := context.NewDataContext()
	dataContext.Add("entity", entity)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(decimalRule)
	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()
	err = eng.Execute(ruleBuilder, true)
	println(entity.Height)
}
