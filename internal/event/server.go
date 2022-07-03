package event

import (
	"github.com/r3labs/sse/v2"
	"teler.app/common"
)

type server struct {
	server  *sse.Server
	version string
	options *common.Options
}
