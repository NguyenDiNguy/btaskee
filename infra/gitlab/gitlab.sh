echo "......build image...... base" 
docker build . -t base -f ./infra/gitlab/Dockerfile.base

echo "......push image...... base"
docker push base:latest

echo "......clean up build stage image......"
docker image prune -f	