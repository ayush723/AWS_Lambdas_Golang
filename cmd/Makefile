build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o=../build/lambdas/main && \
	zip ../build/lambdas/main.zip ../build/lambdas/main &&\
	rm ../build/lambdas/main
