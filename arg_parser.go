package main

import (
	"bytes"
	"errors"
	"os"
	"reflect"
	"strings"
)

type argParser struct {
	args map[string]string
}

func newArgParser() *argParser {
	return &argParser{args: make(map[string]string)}
}

func (a *argParser) parse(args []string) {
	buf := bytes.Buffer{}
	var lastKey string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if lastKey != "" {
				buf.Truncate(buf.Len() - 1)
				a.args[lastKey] = buf.String()
				buf.Reset()
				lastKey = ""
			}
			if strings.HasPrefix(arg, "--") {
				lastKey = arg[2:]
			} else {
				lastKey = arg[1:]
			}
		} else {
			buf.WriteRune(' ')
			buf.WriteString(arg)
		}
	}
}

func (a *argParser) getArg(name, alias string) string {
	return ""
}

func ParseArgs(obj interface{}) error {
	if obj == nil {
		return errors.New("argument should be non-nil")
	}

	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	if objType.Kind() != reflect.Ptr {
		// the argument must be a pointer referring to the raw struct object
		return errors.New("argument must be pointer type")
	}
	if objValue.IsNil() {
		// this means the pointer's value is nil
		return errors.New("argument cannot be null-pointer")
	}

	objType = objType.Elem()   // get the type of *ptr
	objValue = objValue.Elem() // get the value of *ptr

	parser := &argParser{}
	parser.parse(os.Args)

	var field reflect.StructField
	var argValue string
	for i := 0; i < objType.NumField(); i++ {
		field = objType.Field(i)
		argName := field.Tag.Get("arg")
		argAlias := field.Tag.Get("alias")
		argValue = parser.getArg(argName, argAlias)
		if argValue != "" {
			// handle type convert
			objValue.Field(i).Set(reflect.ValueOf(argValue))
		}
		// else: ignore this field
	}
	return nil
}
