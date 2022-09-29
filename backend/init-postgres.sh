if (docker ps -a| grep pg-eDNA)
then
    docker start pg-eDNA
else
    docker run --rm --name pg-eDNA -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 -v pgvolume:/var/lib/postgresql/data postgres
fi