package database

import (
	"sync"

	"github.com/VILJkid/go-simple-grpc-demo/proto"
)

type userDatabase struct {
	userDetails map[int32]*proto.User
	mu          sync.RWMutex
}

func MockDatabaseConnection() *userDatabase {
	return &userDatabase{}
}

func (udb *userDatabase) MockDatabaseBulkEntries() {
	udb.userDetails = make(map[int32]*proto.User)

	udb.userDetails[0] = &proto.User{
		Id:      0,
		Fname:   "Lilah",
		City:    "Jakarta",
		Phone:   1234567890,
		Height:  5.8,
		Married: false,
	}
	udb.userDetails[1] = &proto.User{
		Id:      1,
		Fname:   "Darwin",
		City:    "Hangzhou",
		Phone:   9912672341,
		Height:  5.2,
		Married: true,
	}
	udb.userDetails[2] = &proto.User{
		Id:      2,
		Fname:   "Leon",
		City:    "Zunyi",
		Phone:   7825362733,
		Height:  6.0,
		Married: false,
	}
	udb.userDetails[3] = &proto.User{
		Id:      3,
		Fname:   "Paxton",
		City:    "Belgrade",
		Phone:   9027336214,
		Height:  5.5,
		Married: true,
	}
	udb.userDetails[4] = &proto.User{
		Id:      4,
		Fname:   "Peter",
		City:    "Abu Dhabi",
		Phone:   2736225318,
		Height:  7.2,
		Married: false,
	}
}

func (udb *userDatabase) GetUserById(userId int32) (*proto.User, bool) {
	udb.mu.Lock()
	defer udb.mu.Unlock()
	user, found := udb.userDetails[userId]
	return user, found
}
