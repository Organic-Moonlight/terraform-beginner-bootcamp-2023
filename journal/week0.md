# Terraform Beginner Bootcamp 2023 - Week0

- [Semantic Versioning](#semantic-versioning)
- [Install the Terraform CLI](#install-the-terraform-cli)
   * [Considerations with the Terraform CLI changes](#considerations-with-the-terraform-cli-changes)
   * [Considerations for Linux Distributions](#considerations-for-linux-distributions)
      + [Refactoring into Bash Scripts ](#refactoring-into-bash-scripts)
      + [Shebang Considerations](#shebang-considerations)
      + [Execution Considerations](#execution-considerations)
      + [Linux Considerations](#linux-considerations)
   * [Github Lifecycle (before, Init, Command)](#github-lifecycle-before-init-command)
   * [Working Env Vars](#working-env-vars)
      + [Setting and unsetting Env Vars](#setting-and-unsetting-env-vars)
      + [Printing Vars](#printing-vars)
      + [Scoping of ENv Vars](#scoping-of-env-vars)
      + [Persisiting Env Vars in GItpod](#persisiting-env-vars-in-gitpod)
   * [AWS CLI Installation](#aws-cli-installation)
- [Terraform Basics](#terraform-basics)
   * [Terraform Registry](#terraform-registry)
- [Creating an S3 Bucket with Terraform](#creating-an-s3-bucket-with-terraform)
- [Terraform Cloud](#terraform-cloud)

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


### Working Env Vars

We can list out all Environment Variables (Env Vars) using the `env` command

We can filer specific env vars using grep eg. `env | grep AWS_`

#### Setting and unsetting Env Vars

In the terminal we can set using `export HELLO= 'world'`

In the terminal we can unset using `unset HELLO`

We can set an env var temmporarily when just running a command

```sh
HELLO='world' ./bin/print_message
```
Within a bash script we can set env without writing export eg. 

```sh
#!/usr/bin/env bash

HELLO='world'

echo $HELLO
```

#### Printing Vars

We can print an env var using echo eg. 'echo $HELLO'

#### Scoping of ENv Vars

When you open up new bash terminals in VSCode it will not be aware of env vars that you set in another window. 

If you want Env Vars to persist across all future bash terminals that are open you need to set env vars in your bash profile. eg. `.bash_profile`

#### Persisiting Env Vars in GItpod

We can persist env vars into Gitpod by storing them in Gitpod Secrets Storage.

```
gp env HELLO='world'
```

All future work spaces launched will set the env vars for all bash terminals in those workspaces. 

You can also set env vars in the `.gitpod.yml` but this can only contian non-sensitive env vars

### AWS CLI Installation

AWS CLI is stinalled for this project via the bash script [`./bin/isntall_aws_cli`](./bin/install_aws-cli)

[Getting Started Install (AWS CLI)](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

[AWS CLI Env Vars](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)


We can check if our AWS credentials are configured correctly by running the following command. 

AWS CLI Command:
```sh
aws sts get-caller-identity
```

```json
    "UserId": "AIDA7UC8DQRICINQUHK09",
    "Account": "12345678912",
    "Arn": "arn:aws:iam::98765432198:user/Terraform-Admin"
```
We'll need to generate AWS CLI credits from IAM User to use AWS CLI. 

## Terraform Basics

### Terraform Registry

Terraform sources their providers and modules from the Terraform Registry which is located at [registry.terraform.io](https://registry.terraform.io/)

 - **Providers** is an interface to the APIs that will allow you to create resources in terraform. 
 - **Modules** are a way to make large amount of terrafomr code modular, portable and shareable. 

 [Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

 ### Terraform Console

 We can see a list of all the Terraform commands by simply typing `terrform`

 #### Terraform Init

 At the  start of a new terraform project we will run `terraform init` to download the binaries for the terraform providers that we'll use in this project. 

 #### Terraform Plan
 `Terraform plan`

 This will generate out a changeset, about the state of our infrastructure and what will be changed

 We can output this changeset ie. "plan" to be passed to an apply, but often you can just ignore outputing.

 #### Terraform Apply
 `terraform apply`

 This will run a plan and pass the changeset to be executed by terraform. Apply should prompt yes or no. 

 If you would like to bypass this prompt we can apply the auto approve flag eg. `terraform apply --auto-approve`


 #### Terraform Destory

 `terraform destroy`

 This will destroy resources. 

 You can also us the auto-approve flag to skip the approve prompt eg `terraform apply --auto-approve`

 #### Terraform Lock Files

 `.terraform.lock.hcl` contains the locked versioning for the providers or modules that should be used with this project. 

 The Terraform Lock File **should be commited** to your Version Control System (VSC) eg. Github

 #### Terraform State Files

 `.terraform.tfstate` contains information about the current state of your infrastructure. 

 This file **should not be commited** to your VCS. 

 This file can contain sensentive data. 

 If you lose this file, you lose knowing the state of your infrastucture. 

 `.terraform.tfstate.backup` is the previous state file state. 

 #### Terraform Directory

 `.terraform` directory contains binaries of terraform providers. 


## Creating an S3 Bucket with Terraform

In order to create a bucket in S3 first we would need to add AWS as a provider in `main.tf`

You will grab the provider code that you will need from the terraform registry

[AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest)

Terraform will not allow you to add two providers so will have to modify the code from: 

```
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.44.0"
    }
  }
}

provider "aws" {
  # Configuration options
}
```

to 

```
    aws = {
      source = "hashicorp/aws"
      version = "5.44.0"
    }
  }
}

provider "aws" {
  # Configuration options
}
```

Placing the code underneath the Random Providers that we already installed. 

```
terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.6.0"
    }    
    aws = {
      source = "hashicorp/aws"
      version = "5.44.0"
    }
  }
}

provider "aws" {
  # Configuration options
}
  

  
provider "random" {
  # Configuration options
}
```

When creating an S3 bucket inside of terraform, we used the random provider mentioned above in order to create the S3 bucket name for us. 

When you will need to modify the code to generate a proper bucket name for S3. 

[Bucket Name Rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)

The code will first look like this when grabbing it from the Terraform registry

```
resource "random_string" "random" {
  length           = 16
  special          = true
  override_special = "/@£$"
}
```
Based on the naming rules of the S3 bucket you will want the code to mimic the follow:

```
resource "random_string" "bucket_name" {
  lower = true
  upper =  false
  length   = 16
  special  = false
}
```

## Terraform Cloud

To start, make sure that you have a login for Terraform Cloud. 

What you will want to do is create a project space like so: 

![Terraform Projects](<Terraform Cloud New Project-1.png>)

Then you will need to create a new workspace

![Terraform New Workspace](https://github.com/Organic-Moonlight/terraform-beginner-bootcamp-2023/assets/126846976/4d847489-1c2d-48aa-8722-ce6323417c05)

Since this is a small project you will want to use CLI-driven runs 

![Terraform CLI-Driven Runs](https://github.com/Organic-Moonlight/terraform-beginner-bootcamp-2023/assets/126846976/1bdf11b1-6d0b-4c5a-9615-a056d768c513)

Now we should be following the prompts provided by Terraform: 

You will first want to implement this code in the `main.tf`

```
  cloud {
    organization = "your-organization"

    workspaces {
      name = "Your-Workspace"
    }
  }
```

Once this code is implemented then you will want to run the following comands: 

`terraform login` what this will do is give us access to the terraform cloud in the terminal. 

Once you are inside you will press 'P' so that you will be able to get your terraform token. 

Ctrl+click will lead you to another tab that will give you your token

![Terraform Token](https://github.com/Organic-Moonlight/terraform-beginner-bootcamp-2023/assets/126846976/81a3e935-7952-4f72-a171-55d2988799bd)

Once you have copied the token, go back into the terminal and press 'q' to this will take you out of the terraform cloud and from there you can paste the token. 

Once the token in pasted, you will connected to your workspace. 

Run `terrafrom init` to initialize the workspace. 

Run `terraform apply` to start the first run for this workspace. 

We create a bash script in the bin file so that when we lauch our work space, `terraform login` will be automated. 

[bin/generate_tfrc_credentials](bin/generate_tfrc_credentials)

