package gatev1

func (x Service) Chinese() string {
	if name, ok := map[Service]string{
		Service_SERVICE_HIKVISION: "海康闸机",
		Service_SERVICE_UNIVIEW:   "宇视闸机",
	}[x]; ok {
		return name
	}
	return "未知服务"
}
