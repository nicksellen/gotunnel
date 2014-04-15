gotunnel
========

creates ssh tunnels using simple yaml configuration

run it like this:
````
gotunnel configuration.yml
````

example yaml configuration:
````
hosts:

  app01:
    dropwizard admin http:
      local: 18081
      remote: 8081
      protocol: http

  db01: 
    postgres:
      local: 15432
      remote: 5432
    redis:
      local: 16379
      remote: 6379

  stats:
    graphite http:
      local: 18000
      remote: 80
      protocol: http

  es01: 
    elasticsearch http:
      local: 19200
      remote: 9200
      protocol: http
````

will give you:
````
host           service                    local      remote    
app01          dropwizard admin http      18081      8081       http://localhost:18081
db01           postgres                   15432      5432      
db01           redis                      16379      6379      
stats          graphite http              18000      80         http://localhost:18000
es01           elasticsearch http         19200      9200       http://localhost:19200
````

press CTRL+C to exit
