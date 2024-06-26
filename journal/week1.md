# Terraform Beginner Bootcamp 2023 - Week1

## Root Module Structure

Our root module structure is as follows:

```
PROJECT_ROOT
|
+--- main.tf        # everything else
|
+--- variables.tf   # stores the structure of input variables
|
+--- terraform.tfvars # the data of variables we want to load into our terraform project
|
+--- providers.tf   # defined required providers and their configuration
|
+--- outputs.tf     # stores out outputs
|
+--- README.md      # required for Root modules
```

  [Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform and Input Variables

### Terraform Cloud Variable

In terraform we can set two kinds of variables: 
 - Environment Variables - those you would set in your bash terminal eg. AWS credentials
 - Terraform Variables - those that you would normally set in your tfvars file 

We can set terraform Cloud variables to be sensitive so that they are not shown visibly in the UI. 
     
### Loading Terraform Input Variables

[Terraform Input Variables](https://developer.hashicorp.com/terraform/language/values/variables)

### var-flag
We can use the `-vars` flag to set an input variable or override a variable in the tfvars file eg. `terraform -var user_uuid="my-user_id"`

### var-file flag

 - TODO: research this flag

### terraform.tfvars

This is the default file to load in the terraform variable in blunk

### auto.tfvars

 - TODO: doucment this funcionality for terraform cloud

### Order of terraform variables

 - TODO: document which terraform variables take precedence 

 ## Dealing With Configuration Drift

 ### What happens if we lose our state file? 

 if you lose your state file, you will most likely have to manually tear down all of your cloud infrastructure. 

 You can use Terraform import but it won't work for all cloud resources. You need to check the terraform provider documentation for which resources support import. 

 ### Fix Missing Resources with Terraform Import

 `terraform import aws_s3_bucket.bucket bucket-name`

 [Terraform Import](https://developer.hashicorp.com/terraform/cli/import)

 [AWS S3 Bucket Import](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)

 ### Fix Manual Configuration

  If someone goes and deletes or modifies cloud resources manually through CLickOps. 

  If we run the Terraform Plan, it is with an attempt to put our infrastructure back into the expected state fixing Configuration Drift. 

  ### Fix using Terraform Refresh 

  ```sh
  terraform apply -refresh-only -auto-approve
  ```

  ## Terraform Modules

  ### Terraform Module Structure

  It is recommended to place modules `modules` directory when locally developing modules but you can name it whatever you like. 

  ### Passing input variables

  We can pass input variables to our Module. 
  The module has to declare the terraform variables in its own variables.tf

  ```tf
    module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = vars.user_uuid
  bucket_name = var.bucket_name
}
```

  ### Modules Sources



  Using the source we can import the module from various places eg:
   - locally
   - Github
   - Terraform Registry


  ```tf
  module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
}


[Module Sources](https://developer.hashicorp.com/terraform/language/modules/sources)


## Considerations when using ChatGPT to write Terrafrom

LLMs such as ChatGPT may not be trained on the latest documentation or information about Terraform. 

It may likely produce older examples that could be deprecated. Often affecting providers. 

## Working with Files in Terraform

### File Exist Function

This is a built-in terraform to check the existence of the file. 

```tf
condition = fileexists(var.error_html_filepath)
```

### filemd5

[Terraform filemd5](https://developer.hashicorp.com/terraform/language/functions/filemd5)

### Path Variable

In Terraform there is a special variable called `path` that allows us to reference local paths:
 - path.module = get the path for the current module
 - path.root = get the path for the root module

[Special Path Variable](https://developer.hashicorp.com/terraform/language/expressions/references)

```
resource "aws_s3_object" "index_html" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "index.html"
  source = "${path.root}/public/index.html"
}
```

## Terraform Locals

Locals allow us to define local variables. 
It can be very useful when we need to transform data into another format and have referenced variables. 

[Local Values](https://developer.hashicorp.com/terraform/language/values/locals)

```tf
locals {
    s3_origin_id = "MyS3Origin "
}
```

## Terraform Data Sources

This allows us to source data from cloud resources. 

This is useful when we want to reference cloud resources without importing them. 

```tf
data "aws_caller_identity" "current" {}

output "account_id" {
  value = data.aws_caller_identity.current.account_id
}
```

[Terraform Data Sources](https://developer.hashicorp.com/terraform/language/data-sources)

## Working with JSON

We use the jsonencode to create the json policy inline in the hcl. 


```tf
> jsonencode({"hello"="world"})
{"hello":"world"}
```

[jsonencode](https://developer.hashicorp.com/terraform/language/functions/jsonencode)

### Changing the Lifecycle of Resources

[Meta Arguments Lifecycles](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle)

### Terraform Data

Plain data values such as Local Values and Input Variables don't have any side-effects to plan against and so they aren't valid in replace_triggered_by. You can use terraform_data's behavior of planning an action each time input changes to indirectly use a plain value to trigger replacement.


```tf
  resource "terraform_data" "content_version" {
    input = var.content_version
  }

  lifecycle {
    replace_triggered_by = [terraform_data.content_version.output]
    ignore_changes = [etag]
  }
  ```

[Terraform Data](https://developer.hashicorp.com/terraform/language/resources/terraform-data)


## Provisioners

Provisioners allow you to execute commands on compute instances eg. AWS CLI Command

They are not recommended for use by HashiCorp because Configuration Management tools such as Ansible are a better fit, but the functionality exists. 

### Local-Exec

[Local-Exec](https://developer.hashicorp.com/terraform/language/resources/provisioners/local-exec)

This will execute a command on the machine running the terraform commands eg. plan apply

```tf
resource "aws_instance" "web" {
  # ...

  provisioner "local-exec" {
    command = "echo The server's IP address is ${self.private_ip}"
  }
}
```

### Remote-exec

[Remote-Exec](https://developer.hashicorp.com/terraform/language/resources/provisioners/remote-exec)

this will execute commands on a machine that you target. You must provide credentials such as ssh to get into the machine. 

```tf
resource "aws_instance" "web" {
  # ...

  # Establishes connection to be used by all
  # generic remote provisioners (i.e. file/remote-exec)
  connection {
    type     = "ssh"
    user     = "root"
    password = var.root_password
    host     = self.public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "puppet apply",
      "consul join ${aws_instance.web.private_ip}",
    ]
  }
}
```

## For Each Expressions

For each allows us to enumerate over complex data types

```tf
[for s in var.list : upper(s)]
```
this is mostly useful when you are creating multiples of a cloud resource and you want to reduce the amount of repetitive terraform code. 

[For Each Expressions](https://developer.hashicorp.com/terraform/language/expressions/for)