module github.com/dreamph/logc/zerolog

go 1.23

replace github.com/dreamph/logc => ../

require (
	github.com/dreamph/logc v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.33.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
