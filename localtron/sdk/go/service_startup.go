package sdk

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/router"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// RegisterService registers a service or logs in with credentials loaded
// from the credentials store.
// Every service should have a user account in the user service and this method creates
// that user account.
func RegisterService(serviceEmail, serviceName string, router *router.Router, store datastore.DataStore) (string, error) {
	res, err := store.Query(datastore.All()).Find()
	if err != nil {
		return "", err
	}

	email := serviceEmail
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*usertypes.Credential)
		email = cred.Email
		pw = cred.Password
	} else {
		logger.Info(fmt.Sprintf("Registering the %v service", serviceEmail))

		pw = uuid.New().String()
		err = router.Post(context.Background(), "user-svc", "/register", usertypes.RegisterRequest{
			Email:    email,
			Name:     serviceName,
			Password: pw,
		}, nil)
		if err != nil {
			return "", err
		}
		err = store.Upsert(&usertypes.Credential{
			Email:    email,
			Password: pw,
		})
		if err != nil {
			return "", err
		}
	}

	rsp := usertypes.LoginResponse{}
	err = router.Post(context.Background(), "user-svc", "/login", usertypes.LoginRequest{
		Email:    email,
		Password: pw,
	}, &rsp)
	if err != nil {
		return "", err
	}

	return rsp.Token.Token, nil
}

func RegisterUser(router *router.Router, email, password, username string) (string, error) {
	err := router.Post(context.Background(), "user-svc", "/register", &usertypes.RegisterRequest{
		Email:    email,
		Password: password,
		Name:     username,
	}, nil)
	if err != nil {
		return "", err
	}

	loginRsp := usertypes.LoginResponse{}
	err = router.Post(context.Background(), "user-svc", "/login", &usertypes.LoginRequest{
		Email:    email,
		Password: password,
	}, &loginRsp)
	if err != nil {
		return "", err
	}

	return loginRsp.Token.Token, nil
}
