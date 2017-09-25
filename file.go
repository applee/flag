package flag

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/pelletier/go-toml"
)

func (f *FlagSetEx) ParseTOML(path string) error {
	tree, err := toml.LoadFile(path)
	if err != nil {
		return fmt.Errorf("Parse toml file error: %s", err.Error())
	}
	return f.loadTOMLTree(tree, nil)
}

func (f *FlagSetEx) loadTOMLTree(tree *toml.Tree, path []string) error {
	var err error
	for _, key := range tree.Keys() {
		fullPath := append(path, key)
		value := tree.Get(key)
		if subtree, isTree := value.(*toml.Tree); isTree {
			err := f.loadTOMLTree(subtree, fullPath)
			if err != nil {
				return err
			}
		} else {
			fullPath := strings.Join(append(path, key), ".")
			if f.Lookup(fullPath) != nil {
				// if value is slice of string, int64, uint64, float64, bool,
				// time.Time, marshal it to json
				if reflect.TypeOf(value).Kind() == reflect.Slice {
					var b []byte
					b, err = json.Marshal(value)
					f.Set(fullPath, string(b))
				} else {
					err = f.Set(fullPath, fmt.Sprintf("%v", value))
				}
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
