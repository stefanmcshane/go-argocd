package argocd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stefanmcshane/go-argocd/argocd/client"
	"github.com/stefanmcshane/go-argocd/argocd/client/session_service"
	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// ArgoCDClientOpts allow for customizing the connection method to ArgoCD
type ArgoCDClientOpts struct {
	Host string
	Port string

	// APIToken is the preferred way to communicate with the ArgoCD API.
	// Either APIToken or Username/Password must be supplied.
	// If all 3 are supplied, the APIToken method will be preferred over Username/Password
	APIToken string

	// Username is the username used for communication with the ArgoCD API.
	// If Username is supplied, Password must also be supplied.
	// Either APIToken or Username/Password must be supplied.
	// If all 3 are supplied, the APIToken method will be preferred over Username/Password
	Username string

	// Password is the password used for communication with the ArgoCD API.
	// If Password is supplied, Username must also be supplied.
	// Either APIToken or Username/Password must be supplied.
	// If all 3 are supplied, the APIToken method will be preferred over Username/Password
	Password string

	// Scheme allows for overwriting the connection scheme when connecting to ArgoCD
	// This will default to `https`
	Scheme string

	// Transport allows for overriding HTTP Transport
	// This can be used to set InsecureSkipVerify for example
	// Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	Transport *http.Transport

	// Debug will default to false. Setting this to true will give more context into any calls made to ArgoCD
	Debug bool
}

// ArgoCDClient is a convenience wrapper for accessing ArgoCD API
type ArgoCDClient struct {
	ArgoCD *client.ConsolidateServices
}

// NewArgoCD returns a client which will connect to the given
// hostname and port, and also set the expected Authorization
// header on each request. If Username/Password is provided, this will be
// exchanged for an APIToken for subsequent requests.
// If APIToken, Username, and Password are all supplied, the APIToken method will be preferred over Username/Password
func NewArgoCD(ctx context.Context, conf ArgoCDClientOpts) (*client.ConsolidateServices, error) {
	if conf.Host == "" {
		return nil, errors.New("must provide host")
	}
	if conf.Port == "" || conf.Port == "0" {
		return nil, errors.New("must provide non-zero port")
	}
	if conf.APIToken == "" {
		if conf.Username == "" && conf.Password == "" {
			return nil, errors.New("must provide valid APIToken, or Username and Password")
		}
	}
	if conf.Scheme == "" {
		conf.Scheme = "https"
	}

	t := httptransport.New(net.JoinHostPort(conf.Host, conf.Port), client.DefaultBasePath, []string{conf.Scheme})
	t.Debug = conf.Debug

	var argoClient *client.ConsolidateServices
	if conf.APIToken != "" {
		t.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Bearer %s", conf.APIToken))
		argoClient = client.New(t, strfmt.Default)
	} else if conf.Username != "" && conf.Password != "" {
		temporaryArgoClient := client.New(t, strfmt.Default)

		sessionParams := &session_service.SessionServiceCreateParams{
			Context: ctx,
			Body: &models.SessionSessionCreateRequest{
				Username: conf.Username,
				Password: conf.Password,
			},
		}
		sessionResp, err := temporaryArgoClient.SessionService.SessionServiceCreate(sessionParams)
		if err != nil {
			return nil, fmt.Errorf("error getting APIToken: %w", err)
		}
		conf.APIToken = sessionResp.Payload.Token
		return NewArgoCD(ctx, conf)
	}

	if conf.Transport != nil {
		t.Transport = conf.Transport
	}

	return argoClient, nil
}
