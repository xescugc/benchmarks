.PHONY: help
help: ## Show this help
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v grep -F | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: floattostring
floattostring:
	go test ./floattostring/ -bench . -benchmem -count 5

.PHONY: inttostring
inttostring:
	go test ./inttostring/ -bench . -benchmem -count 5
