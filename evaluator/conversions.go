package evaluator

import (
	"errors"
	"plusLang/objects"
	"strconv"
	"strings"
)

func toNum(arguments []objects.Object) ([]objects.Object, error) {
	output := make([]objects.Object, 0)
	for i := 0; i < len(arguments); i++ {
		switch v := arguments[i].(type) {
		case objects.Function:
			return arguments, errors.New("Failed conversion: unable to convert function to number array")
		case objects.NumberArray:
			output = append(output, v)
			break
		case objects.String:
			stringsToConvert := strings.Fields(v.ToString())
			outputFloats := make([]float64, 0)

			for i := 0; i < len(stringsToConvert); i++ {
				curInt, err := strconv.ParseFloat(stringsToConvert[i], 64)
				if err != nil {
					return arguments, errors.New("Unable to parse number in conversion number to string")
				}
				outputFloats = append(outputFloats, curInt)
			}

			output = append(output, objects.NewNumberArray(outputFloats))
			break
		}
	}

	return output, nil
}

func toString(arguments []objects.Object) ([]objects.Object, error) {
	output := make([]objects.Object, 0)
	for i := 0; i < len(arguments); i++ {
		switch v := arguments[i].(type) {
		case objects.Function:
			return arguments, errors.New("Failed conversion: unable to convert function to number array")
		case objects.NumberArray:
			curString := v.ToString()
			curString = string([]rune(curString)[:len([]rune(curString))-1])
			output = append(output, objects.NewString(curString))
			break
		case objects.String:
			output = append(output, v)
			break
		}
	}

	return output, nil
}
