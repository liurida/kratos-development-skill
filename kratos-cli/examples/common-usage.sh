#!/bin/bash

# Create a new Kratos project
kratos new helloworld

# Navigate into the project directory
cd helloworld

# Add a new protobuf file
kratos proto add api/helloworld/v1/user.proto

# Generate client and server code
kratos proto client api/helloworld/v1/user.proto
kratos proto server api/helloworld/v1/user.proto -t internal/service

# Run the application
kratos run
