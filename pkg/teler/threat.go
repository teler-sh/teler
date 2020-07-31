package teler

import (
	"fmt"

	"github.com/kitabisa/teler/common"
	"github.com/satyrius/gonx"
)

// Analyze logs from threat resources
func Analyze(options *common.Options, log *gonx.Entry) {
	fmt.Printf("%+v, %+v", options, log)
}
