# logc

Golang Logger Wrapper
- Easy for change Log Library Support (zap, zerolog)
- Simple & Easy 

Install
=======
``` sh
go get github.com/dreamph/logc
```

Examples - Use Zap
=======
``` go
package main

import (
	"context"
	"github.com/dreamph/logc"
	"github.com/dreamph/logc/zap"
)

func main() {
	logger := zap.NewLogger(&logc.Options{
		FilePath: "./app.log",
		Level:    "debug",
		Format:   "json",
		MaxAge:   30,
		MaxSize:  10,
	})
	defer logger.Release()

	logger.Info("Test Info")
	logger.Warn("Test Warn")

	d := map[string]interface{}{
		"requestId": "123",
	}
	log := logger.WithLogger(logc.WithValue(context.Background(), d))
	log.Info("Test")

}
```

Examples - Use ZeroLog
=======
``` go
package main

import (
	"context"
	"github.com/dreamph/logc"
	"github.com/dreamph/logc/zerolog"
)

func main() {
	logger := zerolog.NewLogger(&logc.Options{
		FilePath: "./app.log",
		Level:    "debug",
		Format:   "json",
		MaxAge:   30,
		MaxSize:  10,
	})
	defer logger.Release()

	logger.Info("Test Info")
	logger.Warn("Test Warn")

	d := map[string]interface{}{
		"requestId": "123",
	}
	log := logger.WithLogger(logc.WithValue(context.Background(), d))
	log.Info("Test")
}
```


Buy Me a Coffee
=======
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dreamph)