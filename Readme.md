#set up  postgres in docker container:

### start docker container:

`docker-compose up -d`

### stop docker container: 

`docker-compose -f  docker-compose.yml down --volumes
`

### remove container/volumes

`docker volume prune `
`docker container prune`


### create tables + insert data
`cat ./Customers.sql | docker exec -i setupdb_database_1 psql -U postgres -d postgres`


#Dependencies:

Gorm: ORM library

`go get -u gorm.io/gorm `

PQ: postgres driver 

`go get github.com/lib/pq`


MUX: server + routes

`go get github.com/gorilla/mux
`