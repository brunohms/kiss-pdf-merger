build:
	go build -o pdf-merger

build-windows:
	GOOS=windows go build -o pdf-merger.exe