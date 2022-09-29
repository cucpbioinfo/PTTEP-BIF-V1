if (docker ps -a| grep pgadmin)
then
    docker start pgadmin
else
    docker run --name pgadmin -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=pgadmin' -d dpage/pgadmin4
fi