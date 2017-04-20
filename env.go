package flag

import (
	"flag"
	"fmt"
	"strings"
)

// ParseEnv parses flags from environment variables.
// Flags already set will be ignored.
func (f *FlagSetEx) ParseEnv(environ []string) error {
	env := make(map[string]string)
	for _, s := range environ {
		i := strings.Index(s, "=")
		if i < 1 {
			continue
		}
		env[s[0:i]] = s[i+1 : len(s)]
	}

	var err MultiError
	f.VisitAll(func(fi *flag.Flag) {
		envKey := strings.ToUpper(fi.Name)
		if f.envPrefix != "" {
			envKey = f.envPrefix + "_" + envKey
		}
		r := strings.NewReplacer("-", "_", ".", "_")
		envKey = r.Replace(envKey)

		value, isSet := env[envKey]
		if !isSet || len(value) <= 0 {
			return
		}

		//TODO: bool doesn't need value
		if e := fi.Value.Set(value); e != nil {
			err = append(err, fmt.Errorf("invalid value %q for environment variable %s: %v", value, envKey, err))
		}
	})

	if len(err) == 0 {
		return nil
	}
	return err
}
