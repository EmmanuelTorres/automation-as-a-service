* Automation as a Service

A practice project to familiarize myself with microservices.
The first step is to create a system that allows for basic automation services such as pulling a specific project, checking S3 against the Git commit for existing builds, and building then uploading to S3 if S3 does not have it cached.
The end goal is to create a system that allows for more complex automation, such as running a focused regression (gating control) or a complete regression (bug discovery) with minimal input.

This project follows the hexagonal architecture that is prominent in Go. It uses [Ruslan Tsyganok's](https://github.com/ruslantsyganok) [clean architecture example](https://github.com/ruslantsyganok/clean_arcitecture_golang_example) as a template for this pattern.