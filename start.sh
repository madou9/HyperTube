#!/bin/bash

# HyperTube - Start Script
echo "Starting HyperTube application..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Docker is not running. Please start Docker Desktop first."
    exit 1
fi

# Build and start all services
echo "Building and starting all services..."
docker-compose up --build

# Wait a moment for services to start
echo "Waiting for services to initialize..."
sleep 5

# Check service status
echo " Service Status:"
docker-compose ps

# Show logs if any service is not healthy
echo ""
echo "🔍 Checking for any issues..."
if docker-compose ps | grep -q "Exit\|Restarting"; then
    echo " Some services might have issues. Showing recent logs:"
    docker-compose logs --tail=20
else
    echo " All services are running!"
    echo ""
    echo "🌐 Your application is available at:"
    echo "   Frontend: http://localhost:3000"
    echo "   Backend:  http://localhost:8080"
    echo "   Database: localhost:5432"
    echo "   Redis:    localhost:6379"
fi

echo ""
echo "💡 Useful commands:"
echo "   View logs:    docker-compose logs -f"
echo "   Stop all:     ./stop.sh"
echo "   Restart:      ./restart.sh"
echo "   Clean all:    ./clean.sh"
