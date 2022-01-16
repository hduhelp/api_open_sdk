package types

func (x *JSON) MarshalJSON() ([]byte, error) {
	return []byte(x.Value), nil
}

func JSONFromString(str string) *JSON {
	return &JSON{Value: str}
}
