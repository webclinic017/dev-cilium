template := $(shell basename $(shell pwd))

push:
	tar -cvh --exclude .terraform . | coder templates push -y --create -d - "$(template)"
