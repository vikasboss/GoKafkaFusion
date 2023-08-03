Here we'll learn the basic kafka structure.
We'll create simple kafka producer and consumer using Go.
Some important commands:

bash
# to start the project
sudo service docker start
docker-compose up -d

# execute docker containers
docker exec -it go-kafka bash
docker exec -it kafka-container bash
kafka
run these commands inside "kafka-container" container.

# create topic
kafka-topics --create --bootstrap-server localhost:9092 --topic first --partitions 3

# console consumer (for early tests) (to see all messages add --from-beginning)
kafka-console-consumer --bootstrap-server localhost:9092 --topic first --group --first-consumers

# console producer (for early tests)
kafka-console-producer --bootstrap-server localhost:9092 --topic first


go
run these commands inside "go-kafka" container.

# run consumer
go run cmd/consumer/main.go
# run producer
go run cmd/producer/main.go
