package service

import (
	"context"
	"errors"

	"github.com/VILJkid/go-simple-grpc-demo/database"
	"github.com/VILJkid/go-simple-grpc-demo/proto"
)

type userServiceServer struct {
	proto.UserServiceServer
}

func NewUserServiceServer() *userServiceServer {
	return &userServiceServer{}
}

// GetUserById returns the requested single user detail
func (s *userServiceServer) GetUserById(ctx context.Context, req *proto.UserRequest) (*proto.User, error) {
	userId := req.GetUserId()
	if userId < 0 {
		return nil, errors.New("invalid user ID: " + string(userId) + ": must be non-negative")
	}

	db := database.MockDatabaseConnection()
	db.MockDatabaseBulkEntries()

	user, found := db.GetUserById(userId)
	if !found {
		return nil, errors.New("user not found for ID: " + string(userId))
	}
	return user, nil
}

// GetUsersByIds returns requested multiple user details
func (s *userServiceServer) GetUsersByIds(ctx context.Context, req *proto.UserListRequest) (*proto.UserList, error) {
	userIds := req.GetUserIds()
	if len(userIds) == 0 {
		return nil, errors.New("user IDs list is empty")
	}
	var userListDetails []*proto.User
	var err error

	db := database.MockDatabaseConnection()
	db.MockDatabaseBulkEntries()

	for _, userId := range userIds {
		user, found := db.GetUserById(userId)
		if userId < 0 {
			err = errors.Join(err, errors.New("invalid user ID: "+string(userId)+": must be non-negative"))
			continue
		}
		if !found {
			err = errors.Join(err, errors.New("user not found for ID: "+string(userId)))
			continue
		}
		userListDetails = append(userListDetails, user)
	}

	if err != nil {
		return nil, err
	}

	return &proto.UserList{
		Users: userListDetails,
	}, nil
}
