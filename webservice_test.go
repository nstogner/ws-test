package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

func TestWebService(t *testing.T) {
	var config struct {
		Directory      string        `envconfig:"directory" required:"true"`
		Search         string        `envconfig:"search" default:"**/request.*"`
		URL            string        `envconfig:"url" required:"true"`
		Method         string        `envconfig:"method" default:"POST"`
		RequestTimeout time.Duration `envconfig:"request_timeout" default:"3s"`
	}
	envconfig.MustProcess("TEST", &config)
	json.NewEncoder(os.Stdout).Encode(config)
	client := &http.Client{Timeout: config.RequestTimeout}

	requestFiles, err := filepath.Glob(filepath.Join(config.Directory, config.Search))
	require.NoError(t, err)

	for _, reqFile := range requestFiles {
		t.Run(reqFile, func(t *testing.T) {
			reqBody, err := ioutil.ReadFile(reqFile)
			require.NoError(t, err)

			respFile := strings.ReplaceAll(reqFile, "request", "response")
			expectedRespBody, err := os.Open(respFile)
			require.NoError(t, err)

			req, err := http.NewRequest(config.Method, config.URL, bytes.NewReader(reqBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", http.DetectContentType(reqBody))
			resp, err := client.Do(req)
			require.NoError(t, err)

			var obj, expectedObj interface{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&obj))
			require.NoError(t, json.NewDecoder(expectedRespBody).Decode(&expectedObj))

			require.EqualValues(t, expectedObj, obj)
		})
	}

}
