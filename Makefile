TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL := /bin/bash

init:
	@if [ -z "$(shell pipenv --version 1>/dev/null 2>&1)" ]; then \
		pip install --user pipenv; \
	else \
		echo "pipenv is available, good."; \
	fi

install_deps:	init
	@echo "Installing deps..."; \
	pipenv install --dev


tests:	init
	@echo "Running tests..."; \
	pipenv run python3 -B -m unittest discover -s ${TOP_DIR}/tests/ -p '*_test.py'

run:	init
	@echo "Running pkgdiff..."; \
	pipenv run python3 -B -m pkgdiff ${TOP_DIR}/config.yaml