package test

import (
	"context"
	"fmt"

	client "github.com/singulatron/superplatform/clients/go"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/router"
)

func Client(url string) *client.APIClient {
	cli := client.NewAPIClient(&client.Configuration{
		Servers: client.ServerConfigurations{
			{
				URL:         url,
				Description: "Default server",
			},
		},
	})

	return cli
}

func AdminClient(url string) (*client.APIClient, string, error) {
	cli := client.NewAPIClient(&client.Configuration{
		Servers: client.ServerConfigurations{
			{
				URL:         url,
				Description: "Default server",
			},
		},
	})
	userSvc := cli.UserSvcAPI

	adminLoginRsp, _, err := userSvc.Login(context.Background()).Request(client.UserSvcLoginRequest{
		Slug:     client.PtrString("singulatron"),
		Password: client.PtrString("changeme"),
	}).Execute()
	if err != nil {
		return nil, "", err
	}

	adminClient := client.NewAPIClient(&client.Configuration{
		Servers: client.ServerConfigurations{
			{
				URL:         url,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + *adminLoginRsp.Token.Token,
		},
	})

	return adminClient, *adminLoginRsp.Token.Token, nil
}

func MakeClients(router *router.Router, num int) ([]*client.APIClient, error) {
	var ret []*client.APIClient

	for i := 0; i < num; i++ {
		slug := fmt.Sprintf("test-user-slug-%v", i)
		password := fmt.Sprintf("testUserPassword%v", i)
		username := fmt.Sprintf("Test User Name %v", i)

		token, err := sdk.RegisterUser(router, slug, password, username)
		if err != nil {
			return nil, err
		}

		c := client.NewAPIClient(&client.Configuration{
			Servers: client.ServerConfigurations{
				{
					URL:         router.Address(),
					Description: "Default server",
				},
			},
			DefaultHeader: map[string]string{
				"Authorization": "Bearer " + token,
			},
		})

		ret = append(ret, c)
	}

	return ret, nil
}
