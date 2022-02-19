package handlers

// MAP TYPE [JSON OBJ - Driver]
type Map map[string]interface{}

// [type-assertion]
func (m Map) M(key string) Map {return m[key].(map[string]interface{})}
func (m Map) S(key string) string {return m[key].(string)}
func (m Map) I(key string) int {return m[key].(int)}
func (m Map) A(key string) []interface{} {return m[key].([]interface{})}