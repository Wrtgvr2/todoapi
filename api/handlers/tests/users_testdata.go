package handlers_test

import (
	"github.com/wrtgvr/todoapi/models"
)

var (
	testUsername        string = "testUser"
	testPassword        string = "testPassword"
	testUserID          uint64 = 1
	testUsername_BadReq string = "err"
	testPassword_BadReq string = "err"
)

var testUserReqData = models.UserRequest{
	Username:        &testUsername,
	DisplayUsername: &testUsername,
	Password:        &testPassword,
}

var testUserRespData = models.UserResponse{
	ID:              testUserID,
	Username:        testUsername,
	DisplayUsername: testUsername,
}

var testUserData = models.User{
	ID:              testUserID,
	Username:        testUsername,
	DisplayUsername: testUsername,
	Password:        testPassword,
}
