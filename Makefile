PROJECT := fileguard

OS ?= linux
ARCH ?= amd64

PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	windows/amd64 \
	windows/arm64 \
	darwin/amd64 \
	darwin/arm64

test:
	go test ./...

build:
	@OUTPUT=build/$(OS)/$(ARCH)/$(PROJECT); \
	if [ "$(OS)" = "windows" ]; then OUTPUT="$$OUTPUT.exe"; fi; \
	echo "Building $$OUTPUT ..."; \
	GOOS=$(OS) GOARCH=$(ARCH) go build -o $$OUTPUT cmd/fileguard/main.go

build-all:
	@for platform in $(PLATFORMS); do \
		OS=$${platform%/*}; \
		ARCH=$${platform#*/}; \
		OUTPUT="build/$${OS}/$${ARCH}/$(PROJECT)"; \
		if [ "$$OS" = "windows" ]; then OUTPUT="$$OUTPUT.exe"; fi; \
		echo "Building $$OUTPUT ..."; \
		GOOS=$$OS GOARCH=$$ARCH go build -o $$OUTPUT cmd/fileguard/main.go; \
	done

clean:
	rm -rf build/