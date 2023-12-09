package utils

import (
	"fmt"
	"reflect"
	"time"
)

func ExecuteAndLogTime(fn interface{}, args ...interface{}) (interface{}, error) {
	// Check if first param is a function
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("first parameter is not a function")
	}

	// Check if the number of arguments matches the function's input parameters
	if len(args) != fnValue.Type().NumIn() {
		return nil, fmt.Errorf("Number of arguments does not match function's input parameters")
	}

	inArgs := make([]reflect.Value, len(args))
	for i := range args {
		inArgs[i] = reflect.ValueOf(args[i])
	}

	// Call function with the provided arguments
	startTime := time.Now()
	result := fnValue.Call(inArgs)
	elapsedTime := time.Since(startTime).String()
	fmt.Printf("Execution time: %v\n", elapsedTime)

	if len(result) > 0 {
		return result[0].Interface(), nil
	}

	return nil, nil
}
