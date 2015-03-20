package endpoints

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

const (
	NBAStatsDomain = "stats.nba.com"
)

// DefaultRequester is the default Requester using default values for the
// endpoints.
var DefaultRequester = Requester{
	Domain:     NBAStatsDomain,
	PathPrefix: "stats",
}

// Requester performs requests to the stats.nba.com server's endpoints.
type Requester struct {
	Domain     string
	PathPrefix string
	client     http.Client
}

// EndpointURL returns the absolute URL for an endpoint.
func (r *Requester) EndpointURL(endpoint string) string {
	return fmt.Sprintf("http://%s/%s/%s", r.Domain, r.PathPrefix, endpoint)
}

// Request performs a request against the given endpoint with the provided
// parameters.
func (r *Requester) Request(endpoint string, params interface{}, resp interface{}) error {
	endpointURL, err := url.Parse(r.EndpointURL(endpoint))
	if err != nil {
		return err
	}

	urlParams, err := r.makeParams(params)
	if err != nil {
		return err
	}
	endpointURL.RawQuery = urlParams.Encode()

	req, err := http.NewRequest("GET", endpointURL.String(), nil)
	if err != nil {
		return err
	}
	httpResponse, err := r.client.Do(req)
	if err != nil {
		return err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("Endpoint `%s` returned status `%s`", endpointURL, httpResponse.Status)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(httpResponse.Body)
	if err != nil {
		return err
	}
	if err := httpResponse.Body.Close(); err != nil {
		return err
	}
	response, err := NewResponse(buf.Bytes())
	if err != nil {
		return err
	}

	return response.Decode(&resp)
}

func (r *Requester) makeParams(paramStruct interface{}) (url.Values, error) {
	params := url.Values{}
	rv := reflect.ValueOf(paramStruct)
	// Dereference the pointer if it's a pointer.
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		fieldType := rv.Type().Field(i)

		fieldHeader := fieldType.Tag.Get("json")
		if fieldHeader == "" {
			// Skip any fields that don't have a tag.
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			params.Set(fieldHeader, fieldValue.String())
			continue
		case reflect.Int:
			params.Set(fieldHeader, strconv.FormatInt(fieldValue.Int(), 10))
			continue
		}

		return nil, fmt.Errorf("Unsupported request parameter type: %s for tag `%s`", fieldValue.Kind(), fieldHeader)
	}

	return params, nil
}
