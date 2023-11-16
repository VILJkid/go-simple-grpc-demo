package main

import (
	"context"
	"testing"

	"github.com/VILJkid/go-simple-grpc-demo/proto"
	"github.com/VILJkid/go-simple-grpc-demo/service"
)

// TestGetUserByIdExist tests if the user exists
func TestGetUserByIdExist(t *testing.T) {
	server := service.NewUserServiceServer()

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

// TestGetUserByIdNotExist tests if the user doesn't exist
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

// TestGetUsersByIdsExist tests if multiple user exists
func TestGetUsersByIdsExist(t *testing.T) {
	server := service.NewUserServiceServer()

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

// TestGetUsersByIdsNotExist tests if some user doesn't exist
func TestGetUsersByIdsNotExist(t *testing.T) {
	server := service.NewUserServiceServer()

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
