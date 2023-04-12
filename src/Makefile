.PHONY: test
test:
	go clean -testcache
	cd internal && go test $$(go list ./... | grep -v /mocks) -cover