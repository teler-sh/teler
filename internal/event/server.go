package event

import (
	"github.com/kitabisa/teler/common"
	"github.com/r3labs/sse/v2"
)

type server struct {
	server  *sse.Server
	version string
	options *common.Options
}
