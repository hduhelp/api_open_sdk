package health

func (x *Content) MarshalJSON() ([]byte, error) {
	return []byte(x.Value), nil
}
