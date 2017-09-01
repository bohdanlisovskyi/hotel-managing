.PHONY: dependencies
dependencies:
	echo "Installing dependencies"
	glide install

install-helpers:
	echo "Installing Glide"
	go get -u github.com/Masterminds/glide
