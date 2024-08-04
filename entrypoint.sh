#!/bin/sh
set -e

# Debug information
echo "Starting ollama entrypoint script"

# Start the ollama service
echo "Starting ollama service..."
ollama run llama3 &

# Wait for the ollama service to be ready
while ! curl -s http://localhost:11434/api; do
  echo "Waiting for ollama service to be ready..."
  sleep 5
done

# Pull the llama3 model
echo "Pulling llama3 model..."
ollama pull llama3

echo "Ollama service is ready"
