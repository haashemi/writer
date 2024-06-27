module github.com/haashemi/writer/example

go 1.21

replace github.com/haashemi/writer v1.0.0 => ../

require (
	github.com/haashemi/painter v0.1.1-0.20240129225408-5dd5dda5e35f
	github.com/haashemi/writer v1.0.0
)

require (
	github.com/haashemi/go-harfbuzz v0.0.0-20240304202021-7d8c8e99547f // indirect
	github.com/mattn/go-pointer v0.0.1 // indirect
	golang.org/x/image v0.18.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)
