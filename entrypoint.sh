#!/bin/sh
set -e

# Pull the llama3 model
ollama pull llama3

# Start the ollama service
ollama run llama3