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