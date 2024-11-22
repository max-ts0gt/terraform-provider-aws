#!/bin/bash

OLD_PATH="github.com/hashicorp/terraform-provider-aws"
NEW_PATH="github.com/max-ts0gt/terraform-provider-aws"

# Find all .go files and update the import path
find . -type f -name "*.go" -exec sed -i '' "s|${OLD_PATH}|${NEW_PATH}|g" {} +

# Update go.mod if needed
sed -i '' "s|${OLD_PATH}|${NEW_PATH}|g" go.mod

# Clean up and rebuild
go clean -cache
go mod tidy