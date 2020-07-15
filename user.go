package fruitshop

import (
	"context"
	user "fruitshop/gen/user"
	"log"
)

// user service implementation.
// These methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Add implements add feature for the user to the user table
func (s *usersrvc) Add(ctx context.Context, p *user.AddPayload) (res *user.UserManagement, err error) {
	res = &user.UserManagement{}

	s.logger.Print("user.add started")
	newUser := user.UserManagement{
		ID:     p.UserID + p.UserID,
		UserID: p.UserID,
	}
	result, err := CreateUser(&newUser)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}

	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.add completed")
	return &result, err
}

// Get implements get feature for the user from the user table
func (s *usersrvc) Get(ctx context.Context, p *user.GetPayload) (res *user.UserManagement, err error) {
	res = &user.UserManagement{}
	s.logger.Print("user.get")
	newUser := user.UserManagement{
		ID:     p.UserID + p.UserID,
		UserID: p.UserID,
	}
	result, err := GetUser(&newUser)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.get completed")
	return &result, err
}

// Show implements show feature for all users in the user table
func (s *usersrvc) Show(ctx context.Context) (res user.UserManagementCollection, err error) {
	s.logger.Print("user.show started")
	res, err = ListUsers()
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.show completed")
	return
}
