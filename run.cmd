docker stop golang001
docker rm golang001
docker build -t golang001 .
docker run -d --name golang001 -p 6060:6060 -v %cd%:/go/src/myfiles golang001
docker exec -it  golang001 /bin/bash
