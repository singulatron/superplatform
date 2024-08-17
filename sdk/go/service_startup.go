package sdk

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/router"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

// RegisterService registers a service or logs in with credentials loaded
// from the credentials store.
// Every service should have a user account in the user service and this method creates
// that user account.
func RegisterService(serviceSlug, serviceName string, router *router.Router, store datastore.DataStore) (string, error) {
	res, err := store.Query(datastore.All()).Find()
	if err != nil {
		return "", err
	}

	slug := serviceSlug
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*usertypes.Credential)
		slug = cred.Slug
		pw = cred.Password
	} else {
		pw = uuid.New().String()
		err = store.Upsert(&usertypes.Credential{
			Slug:     slug,
			Password: pw,
		})
		if err != nil {
			return "", err
		}
	}

	loginRsp := usertypes.LoginResponse{}
	err = router.Post(context.Background(), "user-svc", "/login", usertypes.LoginRequest{
		Slug:     slug,
		Password: pw,
	}, &loginRsp)

	if err != nil {
		logger.Info(fmt.Sprintf("Registering the %v service", serviceSlug))

		err = router.Post(context.Background(), "user-svc", "/register", usertypes.RegisterRequest{
			Slug:     slug,
			Name:     serviceName,
			Password: pw,
		}, nil)
		if err != nil {
			return "", err
		}

		loginRsp = usertypes.LoginResponse{}
		err = router.Post(context.Background(), "user-svc", "/login", usertypes.LoginRequest{
			Slug:     slug,
			Password: pw,
		}, &loginRsp)
		if err != nil {
			return "", err
		}
	}

	return loginRsp.Token.Token, nil
}

func RegisterUser(router *router.Router, slug, password, username string) (string, error) {
	err := router.Post(context.Background(), "user-svc", "/register", &usertypes.RegisterRequest{
		Slug:     slug,
		Password: password,
		Name:     username,
	}, nil)
	if err != nil {
		return "", err
	}

	loginRsp := usertypes.LoginResponse{}
	err = router.Post(context.Background(), "user-svc", "/login", &usertypes.LoginRequest{
		Slug:     slug,
		Password: password,
	}, &loginRsp)
	if err != nil {
		return "", err
	}

	return loginRsp.Token.Token, nil
}
