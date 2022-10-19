tools:
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.3

generate: tools
	mkdir -p argocd
	swagger generate client --target=./argocd

clean:
	rm -rf argocd/client argocd/models