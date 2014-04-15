gotunnel
========

creates ssh tunnels from simple yaml configuration

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
