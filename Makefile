release:
	git tag $(shell cat VERSION)
	git push origin $(shell cat VERSION)
	goreleaser --clean
