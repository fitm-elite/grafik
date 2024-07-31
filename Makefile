.PHONY: lint
lint: golangci-lint

.PHONY: golangci-lint
golangci-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] golangci-lint: $(mod)" && \
		golangci-lint run --path-prefix $(mod) ./...) &&) true

.PHONY: cover
cover:
	@$(foreach dir,$(COVER_DIRS), ( \
		cd $(dir) && \
		go test -race -coverprofile=cover.out -coverpkg=./... ./... \
		&& go tool cover -html=cover.out -o cover.html) &&) true

.PHONY: test
test:
	go test -race ./...
