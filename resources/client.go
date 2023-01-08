package resources

import (
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	_ "github.com/Azure/go-autorest/autorest"
)

type Client struct {
	TerraformBridge	*bridge.TerraformBridge

	Authorizer		autorest.Authorizer
	Subscriptions		[]string
	CurrentUseSubscription	string

	// TODO You can continue to refine your client

}

func (x Client) CopyWithSubscription(subscriptionId string) *Client {
	c := x
	c.CurrentUseSubscription = subscriptionId
	return &c
}

func newClient() (*Client, error) {
	authorizer, subscriptions, err := getSubscriptions()
	if err != nil {
		return nil, err
	}
	return &Client{
		Authorizer:	authorizer,
		Subscriptions:	subscriptions,
	}, nil
}

func getSubscriptions() (autorest.Authorizer, []string, error) {
	azureAuth, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		azureAuth, err = auth.NewAuthorizerFromEnvironment()
	}
	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()
	svc := subscription.NewSubscriptionsClient()
	svc.Authorizer = azureAuth
	res, err := svc.List(ctx)
	if err != nil {
		return nil, nil, err
	}
	subscriptions := make([]string, 0)
	for res.NotDone() {
		for _, sub := range res.Values() {
			switch sub.State {
			case subscription.Disabled:
				fmt.Println("Not fetching from subscription because it is disabled subscription", *sub.SubscriptionID)
			case subscription.Deleted:
				fmt.Println("Not fetching from subscription because it is deleted subscription", *sub.SubscriptionID)
			default:
				subscriptions = append(subscriptions, *sub.SubscriptionID)
			}
		}
		err := res.NextWithContext(ctx)
		if err != nil {
			return nil, nil, err
		}
	}
	return azureAuth, subscriptions, nil
}
