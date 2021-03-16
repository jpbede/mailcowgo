package transport

import "strconv"

// StrInt64 tries to make a non-type-safe API type safe
// e.g. API returns sometimes a int as string
type StrInt64 int64

// UnmarshalJSON implements Unmarshaler
func (s *StrInt64) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		*s = -1
		return nil
	}

	if string(data) == "false" {
		*s = 0
		return nil
	}

	var parse string
	if string(data[0]) == "\"" && string(data[len(data)-1]) == "\"" {
		parse = string(data[1 : len(data)-1])
	} else {
		parse = string(data)
	}

	// try to parse json data
	v, err := strconv.ParseInt(parse, 10, 64)
	if err != nil { // error cause we have weird data, just set as 0
		v = 0
	}

	*s = StrInt64(v)
	return nil
}
