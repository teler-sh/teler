package runner

import (
	"reflect"

	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/logs"
)

func log(options *common.Options, data map[string]string) {
	m := options.Configs.Logs
	v := reflect.ValueOf(m)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).FieldByName("Active").Bool() {
			switch t.Field(i).Name {
			case "File":
				toFile(options, data)
				// case "Zinc": // TODO
				// 	toZinc(options, data)
			}
		}
	}
}

func toFile(options *common.Options, data map[string]string) {
	err := logs.File(options, data)
	if err != nil {
		errors.Show(err.Error())
	}
}
