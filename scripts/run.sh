# Description: Run the docker containers

# Start the docker containers
docker-compose up -d

# Wait for the containers to start
sleep 2

# Propmpt the user to show the running containers
read -p "Do you want to show the running containers? (y/N): " showContainers

# Show the running containers
if [ "$showContainers" = "y" ]; then
    docker-compose ps
fi

# Propmpt the user to show the logs of the running containers
read -p "Do you want to show the logs of the running containers? (y/N): " showLogs

# Show the logs of the running containers
if [ "$showLogs" = "y" ]; then
    docker-compose logs -f
fi
