docker container stop cryptogo
docker container rm cryptogo
docker build -t cryptogoimage .

docker container run --name cryptogo --rm -it -e PORT=8080 -e INSTANCE_ID="cryptogo-intance" -p 8080:8080 -d cryptogoimage