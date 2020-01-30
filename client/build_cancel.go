package client // import "github.com/docker/docker/client"

import (
	"context"
	"net/url"
)

// BuildCancel requests the daemon to cancel ongoing build request
//  终止Build
func (cli *Client) BuildCancel(ctx context.Context, id string) error {
	//  用于解析url，
	query := url.Values{}
	query.Set("id", id)

	serverResp, err := cli.post(ctx, "/build/cancel", query, nil, nil)
	ensureReaderClosed(serverResp)
	return err
}
