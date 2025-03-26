docker rm -f $(docker ps -aq)

docker volume prune -f

docker rmi -f $(docker images -q)

docker system prune -a --volumes -f

docker-compose up --build