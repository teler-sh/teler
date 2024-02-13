package teler

import (
	"reflect"
	"unicode/utf8"

	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/pkg/matchers"
	"github.com/kitabisa/teler/resource"
)

func getDatasets() {
	datasets = make(map[string]map[string]string)

	rsc := resource.Get()
	for _, threat := range rsc.Threat {
		thr := reflect.ValueOf(threat)
		cat := thr.FieldByName("Category").String()
		con := thr.FieldByName("Content").String()

		if thr.FieldByName("Exclude").Bool() {
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
	for _, item := range whitelist {
		match := matchers.IsMatch(item, find)
		if match {
			return true
		}
	}

	return false
}
