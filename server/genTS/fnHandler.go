package genTypescript

import (
	"fmt"
	"reflect"
	"strings"
)

func genTSFuncFromQuery(stringBuilder *strings.Builder, queryParams, output interface{}, address string) {

	stringBuilder.WriteString("(")
	if queryParams != nil {
		qpType := reflect.TypeOf(queryParams)
		stringBuilder.WriteString("queryParams:")
		stringBuilder.WriteString(GoFieldsToTSObj(qpType))
	}
	stringBuilder.WriteString("):Promise<")

	if output != nil {
		outputType := reflect.TypeOf(output)
		stringBuilder.WriteString(GoFieldsToTSObj(outputType))
	} else {
		stringBuilder.WriteString("void")
	}
	stringBuilder.WriteString(">=>")
	generateQueryFnBody(stringBuilder, queryParams != nil, address)
}

func genTSFuncFromMutation(stringBuilder *strings.Builder, queryParams, input, output interface{}, address string) {

	stringBuilder.WriteString("(")

	isParams := queryParams != nil || input != nil
	if isParams {
		stringBuilder.WriteString("parameters : {")
	}

	if queryParams != nil {
		qpType := reflect.TypeOf(queryParams)
		if qpType.Kind() == reflect.Ptr {
			qpType = qpType.Elem()
		}
		stringBuilder.WriteString(fmt.Sprintf("queryParams:%s,", GoFieldsToTSObj(qpType)))
	}
	if input != nil {
		inputType := reflect.TypeOf(input)
		if inputType.Kind() == reflect.Ptr {
			inputType = inputType.Elem()
		}
		stringBuilder.WriteString(fmt.Sprintf("input:%s", GoFieldsToTSObj(inputType)))
	}

	if isParams {
		stringBuilder.WriteString("}")
	}

	stringBuilder.WriteString("):Promise<")
	if output != nil {
		outputType := reflect.TypeOf(output)
		if outputType.Kind() == reflect.Ptr {
			outputType = outputType.Elem()
		}
		stringBuilder.WriteString(GoFieldsToTSObj(outputType))
	} else {
		stringBuilder.WriteString("void")
	}
	stringBuilder.WriteString(">=>")
	generateMutationFnBody(stringBuilder, isParams, address)
}
func generateQueryFnBody(stringBuilder *strings.Builder, isQuery bool, address string) {
	stringBuilder.WriteString("{return rpcCall(")
	if isQuery {
		stringBuilder.WriteString("{queryParams}")
	} else {
		stringBuilder.WriteString("undefined")
	}
	stringBuilder.WriteString(",")
	stringBuilder.WriteString("'" + address + `'`)

	stringBuilder.WriteString(")}")

}
func generateMutationFnBody(stringBuilder *strings.Builder, isParams bool, address string) {
	stringBuilder.WriteString("{return rpcCall(")
	if isParams {
		stringBuilder.WriteString("parameters")
	} else {
		stringBuilder.WriteString("undefined")
	}
	stringBuilder.WriteString(",")
	stringBuilder.WriteString("'" + address + `'`)

	stringBuilder.WriteString(")}")
}
