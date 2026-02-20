TEST_IMAGE_NAME=yt-auto-archiver-tester
TEST_DOCKERFILE=test.Dockerfile



test:
	docker build -t ${TEST_IMAGE_NAME} -f ${TEST_DOCKERFILE} .
	docker run -v .:/app ${TEST_IMAGE_NAME}  \
		go test -v ./... -coverprofile=coverage.out -covermode=atomic