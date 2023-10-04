# simplewebapp
Simple web application on golang returning current time in required timezone. I built this repo to compare fasthttp library with net/http

## Compares
### GET raw
net/http
```bash
# export GOMAXPROCS=2 ; taskset -c 0-1 go run -tags nethttprawget .
$ curl -s 127.0.0.1:8080/time?timezone=CET -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m GET 127.0.0.1:8080/time?timezone=CET | grep -F 'Reqs/sec' | awk '{print $2}'
67654.53

# export GOMAXPROCS=4 ; taskset -c 0-3 go run -tags nethttprawget .
$ curl -s 127.0.0.1:8080/time?timezone=CET -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m GET 127.0.0.1:8080/time?timezone=CET | grep -F 'Reqs/sec' | awk '{print $2}'
102465.61

# export GOMAXPROCS=6 ; taskset -c 0-5 go run -tags nethttprawget .
$ curl -s 127.0.0.1:8080/time?timezone=CET -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m GET 127.0.0.1:8080/time?timezone=CET | grep -F 'Reqs/sec' | awk '{print $2}'
125686.60
```
fasthttp
```bash
# export GOMAXPROCS=2 ; taskset -c 0-1 go run -tags fasthttprawget .
$ curl -s 127.0.0.1:8080/time?timezone=CET -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m GET 127.0.0.1:8080/time?timezone=CET | grep -F 'Reqs/sec' | awk '{print $2}'
151376.31

# export GOMAXPROCS=4 ; taskset -c 0-3 go run -tags fasthttprawget .
$ curl -s 127.0.0.1:8080/time?timezone=CET -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m GET 127.0.0.1:8080/time?timezone=CET | grep -F 'Reqs/sec' | awk '{print $2}'
263952.38

# export GOMAXPROCS=6 ; taskset -c 0-3 go run -tags fasthttprawget .
$ curl -s 127.0.0.1:8080/time?timezone=CET -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m GET 127.0.0.1:8080/time?timezone=CET | grep -F 'Reqs/sec' | awk '{print $2}'
320146.88
```
### POST raw
net/http
```bash
# export GOMAXPROCS=2 ; taskset -c 0-1 go run -tags nethttprawpost .
$ curl -s -X POST --data 'CET' 127.0.0.1:8080/time -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m POST -b 'CET' 127.0.0.1:8080/time | grep -F 'Reqs/sec' | awk '{print $2}'
64497.85

# export GOMAXPROCS=4 ; taskset -c 0-3 go run -tags nethttprawpost .
$ curl -s -X POST --data 'CET' 127.0.0.1:8080/time -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m POST -b 'CET' 127.0.0.1:8080/time | grep -F 'Reqs/sec' | awk '{print $2}'
104923.02

# export GOMAXPROCS=6 ; taskset -c 0-5 go run -tags nethttprawpost .
$ curl -s -X POST --data 'CET' 127.0.0.1:8080/time -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m POST -b 'CET' 127.0.0.1:8080/time | grep -F 'Reqs/sec' | awk '{print $2}'
126524.78
```
fasthttp
```bash
# export GOMAXPROCS=2 ; taskset -c 0-1 go run -tags fasthttprawpost .
$ curl -s -X POST --data 'CET' 127.0.0.1:8080/time -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m POST -b 'CET' 127.0.0.1:8080/time | grep -F 'Reqs/sec' | awk '{print $2}'
95958.15

# export GOMAXPROCS=4 ; taskset -c 0-3 go run -tags fasthttprawpost .
$ curl -s -X POST --data 'CET' 127.0.0.1:8080/time -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m POST -b 'CET' 127.0.0.1:8080/time | grep -F 'Reqs/sec' | awk '{print $2}'
132285.04

# export GOMAXPROCS=6 ; taskset -c 0-5 go run -tags fasthttprawpost .
$ curl -s -X POST --data 'CET' 127.0.0.1:8080/time -o /dev/null ; bombardier -lk --fasthttp -c 1000 -d 50s -m POST -b 'CET' 127.0.0.1:8080/time | grep -F 'Reqs/sec' | awk '{print $2}'
163599.96
```

### GET with minifier and compressor
TODO

### POST with minifier and compressor
TODO
