package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var realVars map[string]string
var fakeVars map[string]string

func init() {
	realVars = make(map[string]string)
	fakeVars = make(map[string]string)

	for _, value := range(os.Environ()) {
		sep := strings.Index(value, "=")
		realVars[value[0:sep]] = value[sep+1:]
	}
}

func Get(name string) (string, bool) {
	if val, ok := fakeVars[name]; ok {
		return val, ok
	}

	if val, ok := realVars[name]; ok {
		return val, ok
	}

	return "", false
}

func Set(name, val string) {
	fakeVars[name] = val
}

func Unset(name string) {
	delete(fakeVars, name)
}

func GetInt(name string) (int, bool, error) {
	if val, ok := Get(name); ok {
			i, err := strconv.ParseInt(val, 0, 0)
			return int(i), true, err
		} else {
			return 0, false, nil
		}
}

func GetIntDefault(name string, dflt int) (int, bool, error) {
	if val, ok, err := GetInt(name); err != nil {
		return dflt, ok, err
	} else if ok {
		return val, ok, err
	} else {
		return dflt, ok, err
	}
}

func MustGet(name string) string {
	if val, ok := Get(name); ok {
		return val
	} else {
		panic(fmt.Errorf("$%s is required.", name))
	}
}

func MustGetInt(name string) int {
	if val, ok, err := GetInt(name); !ok {
		panic(fmt.Errorf("$%s is required.", name))
	} else if err != nil {
		panic(fmt.Errorf("$%s must be an integer.", name))
	} else {
		return val
	}
}

func MustGetIntDefault(name string, dflt int) int {
	if val, _, err := GetIntDefault(name, dflt); err != nil {
		panic(err)
	} else {
		return val
	}
}
