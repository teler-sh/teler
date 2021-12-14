package logs

import (
	"encoding/json"
	"fmt"

	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/errors"
)

// File write detected threats into it
func File(options *common.Options, log map[string]string) {
	var out string
	file := options.Configs.Logs.File

	if options.Output != nil {
		if file.JSON {
			logJSON, err := json.Marshal(log)
			if err != nil {
				errors.Exit(err.Error())
			}
			out = fmt.Sprintf("%s\n", logJSON)
		} else {
			out = fmt.Sprintf("[%s] [%s] [%s] %s\n",
				log["time_local"],
				log["remote_addr"],
				log["category"],
				log[log["element"]],
			)
		}

		if _, write := options.Output.WriteString(out); write != nil {
			errors.Show(write.Error())
		}
	}
}
