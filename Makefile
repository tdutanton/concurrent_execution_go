SHELL := /bin/sh

CHECKTASK = $(shell which task)

.PHONY: all ensure-task

all: ensure-task
	@echo "Hello! Now you can run the program with Task!"
	@echo "Just enter \"task\" into console."

ensure-task:
ifeq (, $(CHECKTASK))
	@echo "Task is not installed, installing..."
	@go install github.com/go-task/task/v3/cmd/task@latest
	@echo "Task installed"
endif