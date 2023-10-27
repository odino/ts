build:
	which gox > /dev/null || go install github.com/mitchellh/gox@latest
	gox -os="darwin" -os="linux" -os="windows" -arch="amd64" -arch="arm64" -osarch="linux/386" -osarch="windows/386" -output="builds/ts_{{.OS}}_{{.Arch}}"