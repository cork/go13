package action

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"

	"cork/go13/g13"
)

// ConfigActions contains the toml parsed actions
type ConfigActions map[string]map[string]interface{}

// Config struct for input actions and storage
type Config struct {
	Actions ConfigActions
	Folder  string
}

// NewConfig creates a new Config struct
func NewConfig(configFolder string) (*Config, error) {
	if configFolder == "" {
		return nil, errors.New("TOML config folder missing")
	}

	if folder, err := os.Stat(configFolder); configFolder == "" || err != nil || !folder.IsDir() {
		return nil, errors.New("TOML config folder missing")
	}

	return &Config{Folder: configFolder}, nil
}

// LoadTOMLConfig loads name into the config
func (c *Config) LoadTOMLConfig(name string) error {
	name = path.Clean(name)

	if _, err := os.Stat(path.Join(c.Folder, name+".toml")); os.IsNotExist(err) {
		return err
	}

	var actions ConfigActions
	if _, err := toml.DecodeFile(path.Join(c.Folder, name+".toml"), &actions); err != nil {
		return err
	}

	c.Actions = actions

	return nil
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
func (actions *ConfigActions) Validate() ([]error, bool) {
	if len(*actions) < 1 {
		return []error{errors.New("no configurations given")}, true
	}

	var invalid []error

	for _, acts := range *actions {
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
	for profile, acts := range c.Actions {
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

				if actions[profile][g13Keys] != nil && actions[profile][g13Keys].state != nil {
					actions[profile][g13Keys].state.Close()
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

	if _, err := os.Stat(path.Join(c.Folder, name+".toml")); err == nil {
		return errors.New("file exists")
	}

	if err := os.MkdirAll(c.Folder, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(path.Join(c.Folder, name+".toml"))
	if err != nil {
		return err
	}
	defer f.Close()

	if err := toml.NewEncoder(f).Encode(c.Actions); err != nil {
		return err
	}

	f.Sync()

	return nil
}
