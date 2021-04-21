package flatten

// func oneToFlat(in interface{}) []interface{} {
// 	if in == nil {
// 		return []interface{}{}
// 	}
// 	return []interface{}{in}
// }

// func manyToFlat(in interface{}) []interface{} {
// 	out := []interface{}{}
// 	for _, v := range in.([]interface{}) {
// 		out = append(out, Flatten(v)...)
// 	}
// 	return out
// }

// // Flatten returns flattened array
// func Flatten(in interface{}) []interface{} {
// 	if _, ok := in.([]interface{}); !ok {
// 		return oneToFlat(in)
// 	}

// 	return manyToFlat(in)
// }

// Flatten returns flattened array
func Flatten(in interface{}) []interface{} {
	var out = []interface{}{}
	for _, e := range in.([]interface{}) {
		switch v := e.(type) {
		case nil:
		case []interface{}:
			out = append(out, Flatten(v)...)
		default:
			out = append(out, v)
		}
	}
	return out
}
