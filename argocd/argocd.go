package argocd

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stefanmcshane/go-argocd/argocd/client"
)

// ArgoCDClientOpts allow for customizing the connection method to ArgoCD
type ArgoCDClientOpts struct {
	Host     string
	Port     string
	APIToken string

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

// NewArgoCDWithAPIKey returns a client which will connect to the given
// hostname (and optionally port), and will set the expected Authorization
// header on each request
func NewArgoCDWithAPIKey(conf ArgoCDClientOpts) (*client.ConsolidateServices, error) {
	if conf.Host == "" {
		return nil, errors.New("must provide host")
	}
	if conf.Port == "" || conf.Port == "0" {
		return nil, errors.New("must provide non-zero port")
	}
	if conf.APIToken == "" {
		return nil, errors.New("must provide valid APIToken")
	}
	if conf.Scheme == "" {
		conf.Scheme = "https"
	}

	t := httptransport.New(net.JoinHostPort(conf.Host, conf.Port), client.DefaultBasePath, []string{conf.Scheme})
	t.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Bearer %s", conf.APIToken))
	t.Debug = conf.Debug

	if conf.Transport != nil {
		t.Transport = conf.Transport
	}

	argoClient := client.New(t, strfmt.Default)

	return argoClient, nil
}
