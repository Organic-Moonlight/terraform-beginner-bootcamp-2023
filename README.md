# Terraform Beginner Bootcamp 2023

## Semantic Versioning

This poject is going to utilize semantic versioning for its tagging. 
[semver.org](https://semver.org/)


The Format:
**MAJOR.MINOR.PATCH**, eg. `1.0.1`

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes

## Install the Terraform CLI

### Considerations with the Terraform CLI changes

The Terraform CLI installation instructions have changed due to gpg keyring change. Needed to refer to the latest install CLI instructions via Terraform Documentation and change the scripting for install. 

[Install Terraform CLI](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)    

### Considerations for Linux Distributions

This project is built against Ubuntu.
Consider checking your Linux Distrubution and change according to your distribution needs. 
[How to check OS version in Linux](https://www.cyberciti.biz/faq/how-to-check-os-version-in-linux-command-line/)

Example of Checking OS Version:

```
$ cat /etc/os-release

PRETTY_NAME="Ubuntu 22.04.4 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.4 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```

#### Refactoring into Bash Scripts 

While facing the Terraform CLI gpg depreciation issues, the bash script steps had more code. So I decided to create a bash script to install the Terraform CLI. 

- This will keep the Gitpod Task File tidy([.gitpod.yml](.gitpod.yml)).
- This will be easier to debug and execute manually Terraform CLI install
- This will allow better portability for other porjects that need to install Terraform CLI

#### Shebang Considerations

A SheBang tells the bash script what program will interpret the script. eg. `#!/bin/bash`

ChatGPT recommended this format for bash: `#!/usr/bin/env bash`

- for portability for different OS distributions
- will search the user's PATH for the bash executable

https://en.wikipedia.org/wiki/Shebang_(Unix)


#### Execution Considerations

When executing the bash script we can use the `./` notation to execute the bash script. 

eg. `./bin/install_terraform_cli`

If we are using a script in .gitpod.yml we need to point the script to a program to interpet it. 

eg. `source ./bin/install_terraform_cli`

#### Linux Considerations

To make our bash scripts executable we need to change Linux permissions for the file to be executable at the user mode. 

```sh
chmod u+x ./bin/install_terraform_cli
```
alternatively: 

```sh
chmod 744 ./bin/install_terraform_cli
```
https://en.wikipedia.org/wiki/Chmod


### Github Lifecycle (before, Init, Command)

Need to be careful when using the Init because it will not rerun if we restart an existing workspace. 

https://www.gitpod.io/docs/configure/workspaces/tasks