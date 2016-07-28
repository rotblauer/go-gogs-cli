// This sucker is really annoying. Gogs Client says that search is a thing.
// And it is, it does work.
// But... I can't find it in the package. It's just not there. I've looked.
// So this is a re-write of the respondering methods that are nonexported in
// the Gogs Client package. Here they're used exclusively (and tweaked a little, as such)
// for the search command.
// It's dumb. I know.

package cmd

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	gogs "github.com/gogits/go-gogs-client"
	"github.com/spf13/viper"
)

type expRes struct {
	Data []*gogs.Repository `json:"data"`
}

func getResponse(client *http.Client, method, url string, header http.Header, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+viper.GetString("token"))
	for k, v := range header {
		req.Header[k] = v
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 403:
		return nil, errors.New("403 Forbidden")
	case 404:
		return nil, errors.New("404 Not Found")
	}

	if resp.StatusCode/100 != 2 {
		errMap := make(map[string]interface{})
		if err = json.Unmarshal(data, &errMap); err != nil {
			return nil, err
		}
		return nil, errors.New(errMap["message"].(string))
	}

	return data, nil
}

func getParsedResponse(client *http.Client, method, path string, header http.Header, body io.Reader) ([]*gogs.Repository, error) {
	data, err := getResponse(client, method, path, header, body)
	if err != nil {
		return nil, err
	}
	var res expRes
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}
