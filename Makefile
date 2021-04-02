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

build-ui:
	cd ./ui/ && npm install --save-dev && npm run build
	cp -R ./ui/dist/* ./server/static/

build: clean-build build-ui build-mac build-linux build-windows

build-mac:
	cd ./server/ && env env GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o .$(BUILD_PATH)/mac/$(APP_NAME) .;
	zip $(BUILD_PATH)/mac_$(APP_NAME).zip $(BUILD_PATH)/mac/$(APP_NAME);

build-linux:
	cd ./server/ && env env GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o .$(BUILD_PATH)/linux/$(APP_NAME)_64 .;
	zip $(BUILD_PATH)/linux_$(APP_NAME)_64.zip $(BUILD_PATH)/linux/$(APP_NAME)_64;
	cd ./server/ && env env GOOS=linux GOARCH=386 go build -a -installsuffix cgo -o .$(BUILD_PATH)/linux/$(APP_NAME)_32 .;
	zip $(BUILD_PATH)/linux_$(APP_NAME)_32.zip $(BUILD_PATH)/linux/$(APP_NAME)_32;
	cd ./server/ && env env GOOS=linux GOARCH=arm go build -a -installsuffix cgo -o .$(BUILD_PATH)/linux/$(APP_NAME)_arm32 .;
	zip $(BUILD_PATH)/linux_$(APP_NAME)_arm32.zip $(BUILD_PATH)/linux/$(APP_NAME)_arm32;
	cd ./server/ && env env GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o .$(BUILD_PATH)/linux/$(APP_NAME)_arm64 .;
	zip $(BUILD_PATH)/linux_$(APP_NAME)_arm64.zip $(BUILD_PATH)/linux/$(APP_NAME)_arm64;

build-windows:
	cd ./server/ && env env GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o .$(BUILD_PATH)/windows/$(APP_NAME)_64 .;
	zip $(BUILD_PATH)/windows_$(APP_NAME)_amd64.zip $(BUILD_PATH)/windows/$(APP_NAME)_64;
	cd ./server/ && env env GOOS=windows GOARCH=386 go build -a -installsuffix cgo -o .$(BUILD_PATH)/windows/$(APP_NAME)_32 .;
	zip $(BUILD_PATH)/windows_$(APP_NAME)_amd32.zip $(BUILD_PATH)/windows/$(APP_NAME)_32;
