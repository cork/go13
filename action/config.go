package action

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/spf13/pflag"

	"../g13"
)

// Config input actions
type Config map[string]map[string]interface{}

var configFolder = pflag.StringP("toml-path", "p", "config", "Storage folder for configuration files, defaults to working directory")

// LoadTOMLConfig loads a TOML configuration from file if the toml-path is configured
func LoadTOMLConfig(name string) (*Config, error) {
	name = path.Clean(name)

	if *configFolder == "" {
		return nil, errors.New("TOML config folder missing")
	}

	if _, err := os.Stat(path.Join(*configFolder, name+".toml")); os.IsNotExist(err) {
		return nil, err
	}

	var c Config
	if _, err := toml.DecodeFile(path.Join(*configFolder, name+".toml"), &c); err != nil {
		return nil, err
	}
	return &c, nil
}

// ParseTOMLConfig parses a string into a Config map
func ParseTOMLConfig(s string) (*Config, error) {
	var c Config
	_, err := toml.Decode(s, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Validate config input before converting them to actions
func (c *Config) Validate() ([]error, bool) {
	if len(*c) < 1 {
		return []error{errors.New("no configurations given")}, true
	}

	var invalid []error

	for _, acts := range *c {
		for key, act := range acts {
			keys := strings.Split(key, "+")
			for _, key := range keys {
				if _, ok := g13.KEY[key]; !ok {
					if _, ok = EVENTS[key]; !ok {
						invalid = append(invalid, fmt.Errorf("unknown key \"%s\"", key))
					}
				}
			}

			switch act.(type) {
			case string:
			case []interface{}:
				actionKeys := act.([]interface{})[0]
				switch actionKeys.(type) {
				case string:
					keys := strings.Split(strings.TrimSpace(actionKeys.(string)), "+")
					for _, actionKey := range keys {
						if _, ok := KEY[actionKey]; !ok {
							invalid = append(invalid, fmt.Errorf("\"%s\" has unknown action key \"%s\"", key, actionKey))
						}
					}
				default:
					invalid = append(invalid, fmt.Errorf("unsupported keys type \"%s\"", reflect.TypeOf(actionKeys).Name()))
				}
			default:
				invalid = append(invalid, fmt.Errorf("unsupported action type \"%s\"", reflect.TypeOf(act).Name()))
			}
		}
	}
	return invalid, len(invalid) < 1
}

// ToActions converts config values to actions
func (c *Config) ToActions(h *Handler, actions Profiles) {
	for profile, acts := range *c {
		for keys, act := range acts {
			if _, ok := actions[profile]; !ok {
				actions[profile] = Actions{}
			}

			switch act.(type) {
			case string:
				g13Keys, ok := EVENTS[keys]
				if !ok {
					g13Keys = g13.ParseKey(keys)
				}

				action := act.(string)
				actions[profile][g13Keys] = h.BindLua(&action)
			case []interface{}:
				actionKeys := act.([]interface{})[0].(string)
				actions[profile][g13.ParseKey(keys)] = h.KeyPressAction(actionKeys)
			}
		}
	}
}

// SaveAsTOML saves the config to disk if a config folter is configured
func (c *Config) SaveAsTOML(name string) error {
	name = path.Clean(name)

	if _, err := os.Stat(path.Join(*configFolder, name+".toml")); err == nil {
		return errors.New("file exists")
	}

	if err := os.MkdirAll(*configFolder, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(path.Join(*configFolder, name+".toml"))
	if err != nil {
		return err
	}
	defer f.Close()

	if err := toml.NewEncoder(f).Encode(c); err != nil {
		return err
	}

	f.Sync()

	return nil
}
