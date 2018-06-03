package main

type userService struct {

}

func (*userService) getUser(userId int64) User {
	return User{
		Id:        1,
		Name:      "admin",
		Email:     "admin@karspexet.se",
		SpexYears: 10,
	}
}

