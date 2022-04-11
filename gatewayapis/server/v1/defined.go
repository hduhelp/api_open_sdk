package serverv1

import "google.golang.org/grpc"

func (x *PostRegistGRPCMethodsRequest) FormGRPCServer(server interface {
	GetServiceInfo() map[string]grpc.ServiceInfo
}) {
	x.Methods = make([]*GRPCMethod, 0)
	for serviceName, v := range server.GetServiceInfo() {
		for _, method := range v.Methods {
			x.Methods = append(x.Methods, &GRPCMethod{
				ServiceName:    serviceName,
				MethodName:     method.Name,
				IsClientStream: method.IsClientStream,
				IsServerStream: method.IsServerStream,
			})
		}
	}
}
