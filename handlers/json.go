package handlers

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

// MAP TYPE [json-obj driver]
type Map map[string]interface{}

// type assertion
func (m Map) M(key string) Map {return m[key].(Map)}
func (m Map) S(key string) string {return m[key].(string)}
func (m Map) I(key string) int {return m[key].(int)}
func (m Map) A(key string) []interface{} {return m[key].([]interface{})}
// methods
func (m *Map) to_json() ([]byte, error) {
	return json.Marshal(m)
}
func (m *Map) to_string() (string, error) {
	bytes, err := m.to_json()
	if err != nil {return "", err}
	return string(bytes), nil
}



// JSON STRUCT
type JsonFile struct {
	File	string
	Data	Map
}
func (j *JsonFile) Parse() error {
	// File [open]
	file, err := os.Open(j.File)
	if err != nil { return err }

	// File [read]
	// Bytes -> Map[json-obj driver]
	bytes, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal([]byte(bytes), &j.Data); err != nil {
		return err
	}

	// File [close]
	file.Close()
	return nil
}
func (j *JsonFile) Clean() {
	j.Data = nil
}