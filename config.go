package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"

	"./action"
	"./g13"
)

// Config input actions
type Config map[string]map[string]interface{}

// ParseTOMLConfig parses a string into a Config map
func ParseTOMLConfig(s *string) (*Config, error) {
	var c Config
	_, err := toml.Decode(*s, &c)
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
					invalid = append(invalid, fmt.Errorf("unknown key \"%s\"", key))
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
						if _, ok := action.KEY[actionKey]; !ok {
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
func (c *Config) ToActions(h *action.Handler, actions action.Profiles) {
	for profile, acts := range *c {
		for keys, act := range acts {
			if _, ok := actions[profile]; !ok {
				actions[profile] = action.Actions{}
			}

			switch act.(type) {
			case string:
				action := act.(string)
				actions[profile][g13.ParseKey(keys)] = h.BindLua(&action)
			case []interface{}:
				actionKeys := act.([]interface{})[0].(string)
				actions[profile][g13.ParseKey(keys)] = h.KeyPressAction(actionKeys)
			}
		}
	}
}
