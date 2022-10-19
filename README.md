# go-argocd

ArgoCD go module generated from ArgoCD swagger docs.

## Why not use ArgoCD/apiclient?

ArgoCD has a tight dependency with `k8s.io/kubernetes` through their gitops engine. This library decouples that approach due to the myriad of issues with that import

## Generating the library

- Copy the latest swagger.json from [ArgoCD](https://github.com/argoproj/argo-cd/blob/master/assets/swagger.json) to `./swagger.json`
- Ensure [swagger](https://github.com/go-swagger/go-swagger) is installed. This can be done through `make tools` also
- Run `make clean`
- Run `make generate`
