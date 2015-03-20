package results

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	tagKey = "nbagame"
)

var (
	resultSetNameToType = map[string]interface{}{
		"common_all_players": CommonAllPlayersRow{},
	}
)

// Response represents a results from the stats.nba.com API.
type Response struct {
	Resource   string                 `json:"resource"`
	Parameters map[string]interface{} `json:"parameters"`
	ResultSets []*ResultSet           `json:"resultSets"`
}

// ResultSet is a set of results returned for a resource from the stats.nba.com
// API.
type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}

// NewResponse constructs a new Response from a byte array contianing the returned
// json body.
func NewResponse(body []byte) (*Response, error) {
	var resp Response
	return &resp, json.Unmarshal(body, &resp)
}

// Decode decodes the ResultSet into a slice of the appropriate result set
// structs.
func (rs *ResultSet) Decode(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidDecodeError{rv}
	}

	slice := rv.Elem()
	if slice.Kind() != reflect.Slice {
		return &InvalidDecodeError{rv}
	}

	// TODO: More error checking on v.
	sliceType := slice.Type().Elem().Elem()

	headers := rs.makeHeaderMap()
	for _, row := range rs.RowSet {
		if len(row) != len(rs.Headers) {
			return fmt.Errorf("ResultSet headers contain %v columns, row contains %v", len(rs.Headers), len(row))
		}

		newValue := reflect.New(sliceType)
		if err := populateValue(newValue.Elem(), row, headers); err != nil {
			return err
		}
		slice.Set(reflect.Append(slice, newValue))
	}

	return nil
}

func (rs *ResultSet) makeHeaderMap() map[string]int {
	m := make(map[string]int)
	for k, h := range rs.Headers {
		m[h] = k
	}
	return m
}

func populateValue(v reflect.Value, row []interface{}, headers map[string]int) error {
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)

		fieldHeader := fieldType.Tag.Get(tagKey)
		if fieldHeader == "" {
			// Skip any fields that don't have a tag.
			continue
		}

		headerIndex, ok := headers[fieldHeader]
		if !ok {
			return fmt.Errorf("Header `%s` does not exist in %+v", fieldHeader, headers)
		}

		rowValue := reflect.ValueOf(row[headerIndex])
		if !rowValue.IsValid() {
			continue
		}
		rowValueType := rowValue.Type()
		fieldValueType := fieldValue.Type()
		if !rowValueType.ConvertibleTo(fieldValueType) {
			return fmt.Errorf("Cannot convert `%v` to `%v`", rowValue.Type(), fieldValue.Type())
		}
		fieldValue.Set(rowValue.Convert(fieldValue.Type()))
	}
	return nil
}

// An InvalidDecodeError describes an invalid argument passed to Decode.
// (The argument to Decode must be a non-nil pointer to a slice.)
type InvalidDecodeError struct {
	Value reflect.Value
}

func (e *InvalidDecodeError) Error() string {
	if e.Value.IsNil() {
		return "nbagame.results: Decode(nil " + e.Value.Type().String() + ")"
	}
	if e.Value.Kind() != reflect.Ptr {
		return "nbagame.results: Decode(non-pointer " + e.Value.Type().String() + ")"
	}
	if e.Value.Elem().Kind() != reflect.Slice {
		return "nbagame.results: Decode(pointer to non-slice " + e.Value.Type().String() + ")"
	}

	return "nba.results: Decode(" + e.Value.String() + ")"
}
