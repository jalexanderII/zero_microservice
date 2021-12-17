package external_apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ResponseObj interface {
}

// MakeGet gets response from url
func MakeGet(ctx context.Context, url string, obj interface{}, verbose bool) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if verbose {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		body := strings.Trim(string(data), " []")
		if body == "" {
			return nil
		}
		fmt.Printf("body: %s\n", body)

		if err := json.Unmarshal([]byte(body), obj); err != nil {
			fmt.Printf("payload: %s\n", body)
			return err
		}
	} else {
		err = json.NewDecoder(resp.Body).Decode(obj)
	}

	return nil
}
