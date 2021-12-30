TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL := /bin/bash

init:
	if [ ! -d "${TOP_DIR}/.venv/" ]; then \
		virtualenv --version >/dev/null 2>&1 || pip install virtualenv; \
		python3 -B -m virtualenv -p python3 ${TOP_DIR}/.venv/; \
	fi && \
	source ${TOP_DIR}/.venv/bin/activate && \
	pip install -r ${TOP_DIR}/requirements.txt -r ${TOP_DIR}/requirements-dev.txt

install_deps:	init
	@echo "Installing deps..."; \
	source ${TOP_DIR}/.venv/bin/activate && \
	pip install -r ${TOP_DIR}/requirements.txt -r ${TOP_DIR}/requirements-dev.txt


tests:	init
	@echo "Running tests..."; \
	source ${TOP_DIR}/.venv/bin/activate && \
	python3 -B -m unittest discover -s ${TOP_DIR}/tests/ -p '*_test.py'

run:	init
	@echo "Running installed_packages_diff..."; \
	source ${TOP_DIR}/.venv/bin/activate && \
	python3 -B -m installed_packages_diff ${TOP_DIR}/config.yaml

dist:   clean install_deps tests
	@echo "Bulding dist..."; \
	source ${TOP_DIR}/.venv/bin/activate && \
	python3 -B ${TOP_DIR}/setup.py sdist bdist_wheel

clean:
	rm -rf ${TOP_DIR}/dist/ ${TOP_DIR}/build/ *.egg-info/

upload: dist
	@echo "Uploading dist..."; \
	source ${TOP_DIR}/.venv/bin/activate && \
	twine upload dist/*
