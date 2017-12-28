// Package cfg implements functions to load tagged structs from JSON files.
//
// The structure's values can be arbitrary types that are supported in this
// library.  Currently supported types include: `float64`, `int` and `string`.
//
// Example structure:
//
// The only expected tag arguments are as follows (`cfg:"tag0,tag1"`):
//  tag0 -      The first tag item is the key name to look for in the `path' file.
//  tag1 -      The second tag item is an optional command like "required".
//
package cfg

////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// parseTag takes a cfg tag and returns the Key we are looking for and a boolean to
// indicate if the key is a required value or not.
func parseTag(tag string) (string, bool) {
	key := ""
	req := false

	items := strings.Split(tag, ",")
	if len(items) >= 1 {
		key = items[0]
	}
	if len(items) >= 2 {
		req = (strings.ToLower(items[1]) == "required")
	}

	return key, req
}

////////////////////////////////////////////////////////////////////////////////

// Load looks for a valid config file at `path` and if found, attempts to load
// it into memory.  Then it searches the interface{} `v` for any `cfg` tags and
// populates `v` accordingly.
func Load(path string, v interface{}) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Grab some json from the now valid file.
	var m map[string]interface{}
	if err = json.Unmarshal(bs, &m); err != nil {
		return err
	}

	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		vf := val.Field(i)
		tf := val.Type().Field(i)
		tag := tf.Tag.Get("cfg")

		key, req := parseTag(tag)
		if len(key) == 0 && req {
			key = tf.Name
		}

		// Look for the key in the json map if its a valid tag.
		if len(key) > 0 {
			value, ok := m[key]
			if !ok && req {
				return fmt.Errorf("missing required field in config (%v)", key)
			}

			switch tf.Type.Name() {
			case "string":
				vf.SetString(value.(string))
			case "int":
				vf.SetInt(int64(value.(float64)))
			case "float64":
				vf.SetFloat(value.(float64))
			default:
				return fmt.Errorf("unhandled type (%v)", tf.Type.Name())
			}
		} else if req {
			return fmt.Errorf("invalid field (%v)", tf.Name)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
