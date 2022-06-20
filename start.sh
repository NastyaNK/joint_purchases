sudo chmod 666 /var/run/docker.sock
cd docker
docker-compose up -d
export PGPASSWORD=postgres
export POSTGRES_PASSWORD=postgres
# если возникла ошибка нужно увеличить время ожидания
sleep 10
docker exec -it postgres psql -U postgres -c "CREATE DATABASE mvp ENCODING 'UTF8' TEMPLATE template0 LC_COLLATE 'C' LC_CTYPE 'C';"

cd ../
./migrago1 -c migration/config.yaml init
./migrago1 -c migration/config.yaml up
go run main.go
# удалить контейнер
docker rm -f postgres