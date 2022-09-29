if (docker ps -a| grep redis-eDNA)
then
    docker start redis-eDNA
else
    docker run --name redis-eDNA -p 6379:6379 -d redis
fi