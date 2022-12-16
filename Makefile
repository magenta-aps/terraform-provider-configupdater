HOSTNAME=registry.terraform.io
NAMESPACE=magenta-aps
NAME=configupdater
VERSION=0.1.0
OS_ARCH=linux_amd64
PLUGIN_PATH=terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}


terraform-provider-configupdater: *.go
	go build -o terraform-provider-configupdater
	mkdir -p ${PLUGIN_PATH}
	cp terraform-provider-configupdater ${PLUGIN_PATH}


format:
	terraform fmt
	go fmt


init: terraform-provider-configupdater main.tf
	-rm .terraform.lock.hcl
	terraform init
