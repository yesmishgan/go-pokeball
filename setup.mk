# project info
USERNAME?=yesmishgan
PROJECT_NAME?=go-pokeball

OLD_PROJECT_NAME?=yesmishgan/go-pokeball
NEW_PROJECT_NAME?=$(USERNAME)/$(PROJECT_NAME)

export REMOTE_ORIGIN_URL := $(shell git config --get remote.origin.url)

setup-project: clean-up-git
	find . -not -path '*/.*' -exec sed -i '' "s%$(OLD_PROJECT_NAME)%$(NEW_PROJECT_NAME)%g" {} \;
