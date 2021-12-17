package dolyame

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func (c *Client) setAuth(r *http.Request) {
	r.Header.Add(
		"Authorization",
		fmt.Sprintf(
			"Basic %v",
			base64.RawStdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v:%v", c.login, c.password)),
			),
		),
	)
}
