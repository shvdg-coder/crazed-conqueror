# Makefile for project initializing project by generating artifacts and manifests.
IMAGE_NAME = proto-artifacts:latest

build-proto-image:
	podman build -f Dockerfile.proto -t $(IMAGE_NAME) .

go-artifacts: build-proto-image
	podman rm -f proto_tmp 2>/dev/null || true; CID=$$(podman create --name proto_tmp $(IMAGE_NAME)); podman cp $$CID:/artifacts/go/. ./; podman rm $$CID

initialize: go-artifacts

clean-go:
	find ./05-backend -name "*.pb.go" -type f -delete

clean: clean-go

clean-initialize: clean initialize