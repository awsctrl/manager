
# Image URL to use all building/pushing image targets
IMG ?= r.awsctrl.io/controller:latest

all: test manager

# Run tests
test: generate fmt vet manifests
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager awsctrl.io/cmd/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crds
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd \
		paths=./pkg/apis/... \
		output:crd:dir=./config/crds/ \
		crd:trivialVersions=true

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go paths=./pkg/apis \
		object:headerFile=./hack/boilerplate.go.txt

# Build the docker image
docker-build: test
	docker build . -t ${IMG}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
docker-push:
	docker push ${IMG}

install-ci:
	curl -sL https://go.kubebuilder.io/dl/2.0.0/linux/amd64 | tar -xz -C /tmp/
	sudo mv /tmp/kubebuilder_2.0.0_linux_amd64 /usr/local/kubebuilder
	export PATH=\$PATH:/usr/local/kubebuilder/bin
	curl -s https://api.github.com/repos/kubernetes-sigs/kustomize/releases/latest |\
  	grep browser_download |\
 	grep linux |\
  	cut -d '"' -f 4 |\
  	xargs curl -O -L
	mv kustomize_*_linux_amd64 kustomize
	chmod u+x kustomize

# Create kind cluster for testing
kind-create:
	kind create cluster --name awsctrl --config config/kind/kind.yaml

# Delete kind cluster for testing
kind-delete:
	kind delete cluster --name awsctrl 