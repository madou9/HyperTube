#!/bin/bash

# HyperTube - Clean Script (DESTROYS ALL DATA)
echo "HyperTube Clean Script"
echo "WARNING: This will destroy ALL data including database!"
echo ""

read -p "Are you sure you want to continue? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Clean operation cancelled."
    exit 1
fi

echo " Stopping and removing everything..."

# Stop and remove containers, networks, volumes, and images
docker-compose down -v --rmi all

# Remove any dangling images and build cache
echo "Cleaning Docker build cache..."
docker system prune -f

echo "Everything cleaned!"
echo ""
echo "To start fresh: ./start.sh"
