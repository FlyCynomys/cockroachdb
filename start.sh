./cockroach start --insecure --background --store=node1 --host=localhost --port=26257 --http-port=8081
sleep 2s
./cockroach start --insecure --background --store=node2 --host=localhost --port=26258 --http-port=8082 --join=localhost:26257 
sleep 2s
./cockroach start --insecure --background --store=node3 --host=localhost --port=26259 --http-port=8083 --join=localhost:26257 
sleep 2s
./cockroach start --insecure --background --store=node4 --host=localhost --port=26260 --http-port=8084 --join=localhost:26257 
sleep 2s
./cockroach start --insecure --background --store=node5 --host=localhost --port=26261 --http-port=8085 --join=localhost:26257 

echo "start cluater over"

#./cockroach start --insecure --background --store=node6 --host=localhost --port=26262 --http-port=8086 --join=localhost:26257 
#./cockroach start --insecure --background --store=node7 --host=localhost --port=26263 --http-port=8087 --join=localhost:26257 
#echo 'num_replicas: 5' | cockroach zone set .default --insecure -f -
#cockroach quit --insecure