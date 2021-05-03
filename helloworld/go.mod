module helloworld

go 1.16

require (
	github.com/myuser/calculator v0.0.0
	github.com/rs/zerolog v1.21.0 // indirect
)

replace github.com/myuser/calculator => ../calculator
