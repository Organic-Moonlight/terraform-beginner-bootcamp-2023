variable "user_uuid" {
  description = "UUID for the user"
  type        = string

  # Validation rules
  validation {
    condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "The value for user_uuid must be a valid UUID."
  }
}

variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string

}

variable "index_html_filepath" {
  description = "Path to the index.html file"
  type        = string

  validation {
    condition     = fileexists(var.index_html_filepath)
    error_message = "The provided path for index_html_filepath is invalid or does not exist."
  }
}

variable "error_html_filepath" {
  description = "Path to the error.html file"
  type        = string

  validation {
    condition     = fileexists(var.error_html_filepath)
    error_message = "The provided path for error_html_filepath is invalid or does not exist."
  }
}


variable "content_version" {
  description = "The content version for the resource"
  type        = number

  validation {
    condition     = var.content_version >= 1 && floor(var.content_version) == var.content_version
    error_message = "The content_version value must be a positive integer starting from 1."
  }
}
