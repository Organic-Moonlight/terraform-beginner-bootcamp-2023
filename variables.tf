variable "user_uuid" {
  description = "UUID for the user"
  type        = string

  # Validation rules
  validation {
    condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "The value for user_uuid must be a valid UUID."
  }
}