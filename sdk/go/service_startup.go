package sdk

import (
	"context"
	"fmt"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/logger"
	"github.com/singulatron/superplatform/sdk/go/router"

	client "github.com/singulatron/superplatform/clients/go"
)

// RegisterService registers a service or logs in with credentials loaded from the credentials store.
// Every service should have a user account in the user service and this method creates that user account.
func RegisterService(serviceSlug, serviceName string, router *router.Router, store datastore.DataStore) (string, error) {
	res, err := store.Query().Find()
	if err != nil {
		return "", err
	}

	slug := serviceSlug
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*Credential)
		slug = cred.Slug
		pw = cred.Password
	} else {
		pw = Id("cred")
		err = store.Upsert(&Credential{
			Slug:     slug,
			Password: pw,
		})
		if err != nil {
			return "", err
		}
	}

	loginRsp := client.UserSvcLoginResponse{}
	err = router.Post(context.Background(), "user-svc", "/login", client.UserSvcLoginRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(pw),
	}, &loginRsp)

	if err != nil {
		logger.Info(fmt.Sprintf("Registering the %v service", serviceSlug))

		err = router.Post(context.Background(), "user-svc", "/register", client.UserSvcRegisterRequest{
			Slug:     client.PtrString(slug),
			Name:     client.PtrString(serviceName),
			Password: client.PtrString(pw),
		}, nil)
		if err != nil {
			return "", err
		}

		loginRsp = client.UserSvcLoginResponse{}
		err = router.Post(context.Background(), "user-svc", "/login", client.UserSvcLoginRequest{
			Slug:     client.PtrString(slug),
			Password: client.PtrString(pw),
		}, &loginRsp)
		if err != nil {
			return "", err
		}
	}

	return *loginRsp.Token.Token, nil
}

func RegisterServiceNoRouter(userService *client.UserSvcAPIService, serviceSlug, serviceName string, store datastore.DataStore) (string, error) {
	ctx := context.Background()

	res, err := store.Query().Find()
	if err != nil {
		return "", err
	}

	slug := serviceSlug
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*Credential)
		slug = cred.Slug
		pw = cred.Password
	} else {
		pw = Id("cred")
		err = store.Upsert(&Credential{
			Slug:     slug,
			Password: pw,
		})
		if err != nil {
			return "", err
		}
	}

	loginRsp, _, err := userService.Login(ctx).Request(client.UserSvcLoginRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(pw),
	}).Execute()

	if err != nil {
		logger.Info(fmt.Sprintf("Registering the %v service", serviceSlug))

		_, _, err = userService.Register(ctx).Body(client.UserSvcRegisterRequest{
			Slug:     client.PtrString(slug),
			Name:     client.PtrString(serviceName),
			Password: client.PtrString(pw),
		}).Execute()
		if err != nil {
			return "", err
		}

		loginRsp, _, err = userService.Login(ctx).Request(client.UserSvcLoginRequest{
			Slug:     client.PtrString(slug),
			Password: client.PtrString(pw),
		}).Execute()
		if err != nil {
			return "", err
		}
	}

	return *loginRsp.Token.Token, nil
}

func RegisterUser(router *router.Router, slug, password, username string) (string, error) {
	err := router.Post(context.Background(), "user-svc", "/register", &client.UserSvcRegisterRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(password),
		Name:     client.PtrString(username),
	}, nil)
	if err != nil {
		return "", err
	}

	loginRsp := client.UserSvcLoginResponse{}
	err = router.Post(context.Background(), "user-svc", "/login", &client.UserSvcLoginRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(password),
	}, &loginRsp)
	if err != nil {
		return "", err
	}

	return *loginRsp.Token.Token, nil
}
