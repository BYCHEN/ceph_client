package s3

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

//DownloadObject ...
func (s *S3Config) DownloadObject(object string) (*http.Response, error) {
	res, err := s.Send(http.MethodGet, fmt.Sprintf("%s/%s", s.Bucket, object))
	if err != nil {
		return res, err
	}
	if res.StatusCode != 200 {
		defer res.Body.Close()
		body := S3Error{}
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		xml.Unmarshal(bodyBytes, &body)

		return res, errors.New(body.Code)
	}
	return res, err
}
