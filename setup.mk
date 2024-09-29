# project info
GITHUB_USERNAME?=yesmishgan
PROJECT_NAME?=go-project-template

OLD_PROJECT_NAME?=yesmishgan/go-project-template
NEW_PROJECT_NAME?=$(GITHUB_USERNAME)/$(PROJECT_NAME)

export REMOTE_ORIGIN_URL := $(shell git config --get remote.origin.url)

setup-project: clean-up-git
	find . -not -path '*/.*' -exec sed -i '' "s%$(OLD_PROJECT_NAME)%$(NEW_PROJECT_NAME)%g" {} \;

clean-up-git:
	@if [ $(REMOTE_ORIGIN_URL) == "git@github.com:yesmishgan/go-project-template.git" ]; then\
		rm -rf .git;\
		git init;\
	fi
