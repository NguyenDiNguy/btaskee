.PHONY: proto
proto:
	buf mod update && buf generate && sh ./inject_tags.sh

.PHONY: setup
setup:
	bash ./infra/setup.sh

.PHONY: base
base:	
	bash ./infra/gitlab/gitlab.sh

.PHONY: build
build:	
	bash ./infra/build.sh $(svc)

.PHONY: deploy
deploy:	
	bash ./infra/deploy.sh $(svc)

.PHONY: test
test:	
	bash ./test/test.sh $(task)