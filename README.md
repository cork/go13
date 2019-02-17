# go13 a Logitech G13 action server written in GO

This is very much a WIP. The basics are implemented but there is a long way left until what I want as a minimum is implemented.

## Dependencies

### System

- Depends on system installation of libusb (for now might use raw access in the future)

### GO

You will need:
- github.com/BurntSushi/toml
- github.com/google/gousb
- github.com/ogier/pflag
- github.com/sashko/go-uinput
- github.com/yuin/gopher-lua
- layeh.com/gopher-luar

## Install

```bash
go get github.com/BurntSushi/toml
go get github.com/google/gousb
go get github.com/ogier/pflag
go get github.com/sashko/go-uinput
go get github.com/yuin/gopher-lua
go get layeh.com/gopher-luar
go build
```

## Running

```bash
./go13
```

It will read configuration from test.toml for now.

## TODO

- [x] G13 buttons to single keyboard key mapping.
- [x] G13 buttons to multiple keyboard keys mapping.
- [x] G13 buttons to lua scripting actions mapping.
- [x] Parsing profile and actions from TOML configuration file.
- [x] Stick mouse movement support.
- [x] TOML configuration folder.
- [x] Action to switch configuration.
- [ ] HTTP endpoint for adding/modifying actions and profiles.
- [ ] HTTP endpoint for removing actions and profiles.
- [ ] HTTP endpoint for switcing profile.
- [ ] HTTP endpoint for switching configuration.
- [ ] Web configuration interface.
- [ ] Grapical interactive Web configuration interface.
