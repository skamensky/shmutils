package calc

import (
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/manifoldco/promptui"
	"strings"
)

// TODO, allow for defining variables in terminal before running evaluation
func Calculate(expression string, interactiveMode bool) (interface{}, error) {
	if !interactiveMode {
		return gval.Evaluate(expression, nil)
	} else {
		variables := make(map[interface{}]interface{})
		for {
			p := promptui.Prompt{
				Label: "Enter expression. Define variables before the expression using the format 'varName=value'",
			}
			res, err := p.Run()
			if err != nil {
				fmt.Println("Something weird happened: %v", err)
			}

			if strings.Contains(res, "=") {
				variables[strings.Split(res, "=")[0]] = strings.Split(res, "=")[1]
				continue
			} else {
				return gval.Evaluate(res, variables)
			}
		}
	}

}
