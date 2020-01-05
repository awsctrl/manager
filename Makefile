
# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"
# Use AWS Client
USE_AWS_CLIENT ?= false
# Use existing Cluster
USE_EXISTING_CLUSTER ?= false
# Use Go Modules with the awsctrl
GO111MODULE ?= on

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager

# Stub for ci
test-unit: test

# Run tests
test: generate fmt vet manifests
	go test `go list ./... | grep -v e2e` -coverprofile unit.out -covermode atomic

test-e2e-%: 
ifeq (true,${USE_EXISTING_CLUSTER})
	export KUBECONFIG=${PWD}/kubeconfig-e2e-$*
	@$(MAKE) kind-create-awsctrl-$*
endif
	go test -coverprofile e2e-$*.out -covermode atomic -v -coverpkg ./controllers/... ./e2e/$*/...
ifeq (true,${USE_EXISTING_CLUSTER})
	@$(MAKE) kind-delete-awsctrl-$*
endif

# Run e2e
test-e2e: generate fmt vet manifests
	@$(MAKE) test-e2e-apigateway
	@$(MAKE) test-e2e-cloud9
	@$(MAKE) test-e2e-cloudformation
	@$(MAKE) test-e2e-ecr
	@$(MAKE) test-e2e-route53
	@$(MAKE) test-e2e-iam


# Build manager binary
manager: generate fmt vet
	go build -o bin/awsctrl main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go start

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: generator controller-gen
	$(GENERATOR) run
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

# Build the docker image
docker-build: test
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.4
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# Install CI will configure
install-ci: kubebuilder set-env kind #kubectl kubectl-context kubectl-verify 

# Install kind if not installed
kind:
ifeq (, $(shell which kind))
	go get sigs.k8s.io/kind
KIND=$(GOBIN)/kind
else
KIND=$(shell which kind)
endif

# Create kind cluster for testing
kind-create-%: kind
	$(KIND) create cluster --name $* --config config/kind/config.yaml -q

kind-create:
	@$(MAKE) kind-create-awsctrl

# Delete kind cluster for testing
kind-delete-%: kind
	$(KIND) delete cluster --name $* 

kind-delete:
	@$(MAKE) kind-delete-awsctrl

# Install Kubectl
kubectl:
	curl -LO https://storage.googleapis.com/kubernetes-release/release/\$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	sudo mkdir -p /usr/local/bin/
	sudo mv ./kubectl /usr/local/bin/kubectl
	export PATH=\$PATH:/usr/local/bin

# Set Kubectl context
kubectl-context:
	kubectl config set-context kind-awsctrl.io

# Test cluster avilable
kubectl-verify:
	kubectl get nodes -o wide

# Install Kubebuilder
kubebuilder:
	curl -sL https://go.kubebuilder.io/dl/2.2.0/linux/amd64 | tar -xz -C /tmp/
	sudo mv /tmp/kubebuilder_2.2.0_linux_amd64 /usr/local/kubebuilder
	export PATH=\$PATH:/usr/local/kubebuilder/bin

set-env:
	export GO111MODULE=on
	export USE_EXISTING_CLUSTER=true
	export POD_NAMESPACE=default

# find or download generator
# download generator if necessary
generator:
ifeq (, $(shell which generator))
	go install go.awsctrl.io/generator
GENERATOR=$(GOBIN)/generator
else
GENERATOR=$(shell which generator)
endif