.PHONY: test
test:
	go clean -testcache
	cd internal && go test $$(go list ./... | grep -v /mocks) -cover

integration-test:
	docker compose down && docker compose up -d
	cd test_integration && go test .

start-docker:
	docker compose down && docker compose up -d
