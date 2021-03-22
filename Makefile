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

build-mac: clean-build build-ui
	cd ./server/ && GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o .$(BUILD_PATH)/mac/$(APP_NAME) .;

release-mac: clean-release build-mac
	mkdir -p $(RELEASE_PATH)/mac/$(APP_NAME).app/Contents/MacOS
	mkdir -p $(RELEASE_PATH)/mac/$(APP_NAME).app/Contents/Resources
	cp ./build_assets/mac/Info.plist $(RELEASE_PATH)/mac/$(APP_NAME).app/Contents/
	cp ./build_assets/mac/$(APP_NAME).icns $(RELEASE_PATH)/mac/$(APP_NAME).app/Contents/Resources/
	cp $(BUILD_PATH)/mac/$(APP_NAME) $(RELEASE_PATH)/mac/$(APP_NAME).app/Contents/MacOS/$(APP_NAME)
