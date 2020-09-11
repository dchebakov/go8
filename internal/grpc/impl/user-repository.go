package impl

import (
	grpc2 "google.golang.org/grpc"
)

type RepositoryServiceGrpcImpl struct {
}

//NewRepositoryServiceGrpcImpl returns the pointer to the implementation.
func NewRepositoryServiceGrpcImpl(*grpc2.ClientConn) *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}
