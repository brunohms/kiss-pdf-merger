build:
	go build -o builds/pdf-merger

build-windows:
	GOOS=windows go build -o builds/pdf-merger.exe