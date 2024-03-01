#!/bin/bash

# Update and install Docker
sudo pacman -Syu --noconfirm
sudo pacman -S --noconfirm docker

# Start and enable Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Install Docker Compose
sudo pacman -S --noconfirm docker-compose