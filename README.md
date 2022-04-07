# simplewebapp
Simple web application on golang returning current time in required timezone

## Little compares
```bash
$ for tag in nethttpget fasthttpget; do GOMAXPROCS=1 timeout 10 go run -tags ${tag} . & sleep 1 && ab -klqdS -c 1000 -n 100000 127.0.0.1:8080/time | grep -F 'Requests per second' ; sleep 9; done
[1] 10895
Requests per second:    36763.46 [#/sec] (mean)
[1]+  Выход 124          GOMAXPROCS=1 timeout 10 go run -tags ${tag} .
[1] 10923
Requests per second:    70527.38 [#/sec] (mean)
[1]+  Выход 124          GOMAXPROCS=1 timeout 10 go run -tags ${tag} .
$ for tag in nethttppost fasthttppost; do GOMAXPROCS=1 timeout 10 go run -tags ${tag} . & sleep 1 && echo 'Asia/Tokyo' > postdata && ab -klqdS -p postdata -c 1000 -n 100000 127.0.0.1:8080/time | grep -F 'Requests per second' ; rm -f postdata; sleep 9; done
[1] 11823
Requests per second:    16064.35 [#/sec] (mean)
[1]+  Выход 124          GOMAXPROCS=1 timeout 10 go run -tags ${tag} .
[1] 12037
Requests per second:    19728.09 [#/sec] (mean)
[1]+  Выход 124          GOMAXPROCS=1 timeout 10 go run -tags ${tag} .
```
