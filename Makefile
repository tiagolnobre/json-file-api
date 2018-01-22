NAME=tiagonobre/json-file-api
VERSION=`git rev-parse --short HEAD`
VERSION=`git rev-parse --short HEAD`

.PHONY: all build build-dev push-version push-dev push

all: build

build-dev:
	@docker build -t $(NAME):dev-$(VERSION) --rm .

build:
	@docker build -t $(NAME):$(VERSION) --no-cache --build-arg app_env=production --rm .

push-version:
	@docker push $(NAME):$(VERSION)

push-dev-version:
	@docker push $(NAME):dev-$(VERSION)

push-dev: build-dev push-dev-version
	@docker tag $(NAME):dev-$(VERSION) $(NAME):latest-dev
	@docker push $(NAME):latest-dev

push: build push-version
	@docker tag $(NAME):$(VERSION) $(NAME):latest
	@docker push $(NAME):latest
