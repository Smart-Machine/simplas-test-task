### Usage
There is a Makefile for a easier usage. To start the project you can run, from the root of the project:
```bash
make
```

To delete the containers, images, volumes and everything related to docker, and start fresh, you can run:
```bash
make clean
```

----
### API Contracts

##### Insert an object 
```bash
curl --location --request POST 'http://localhost:8001/advertisement/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "_id": "1",
        "categories": {"subcategory": "1407"},
        "title": {
          "ro": "baza", 
          "ru": "baza"
        },
        "type": "zpecial",
        "posted": 1486556302.101039
    }'
```

#### Search an list of object
```bash
curl --location --request GET 'http://localhost:8001/advertisement/?title=Casa'
```

#### Search an object
```bash
curl --location --request GET 'http://localhost:8001/advertisement/32679487'
```

### Update an object
```bash
curl --location --request PUT 'http://localhost:8001/advertisement/32679487' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "_id": "38784587",
        "categories": {"subcategory": "1407"},
        "title": {
          "ro": "Apartament in centrul Chisinaului", 
          "ru": "Центр рышкановки. 3х комнатная"
        },
        "type": "standard",
        "posted": 1486556302.101039
    }'
```

### Delete an object
```bash
curl --location --request DELETE 'http://localhost:8001/advertisement/32679487'
```

-----
### Todo List:
##### Implement a proper healthcheck for the gRPC service.
```
healthcheck:
            test: ["CMD", "bin/grpc_health_probe-linux-amd64", "-addr=localhost:50051"]
            interval: 30s
            timeout: 30s
            retries: 3
```
Source: https://pkg.go.dev/google.golang.org/grpc/health/grpc_health_v1

#### Migrate elasticsearch and gRPC clients to credential usage

#### Unit-testing for all the components

#### Add documentation

#### Dynamically picking up the `service` ip for the gRCP client. Right now the service's client is having a hardcoded docker IP.

#### Expand the proto for the service, so each endpoint returns a specific type in the correlation with its purpose. Although, the very design of the task isn't the best, as the service itself handles the data ingestion and the CRUD. This is a questionable decision, as it creates a critical point, where one single change in the system requires the change of the contracts between the API and the gRPC service's endpoints, the proto types and so on. It would be better to split the responsabilities. For example let the API query ElasticSearch directly. Not mentioning that this way the system is even less efficient. Not querying the data directly adds up a layer of complexity that is redundant. A well made API can query the elasticsearch as well as gRPC does, and it is less work to be done when a new change appear, and it is easier to maintain, not mentioning the faster speed and more isolation between the endpoints.

#### Adding CORS to the API.

#### Adding all the extra middleware to the API, like logging, authorization (JWT tokens ideally).

#### If migrate to Elasticsearch via docker, on a local machine or owned server, generate certs and bind a volume of the CA certs, keys and everything related to the security. Generating them locally on each run adds complexity, downtime, and access permission issues.


----
### Resources:
* ElasticSearch go-client docs: https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8/esapi@v8.13.1#Get

