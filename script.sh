docker build -t forum .
docker run -d -p 8082:8082 --name forum-container forum
docker system prune -a