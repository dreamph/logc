module github.com/dreamph/logc/zap

go 1.23

replace github.com/dreamph/logc => ../

require (
	github.com/dreamph/logc v1.0.1
	go.uber.org/zap v1.27.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require go.uber.org/multierr v1.11.0 // indirect
