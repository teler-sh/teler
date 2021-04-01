package teler

import (
	"reflect"

	"ktbs.dev/teler/resource"
)

func getDatasets() {
	datasets = make(map[string]map[string]string)
	rsc := resource.Get()
	for i := 0; i < len(rsc.Threat); i++ {
		threat := reflect.ValueOf(&rsc.Threat[i]).Elem()
		cat := threat.FieldByName("Category").String()
		con := threat.FieldByName("Content").String()

		if threat.FieldByName("Exclude").Bool() {
			continue
		}

		datasets[cat] = map[string]string{}
		datasets[cat]["content"] = con
	}
}
