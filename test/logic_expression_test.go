package test

import (
	"fmt"
	"testing"

	"github.com/cookedsteak/gengine/builder"
	"github.com/cookedsteak/gengine/context"
	"github.com/cookedsteak/gengine/engine"
)

const logic_expr = `
rule "name test" "i can" salience 10
begin

res = a sin "1,2,3,4"
println(res)
res2 = c snin "findPassword_bottom_e500_email_sms,findPassword_bottom_e_email_sms,findPassword_bottom_p_email_sms,findPassword_bottom_p_sms,findPassword_bottom_p_sms_ga,findPassword_clusterRisk_e_email_face,findPassword_clusterRisk_e_email_sms,findPassword_clusterRisk_p_email_sms,findPassword_clusterRisk_p_sms,findPassword_clusterRisk_p_sms_ga,findPassword_secEnv_sms,findPassword_secEnv_email,findPassword_clusterRisk_e_email_face"
println(res2)
res3 = b sinc "what"
println(res3)
res4 = d sninc "what"
println(res4)

end
`

func Test_logic_expr(t *testing.T) {
	dataContext := context.NewDataContext()
	dataContext.Add("println", fmt.Println)
	dataContext.Add("a", "22")
	dataContext.Add("c", "findPassword_bottom_e500_email_sms")
	dataContext.Add("b", "what are you doing")
	dataContext.Add("d", "what are you doing")

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	e1 := ruleBuilder.BuildRuleFromString(logic_expr)
	if e1 != nil {
		panic(e1)
	}

	eng := engine.NewGengine()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
	e2 := eng.Execute(ruleBuilder, true)
	if e2 != nil {
		panic(e2)
	}

}
