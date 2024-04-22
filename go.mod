module github.com/Zurickata/Lab_2_Distribuidos

go 1.13

require (
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

replace (
    github.com/Zurickata/Lab_2_Distribuidos/proto => ../proto
)