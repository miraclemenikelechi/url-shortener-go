variable "db_password" {
  type        = string
  description = "The password for the PostgreSQL database"
  sensitive   = true
}

variable "db_user" {
  type        = string
  description = "The username for the PostgreSQL database"
  default     = "postgres"
}

variable "db_name" {
  type        = string
  description = "The name of the PostgreSQL database"
  default     = "url_shortener"
}

variable "app_port" {
  type        = number
  description = "The port number the application listens on"
  default     = 8649
}
