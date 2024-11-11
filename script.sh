docker build -t forum .
docker run -d -p 8080:8080 --name forum-container forum
docker container ps
docker images
docker logs forum-container 
docker system prune -a