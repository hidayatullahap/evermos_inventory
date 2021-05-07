package pkg

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// MarshalEnv returns the env encoding of v.
// It assign value in ENV_VARIABLE to v, support common type as string, int(8,16,32,64), uint(8,16,32,64)
// and float(32,64). Assign value requires ENV_VARIABLE value type is same with assigned struct field or it will skipped
// field struct with pointer type also will skipped
//
// Examples of struct field tags and their meanings:
//
//   // Field appears in ENV as key "db.name".
//   Field string `env:"db.name"`
//
//   // Name appears in ENV as key "my.name" it required and have default value 10.
//   Name string `env:"my.name,required,default=10"`
//
// Available tag options
//
// - required
//
// - default; ex: default=8000
func MarshalEnv(v interface{}) error {
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()

		typeOf := reflect.TypeOf(v)
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()

			var fieldNames []reflect.StructField

			for i := 0; i < typeOf.NumField(); i++ {
				fieldNames = append(fieldNames, typeOf.Field(i))
			}

			for _, field := range fieldNames {
				f := valueOf.FieldByName(field.Name)

				tag := field.Tag.Get("env")
				confValue, err := getConfigValue(tag)
				if err != nil {
					return err
				}

				if confValue != "" && f.CanSet() {
					convertType(f, confValue)
				}
			}
		}
	}

	return nil
}

func convertType(f reflect.Value, value string) {
	switch f.Kind() {
	case reflect.String:
		f.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			f.SetInt(v)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			f.SetUint(uint64(v))
		}
	case reflect.Float32:
		v, err := strconv.ParseFloat(value, 32)
		if err == nil {
			f.SetFloat(v)
		}
	case reflect.Float64:
		v, err := strconv.ParseFloat(value, 64)
		if err == nil {
			f.SetFloat(v)
		}
	}
}

func getConfigValue(key string) (string, error) {
	value := ""

	opt := splitTagOpt(key)
	if opt["key"] != "" {
		configKey := opt["key"]
		value = os.Getenv(configKey)

		if value == "" && opt["default"] != "" {
			value = opt["default"]
		}

		if value == "" && opt["required"] == "true" {
			return value, errors.New(fmt.Sprintf("env config key %s not set", configKey))
		}
	}

	return value, nil
}

func splitTagOpt(key string) map[string]string {
	opt := make(map[string]string)
	opt["key"] = ""
	opt["required"] = "false"
	opt["default"] = ""

	s := strings.Split(key, ",")
	sLen := len(s)

	if sLen > 0 {
		for i, key := range s {
			if i == 0 {
				opt["key"] = key
			}

			if key == "required" {
				opt["required"] = "true"
			}

			if strings.Contains(key, "default") {
				def := strings.Split(key, "=")
				if len(def) > 0 {
					opt["default"] = def[1]
				}
			}
		}
	}

	return opt
}
