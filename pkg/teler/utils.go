package teler

import (
	"reflect"
	"unicode/utf8"

	"teler.app/common"
	"teler.app/pkg/matchers"
	"teler.app/resource"
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

		datasets[cat] = make(map[string]string)
		datasets[cat]["content"] = con
	}
}

func trimFirst(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func isWhitelist(options *common.Options, find string) bool {
	whitelist := options.Configs.Rules.Threat.Whitelists
	for i := 0; i < len(whitelist); i++ {
		match := matchers.IsMatch(whitelist[i], find)
		if match {
			return true
		}
	}

	return false
}
