# logrusltsv
LTSV logger for [logrus](https://github.com/Sirupsen/logrus).

In this package, keys which include invalid character are ignored and not logged.
See [LTSV documentation](http://ltsv.org/).

```go
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hitsumabushi/logrusltsv"
)

func main() {
	log.SetFormatter(&logrusltsv.LtsvFormatter{})

	log.WithFields(log.Fields{
		"animal":    "walrus",
		"size":      10,
		"colon:key": "ignored",
	}).Info("A walrus appears")
	// #=> animal:walrus	level:info	message:A walrus appears	size:10
}
```

