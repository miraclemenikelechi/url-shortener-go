terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "4.4.0"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

resource "docker_image" "app" {
  name = "url-shortener-go-application-image"
  build {
    context    = abspath("${path.module}/..")
    dockerfile = abspath("${path.module}/../Dockerfile")
  }
  keep_locally = false
}


resource "docker_container" "app" {
  image = docker_image.app.image_id
  name  = "url-shortener-go-application-container"

  env=[
    "PORT=${var.app_port}",
    "DATABASE_URL=postgres://${var.db_user}:${var.db_password}@postgres-container:5432/${var.db_name}?sslmode=disable"
  ]

  ports {
    internal = var.app_port
    external = var.app_port
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [
    docker_container.postgres,
  ]
}


resource "docker_image" "postgres" {
  name         = "postgres:18-alpine"
  keep_locally = false
}

resource "docker_container" "postgres" {
  image = docker_image.postgres.image_id
  name  = "postgres-container"

  env=[
    "POSTGRES_USER=${var.db_user}",
    "POSTGRES_PASSWORD=${var.db_password}",
    "POSTGRES_DB=${var.db_name}",
  ]

  networks_advanced {
    name = docker_network.app_network.name
  }
}


resource "docker_network" "app_network" {
  name = "url-shortener-go-application-network"
}
