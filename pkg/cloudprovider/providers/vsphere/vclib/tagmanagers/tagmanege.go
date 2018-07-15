package tagmanagers

import (
	"context"
	"net/url"

	"k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
)

type TagManagers struct {
	Client *RestClient
}

func NewTagManagers(connection *vclib.VSphereConnection) TagManagers {
	vsUrl := connection.Client.URL()
	vsUrl.User = url.UserPassword(connection.Username, connection.Password)

	return TagManagers{
		Client: NewClient(
			vsUrl,
			true,
			"",
		),
	}
}

func (t *TagManagers) WithLogout(ctx context.Context, f func(client *RestClient) error) error {
	err := t.Client.Login(ctx)

	if err != nil {
		return err
	}
	defer t.Client.Logout(ctx)

	return f(t.Client)

}
