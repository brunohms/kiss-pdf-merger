build:
	go build -o builds/pdf-merger
	GOOS=windows go build -o builds/pdf-merger.exe
