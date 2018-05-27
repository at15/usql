package cassandra

import (
	"github.com/xo/usql/drivers"
)

// https://github.com/gocql/gocql

func init() {
	drivers.Register("cassandra", drivers.Driver{
		AllowMultilineComments: true,
		AllowCComments:         true,
	})
}