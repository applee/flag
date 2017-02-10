package flag

import (
	"fmt"
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

func (f *FlagSetEx) loadTOMLTree(tree *toml.TomlTree, path []string) error {
	for _, key := range tree.Keys() {
		fullPath := append(path, key)
		value := tree.Get(key)
		if subtree, isTree := value.(*toml.TomlTree); isTree {
			err := f.loadTOMLTree(subtree, fullPath)
			if err != nil {
				return err
			}
		} else {
			fullPath := strings.Join(append(path, key), ".")
			if f.Lookup(fullPath) != nil { // 允许toml存在冗余选项
				err := f.Set(fullPath, fmt.Sprintf("%v", value))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
