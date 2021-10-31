package baseStaff

import "encoding/json"

func (x *InfoMapList) Append(list ...*Info) {
	for _, v := range list {
		x.Items[v.StaffID] = v
	}
}

func (x *InfoMapList) MarshalJSON() ([]byte, error) {
	list := make([]*Info, 0)
	for _, v := range x.Items {
		list = append(list, v)
	}
	return json.Marshal(list)
}

func (x *InfoMapList) UnmarshalJSON(bytes []byte) error {
	list := make([]*Info, 0)
	err := json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}
	x.Items = make(map[string]*Info)
	for _, v := range list {
		x.Items[v.StaffID] = v
	}
	return nil
}
