package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/stefanmcshane/go-argocd/argocd"
	"github.com/stefanmcshane/go-argocd/argocd/client/application_service"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func run() error {
	argo, err := NewArgoCD()
	if err != nil {
		return err
	}

	_ = argo
	return nil
}

func NewArgoCD() (string, error) {
	var cli string

	argoConf := argocd.ArgoCDClientOpts{
		Host:      "localhost",
		Port:      "8080",
		APIToken:  "YOUR_API_TOKEN",
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Debug:     false,
	}

	argoClient, err := argocd.NewArgoCDWithAPIKey(argoConf)
	if err != nil {
		return cli, err
	}

	params := application_service.ApplicationServiceListParams{
		Context: context.Background(),
	}
	resp, err := argoClient.ApplicationService.ApplicationServiceList(&params)
	if err != nil {
		return cli, err
	}

	fmt.Println(resp.Payload.Items[0].Metadata)

	return cli, nil

}
