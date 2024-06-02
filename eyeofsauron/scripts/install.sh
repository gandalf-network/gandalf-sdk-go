#!/bin/bash

# Check if go.mod exists
if [ ! -f go.mod ]; then
    echo "go.mod not found. Please run this script in the root of your Go module."
    exit 1
fi

# Add the replace directives
go mod edit -replace github.com/machinebox/graphql=github.com/machinebox/graphql@v0.2.3-0.20181106130121-3a9253180225
go mod edit -replace github.com/Khan/genqlient=github.com/gandalf-network/genqlient@v0.0.0-20240602164016-6dfd436f2b5b

# Tidy up the module
go mod tidy

echo "Replacements added successfully."
