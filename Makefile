HOSTNAME=registry.terraform.io
NAMESPACE=magenta-aps
NAME=configupdater
VERSION=0.1.0
OS_ARCH=linux_amd64
PLUGIN_PATH=terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}


terraform-provider-configupdater: *.go
	go build -o $@


format:
	terraform fmt
	go fmt


.DEFAULT_GOAL := init
init: terraform-provider-configupdater main.tf
	mkdir -p ${PLUGIN_PATH}
	cp $< ${PLUGIN_PATH}
	-rm .terraform.lock.hcl
	terraform init
