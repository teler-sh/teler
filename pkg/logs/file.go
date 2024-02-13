package logs

import (
	"encoding/json"
	"fmt"

	"github.com/kitabisa/teler/common"
)

// File write detected threats into it
func File(options *common.Options, data map[string]string) error {
	var out string
	file := options.Configs.Logs.File

	if options.Output != nil {
		if file.JSON {
			logJSON, err := json.Marshal(data)
			if err != nil {
				return err
			}
			out = fmt.Sprintf("%s\n", logJSON)
		} else {
			out = fmt.Sprintf("[%s] [%s] [%s] %s\n",
				data["time_local"],
				data["remote_addr"],
				data["category"],
				data[data["element"]],
			)
		}

		if _, write := options.Output.WriteString(out); write != nil {
			return write
		}
	}

	return nil
}
