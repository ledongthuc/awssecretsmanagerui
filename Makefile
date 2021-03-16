BUILD_PATH=./build
RELEASE_PATH=./release
APP_NAME="aws-secrets-manager-ui"

clean: clean-build clean-release

clean-build:
	rm -rf $(BUILD_PATH);
	mkdir $(BUILD_PATH);

clean-release:
	rm -rf $(RELEASE_PATH);
	mkdir $(RELEASE_PATH);

build: clean-build build-ui build-mac

build-ui:
	cd ./ui/ && npm install --save-dev && npm run build
	cp -R ./ui/dist/* ./server/static/

build-docker:
	docker build -f Dockerfile -t awssecretsmanagerui:latest .;

build-mac:
	cd ./server/ && GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ../build/$(APP_NAME) .;
