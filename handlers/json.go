package handlers

// import (
// 	"encoding/json"
// )

// // JSON TYPE [Middleware to the rest of datatypes]
// type Json []byte
// // [constructor]
// func JsonNew(data interface{}) (Json, error) {
// 	return json.Marshal(data)
// }
// func JsonNewFromFile(filename string) (Json, error) {
// 	bytes, err := FileToBytes(filename)
// 	if err != nil { return nil, err }

// 	return json.Marshal(bytes)
// }
// // [methods]
// func (j *Json) To_Map() *Map {
// 	var data = &Map{}
// 	if err := json.Unmarshal([]byte(*j), &data); err != nil {
// 		return nil
// 	}
// 	return data
// }
// func (j *Json) To_String() string {
// 	return string(*j)
// }