var store = make(map[string]string)

// Put function takes two parameters of type string as inputs  and returns an error type
func Put(key string, value string) error {
	store[key] = value

	return nil
}

var ErrorNoSuchKey = errors.New("no such key")

// Get function takes one parameter of type string as input and returns error and string types
func Get(key string) (string, error) {
	value, ok : store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)

	return nil
}