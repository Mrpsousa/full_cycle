## commands - only kafka in docker-compose
 -  kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3
 -  kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe
 -  consumir topics : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste
 -  consumir topics do inicio : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste      --from-beginning
 -  consumir topics do inicio + grupo : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning --group=x
 -  criar msgs : kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

 ## installs
 -  go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka