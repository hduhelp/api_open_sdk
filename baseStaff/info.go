package baseStaff

import "encoding/json"

//type InfoMapList map[string]*Info

func (m InfoMapList) Append(list ...*Info) {
	for _, v := range list {
		m.Items[v.StaffID] = v
	}
}

func (m InfoMapList) MarshalJSON() ([]byte, error) {
	list := make([]*Info, 0)
	for _, v := range m.Items {
		list = append(list, v)
	}
	return json.Marshal(list)
}
