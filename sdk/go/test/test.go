package test

import (
	"context"
	"fmt"

	client "github.com/singulatron/singulatron/clients/go"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/router"
)

func AdminClient(url string, username, password string) (*client.APIClient, error) {
	cli := client.NewAPIClient(&client.Configuration{
		Servers: client.ServerConfigurations{
			{
				URL:         url,
				Description: "Default server",
			},
		},
	})
	userSvc := cli.UserSvcAPI

	adminLoginRsp, _, err := userSvc.Login(context.Background()).Request(clients.UserSvcLoginRequest{
		Slug:     client.PtrString("singulatron"),
		Password: client.PtrString("changeme"),
	}).Execute()
	if err != nil {
		return nil, err
	}

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
