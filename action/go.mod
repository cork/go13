module cork/go13/action

go 1.18

replace cork/go13/g13 => ../g13

require (
	cork/go13/g13 v0.0.0-00010101000000-000000000000
	github.com/BurntSushi/toml v1.0.0
	github.com/sashko/go-uinput v0.0.0-20200718185411-c753d6644126
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9
	layeh.com/gopher-luar v1.0.10
)

require github.com/google/gousb v1.1.1 // indirect
