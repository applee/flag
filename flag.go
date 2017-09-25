package flag

import (
	"errors"
	"flag"
	"io"
	"os"
	"path/filepath"
	"time"
)

type FlagSetEx struct {
	*flag.FlagSet
	envPrefix string
}

var DefaultConfigFlagName = "config"

var ex = &FlagSetEx{
	flag.NewFlagSet(os.Args[0], flag.ContinueOnError),
	"",
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool) *bool {
	return ex.FlagSet.Bool(name, value, "")
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	ex.FlagSet.BoolVar(p, name, value, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func Int(name string, value int) *int {
	return ex.FlagSet.Int(name, value, "")
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVar(p *int, name string, value int, usage string) {
	ex.FlagSet.IntVar(p, name, value, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64(name string, value int64) *int64 {
	return ex.FlagSet.Int64(name, value, "")
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64Var(p *int64, name string, value int64, usage string) {
	ex.FlagSet.Int64Var(p, name, value, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func Uint(name string, value uint) *uint {
	return ex.FlagSet.Uint(name, value, "")
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint  variable in which to store the value of the flag.
func UintVar(p *uint, name string, value uint, usage string) {
	ex.FlagSet.UintVar(p, name, value, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, value uint64, usage string) *uint64 {
	return ex.FlagSet.Uint64(name, value, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	ex.FlagSet.Uint64Var(p, name, value, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, value string, usage string) *string {
	return ex.FlagSet.String(name, value, "")
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	ex.FlagSet.StringVar(p, name, value, usage)
}

// Float64 defines a float64 config variable with a given name and default value.
func Float64(name string, value float64, usage string) *float64 {
	return ex.FlagSet.Float64(name, value, usage)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
func Float64Var(p *float64, name string, value float64, usage string) {
	ex.FlagSet.Float64Var(p, name, value, usage)
}

// Duration defines a time.Duration config variable with a given name and default value.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return ex.FlagSet.Duration(name, value, usage)
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	ex.FlagSet.DurationVar(p, name, value, usage)
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func Var(p flag.Value, name string, usage string) {
	ex.FlagSet.Var(p, name, usage)
}

// Parse parses the command-line, environment variables and config file flags
// into the global ConfigSet.
// This must be called after all config flags have been defined but before the
// flags are accessed by the program.
func Parse() error {
	err := ex.FlagSet.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	err = ex.ParseEnv(os.Environ())
	if err != nil {
		return err
	}

	cf := ex.FlagSet.Lookup(DefaultConfigFlagName)
	if cf == nil {
		return nil
	}
	path := cf.Value.String()
	if len(path) > 0 {
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			return errors.New("Invalid config file.")
		}
		ext := filepath.Ext(path)
		switch ext {
		case ".toml":
			return ex.ParseTOML(path)
		default:
			return errors.New("Unsupported config file.")
		}
	}
	return nil
}

func SetOutput(output io.Writer) {
	ex.FlagSet.SetOutput(output)
}

func Init(name string, errorHandling flag.ErrorHandling) {
	ex.FlagSet.Init(name, errorHandling)
}

func Set(name, value string) error {
	return ex.FlagSet.Set(name, value)
}
