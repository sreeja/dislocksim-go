docker exec -it houston tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it houston tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it houston tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 125ms

docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:2
docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.61 flowid 1:2
docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.62 flowid 1:2
docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.63 flowid 1:2

docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.72 flowid 1:2
docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.73 flowid 1:2

docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.82 flowid 1:2
docker exec -it houston tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it paris tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it paris tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it paris tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it paris tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it singapore tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it singapore tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 125ms
docker exec -it singapore tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:3

docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.61 flowid 1:3
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.62 flowid 1:3
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.63 flowid 1:3

docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.72 flowid 1:3
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.73 flowid 1:3

docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it singapore tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.82 flowid 1:3




docker exec -it etcd0-clust tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd0-clust tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd0-clust tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 125ms

docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:2
docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.61 flowid 1:2
docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.62 flowid 1:2
docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.63 flowid 1:2

docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.72 flowid 1:2
docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.73 flowid 1:2

docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.82 flowid 1:2
docker exec -it etcd0-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd0-dist tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd0-dist tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd0-dist tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 125ms

docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:2
docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.61 flowid 1:2
docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.62 flowid 1:2
docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.63 flowid 1:2

docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.72 flowid 1:2
docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.73 flowid 1:2

docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.82 flowid 1:2
docker exec -it etcd0-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3



docker exec -it etcd0-cent tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd0-cent tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd0-cent tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd0-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd0-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd0-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it etcd0-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd0-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd1-cent tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd1-cent tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd1-cent tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd1-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd1-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd1-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it etcd1-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd1-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd2-cent tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd2-cent tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd2-cent tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd2-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd2-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd2-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it etcd2-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd2-cent tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd1-clust tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd1-clust tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd1-clust tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd1-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd1-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd1-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it etcd1-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd1-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd2-clust tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd2-clust tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd2-clust tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd2-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd2-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd2-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it etcd2-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd2-clust tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd1-dist tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd1-dist tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 55ms
docker exec -it etcd1-dist tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd1-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd1-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.53 flowid 1:3

docker exec -it etcd1-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2

docker exec -it etcd1-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd1-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.83 flowid 1:3


docker exec -it etcd2-dist tc qdisc add dev eth0 root handle 1: prio bands 16 priomap 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

docker exec -it etcd2-dist tc qdisc add dev eth0 parent 1:2 handle 2: netem delay 125ms
docker exec -it etcd2-dist tc qdisc add dev eth0 parent 1:3 handle 3: netem delay 75ms

docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.51 flowid 1:2
docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.52 flowid 1:3

docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.61 flowid 1:3
docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.62 flowid 1:3
docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.63 flowid 1:3

docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.71 flowid 1:2
docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.72 flowid 1:3
docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.73 flowid 1:3

docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.81 flowid 1:2
docker exec -it etcd2-dist tc filter add dev eth0 parent 1:0 protocol ip u32 match ip dst 172.21.0.82 flowid 1:3



