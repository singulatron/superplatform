#!/bin/bash

# Find all interface.go files
find . -name 'interface.go' | while read file; do
    # Get the directory of the file
    dir=$(dirname "${file}")
    # Extract the package name from the source file
    pkg=$(grep -E '^package ' "${file}" | awk '{print $2}')
    # Define the output file for the mock
    output="${dir}/mock_${pkg}.go"
    
    # Generate the mock with the correct package name
    mockgen -source="${file}" -destination="${output}" -package="${pkg}"
    
    echo "Mock generated for ${file} in ${output} with package ${pkg}"
done
