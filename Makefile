SHELL := /bin/bash
.DEFAULT_GOAL := help
.PHONY : help install_cli configure_cli
-include .env/main.env


exercism_version=3.2.0
cli_url="https://github.com/exercism/cli/releases/download/v$(exercism_version)/exercism-$(exercism_version)-linux-x86_64.tar.gz"
source_code_url="https://raw.githubusercontent.com/exercism/cli/v$(exercism_version)/shell/exercism_completion.bash"
bin_dir=$(CURDIR)/.bin
tmp_dir=$(CURDIR)/.tmp
conf_dir=$(CURDIR)/.conf
bash_profiles_dir=$(conf_dir)/bash
bashrc_path="${HOME}/.bashrc"
source_command=source $(bash_profiles_dir)/*


help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

all: install conf autocomplete

install:		## Install exercism cli
	mkdir -p $(tmp_dir)
	curl -L $(cli_url) -o $(tmp_dir)/exercism.tar.gz
	tar -xvzf $(tmp_dir)/exercism.tar.gz -C $(tmp_dir)

	mkdir -p $(bin_dir)
	mv $(tmp_dir)/exercism $(bin_dir)
	chmod +x $(bin_dir)/exercism

	$(bin_dir)/exercism --help

	rm -rf $(tmp_dir)

autocomplete:		## Add exercism cli autocomplete
	mkdir -p $(bash_profiles_dir)
	curl -L $(source_code_url) -o $(bash_profiles_dir)/exercism_completion.bash

	grep "${source_command}" $(bashrc_path) || echo "${source_command}" >> $(bashrc_path)

conf:		## Configure exercism cli
	exercism configure --token=$(EXERCISM_TOKEN)