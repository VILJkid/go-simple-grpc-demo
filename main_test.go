package main

import (
	"context"
	"testing"

	"github.com/VILJkid/go-simple-grpc-demo/proto"
	"github.com/VILJkid/go-simple-grpc-demo/service"
)

func TestGetUserByIdExist(t *testing.T) {
	server := service.NewUserServiceServer()

	// Test case: User exists
	userID := int32(1)
	request := &proto.UserRequest{UserId: userID}
	response, err := server.GetUserById(context.Background(), request)

	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	if response.GetId() != userID {
		t.Errorf("Expected user ID %d, got %d", userID, response.GetId())
	}
}

func TestGetUserByIdNotExist(t *testing.T) {
	server := service.NewUserServiceServer()

	// Test case: User does not exist
	nonExistentUserID := int32(10)
	nonExistentRequest := &proto.UserRequest{UserId: nonExistentUserID}
	nonExistentResponse, err := server.GetUserById(context.Background(), nonExistentRequest)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if nonExistentResponse != nil {
		t.Errorf("nonExistentResponse should be nil, got %s", nonExistentResponse.String())
	}
}

func TestGetUsersByIdsExist(t *testing.T) {
	server := service.NewUserServiceServer()

	// Test case: Users exist
	userIDs := []int32{0, 2, 4}
	request := &proto.UserListRequest{UserIds: userIDs}
	response, err := server.GetUsersByIds(context.Background(), request)

	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	if len(response.GetUsers()) != len(userIDs) {
		t.Errorf("Expected %d users, got %d", len(userIDs), len(response.GetUsers()))
	}
}

func TestGetUsersByIdsNotExist(t *testing.T) {
	server := service.NewUserServiceServer()

	// Test case: Some users do not exist
	invalidUserIDs := []int32{0, 10, 4}
	invalidRequest := &proto.UserListRequest{UserIds: invalidUserIDs}
	invalidResponse, err := server.GetUsersByIds(context.Background(), invalidRequest)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if invalidResponse != nil {
		t.Errorf("invalidResponse should be nil, got %s", invalidResponse.String())
	}
}
