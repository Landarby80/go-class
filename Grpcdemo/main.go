package main
/*
Case 1  : If you wants to genarate *.pb.go in same directory 
Run     : protoc --go_out=. *.proto (In the directory where *.proto files exist)

Case 2  : If you wants to genarate *.pb.go file to a directory(pb) other then proto folder 
         (proto folder contain *.proto)     (pb is target folder where *.pb.go file genarate
        )
         protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

*/
import (
	"GoGRPCConnect/proto"
)
func main(){

}