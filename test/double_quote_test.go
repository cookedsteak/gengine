package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cookedsteak/gengine/engine"
)

func TestDoubleQuote(t *testing.T) {
	rule := `
	rule "1"  salience 1
	begin
		pp(s1,s2)
		s3 = "{\"code\"}"
		return s1==s3 && s2==s3
	end
`
	ep, err := engine.NewGenginePool(1, 2, 2, string(rule), nil)
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["pp"] = fmt.Println
	data["s1"] = `{"code"}`
	data["s2"] = "{\"code\"}"
	err, res := ep.Execute(data, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, res["1"])
}
