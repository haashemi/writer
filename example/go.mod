module github.com/haashemi/writer/example

go 1.23.0

toolchain go1.24.1

replace github.com/haashemi/writer v1.0.0 => ../

require (
	github.com/haashemi/painter v0.1.1-0.20240129225408-5dd5dda5e35f
	github.com/haashemi/writer v1.0.0
)

require (
	github.com/mattn/go-pointer v0.0.1 // indirect
	golang.org/x/image v0.25.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
