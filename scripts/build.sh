# Command to build the application for macOS, windows and linux
wails build -platform darwin/arm64,linux/amd64,windows/amd64

# Command to zip the application for macOS and windows
zip Dockernetes_MACOS_arm64.zip build/bin/Dockernetes.app
zip Dockernetes_MACOS_Windows_AMD64.zip build/bin/Dockernetes-amd64.exe