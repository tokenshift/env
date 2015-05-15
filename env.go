package env

import "fmt"
import "os"
import "strconv"
import "strings"

// Parses an environment variable of the form "KEY=VALUE".
func parseEntry(e string) (key, value string, err error) {
	sep := strings.Index(e, "=")

	if sep < 1 {
		return "", "", fmt.Errorf("Malformed environment variable: %s", e)
	}

	return e[0:sep], e[sep+1:], nil
}

// Returns a map of all environment variables.
func All() (map[string] string) {
	m := make(map[string] string)

	for _, entry := range os.Environ() {
		k, v, err := parseEntry(entry)
		if err == nil {
			m[k] = v
		}
	}

	return m
}

// Gets the specified environment variable as a string.
func Get(key string) (val string, ok bool) {
	for _, entry := range os.Environ() {
		k, v, err := parseEntry(entry)
		if err == nil && k == key {
			return v, true
		}
	}

	return "", false
}

// Gets the specified environment variable as a double.
func GetFloat(key string) (val float64, ok bool) {
	s, ok := Get(key)
	if !ok {
		return
	}

	val, err := strconv.ParseFloat(s, 64)
	ok = err == nil
	return
}

// Gets the specified environment variable as an integer.
func GetInt(key string) (val int, ok bool) {
	s, ok := Get(key)
	if !ok {
		return
	}

	i, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		ok = false
		return
	}

	return int(i), true
}

// Gets the environment variable as a list (colon-separated strings).
func GetList(key string) (vals []string, ok bool) {
	s, ok := Get(key)
	if !ok {
		return
	}

	vals = make([]string, 0)

	i0 := 0
	for i1 := 0; i1 < len(s); i1 += 1 {
		if s[i1] == ':' && (i1 <= i0 || s[i1-1] != '\\') {
			// Unescaped ':'; process the delimited value.
			vals = append(vals, strings.Replace(s[i0:i1], "\\:", ":", -1))
			i0 = i1 + 1
		}
	}

	return
}

func MustGet(key string) (val string) {
	val, ok := Get(key)
	if !ok {
		panic(fmt.Errorf("Missing environment variable $%s.", key))
	}
	return
}

func MustGetFloat(key string) (val float64) {
	val, ok := GetFloat(key)
	if !ok {
		panic(fmt.Errorf("Missing environment variable $%s (float).", key))
	}
	return
}

func MustGetInt(key string) (val int) {
	val, ok := GetInt(key)
	if !ok {
		panic(fmt.Errorf("Missing environment variable $%s (int).", key))
	}
	return
}

func MustGetList(key string) (vals []string) {
	vals, ok := GetList(key)
	if !ok {
		panic(fmt.Errorf("Missing environment variable $%s (list).", key))
	}
	return
}
