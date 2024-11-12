docker build -t forum .
docker run -d -p 8082:8082 --name forum-container forum
docker container ps
docker images
docker logs forum-container 
docker system prune -a