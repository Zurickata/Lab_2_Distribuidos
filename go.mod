module github.com/Zurickata/Lab_2_Distribuidos

go 1.13

require (
	golang.org/x/net v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

replace (
    github.com/Zurickata/Lab_2_Distribuidos/proto => ../proto
)