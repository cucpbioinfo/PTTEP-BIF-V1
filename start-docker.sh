#!/bin/bash

skip_clean_up=$1

#1 create missing directory
echo "1.) Create upload directory"
cd ./backend
mkdir upload
cd upload
mkdir species-image
mkdir species-video
cd ../..

#2 init config
echo "2.) Init config"
cp ./frontend/.env.example ./frontend/.env.local
cp ./backend/.env.local ./backend/.env
cp ./backend/.env.local ./database/.env
cp ./backend/config.local.yaml ./backend/config.yaml

echo "3.) Stop services"
sudo docker-compose down

if [ "$skip_clean_up" != "true" ]; then
    echo "4.) Cleanup docker"
    sudo docker system prune -f
    sudo docker volume rm $(sudo docker volume ls -q)
else
    echo "4.) Cleanup docker --Skipping"
fi

echo "5.) Start docker"
sudo docker-compose up -d --build

echo "6.) Migrate database"
cd ./database/migrations
go run *.go init
go run *.go
cd ../..

echo "7.) Health Check for API"
curl --write-out "%{http_code}\n" --silent --output /dev/null http://localhost/api/healthz
if [ "$http_response" != "200" ]; then
    echo "Server is healthy."
else
    echo "Server is unhealthy."
fi
