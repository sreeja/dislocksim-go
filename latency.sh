docker exec -it houston tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it houston tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 109ms
docker exec -it houston tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 251ms

docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:2
docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3


docker exec -it paris tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it paris tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 109ms
docker exec -it paris tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 150ms

docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.41 flowid 1:2
docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.42 flowid 1:2
docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.43 flowid 1:2


docker exec -it singapore tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it singapore tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 251ms
docker exec -it singapore tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 150ms

docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:3

docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.41 flowid 1:2
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.42 flowid 1:2
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.43 flowid 1:2

