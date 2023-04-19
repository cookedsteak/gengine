package test

import (
	"fmt"
	"testing"

	"github.com/cookedsteak/gengine/builder"
	"github.com/cookedsteak/gengine/context"
	"github.com/cookedsteak/gengine/engine"
)

func IsIn(p1 string, pa ...string) bool {
	for _, v := range pa {
		if p1 == v {
			return true
		}
	}
	return false
}

func Test_in(t *testing.T) {

	dc := context.NewDataContext()
	dc.Add("isIn", IsIn)
	dc.Add("println", fmt.Println)
	rb := builder.NewRuleBuilder(dc)
	e := rb.BuildRuleFromString(`
rule "isIn"
begin
	if isIn("hello", "hello", "world", "x1", "x2") {
		println("hello")
	}
end`)

	if e != nil {
		panic(e)
	}

	gengine := engine.NewGengine()
	e = gengine.Execute(rb, true)
	if e != nil {
		panic(e)
	}

}
