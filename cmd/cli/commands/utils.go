package commands

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"

	"github.com/livekit/livekit-server/pkg/auth"
)

var (
	roomFlag = &cli.StringFlag{
		Name:  "room",
		Usage: "name or id of the room",
	}
	roomHostFlag = &cli.StringFlag{
		Name:  "host",
		Value: "http://localhost:7880",
	}
	rtcHostFlag = &cli.StringFlag{
		Name:  "host",
		Value: "ws://localhost:7881",
	}
	apiKeyFlag = &cli.StringFlag{
		Name:     "api-key",
		EnvVars:  []string{"LK_API_KEY"},
		Required: true,
	}
	secretFlag = &cli.StringFlag{
		Name:     "api-secret",
		EnvVars:  []string{"LK_API_SECRET"},
		Required: true,
	}
)

func PrintJSON(obj interface{}) {
	txt, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(txt))
}

func ExpandUser(p string) string {
	if strings.HasPrefix(p, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, p[1:])
	}

	return p
}

func accessToken(c *cli.Context, grant *auth.VideoGrant, identity string) (value string, err error) {
	apiKey := c.String("api-key")
	apiSecret := c.String("api-secret")
	if apiKey == "" && apiSecret == "" {
		// not provided, don't sign request
		return
	}

	at := auth.NewAccessToken(apiKey, apiSecret).
		AddGrant(grant).
		SetIdentity(identity)
	return at.ToJWT()
}
