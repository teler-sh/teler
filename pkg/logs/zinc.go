package logs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	e "ktbs.dev/teler/pkg/errors"
)

// Zinc logs insertion
func Zinc(base string, index string, auth string, log map[string]string) error {
	var res map[string]string
	client := &http.Client{}

	data, err := json.Marshal(log)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprint(base, "/api/", index, "/document"), bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return err
	}

	if res["id"] != "" {
		return nil
	}

	return errors.New(e.ErrInsertLogZinc)
}
