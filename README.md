# An example of go application layering

## About

This project is an attempt to demonstrate layering of a go web app, and to show some rules of communication between layers.

For the purpose of demonstration, this project is a dumb key-value storage app that exposes a REST API. Example:
1. `POST /items` with body `{"key": "foo", "val":  "123"}`,
2. `GET /items/foo` returns `{"key": "foo", "val":  "123"}`.

The layers are as follows, from top to bottom:
1. **Transport** layer. It provides functionality to receive and send requests. The package `rest` belongs to this layer.
2. **Application** layer. It is where domain a.k.a. business logic lives. Package `app` contains a facade interface for this layer and package `appservice` is its implementation.
3. **Storage** layer. It provides means to store and retrieve data from persistent storage. Package `storage` is the facade for this layer, `storagemysql` is its implementation.

### Rules for communication and structure

1. Interlayer code dependencies and calls are only allowed to the layer immediately below;
2. Layers expose interfaces and interlayer calls are only done via interfaces;
3. Models are not re-used between layers: each layer has its own copy of the models it needs;
4. Interfaces live in separate packages from their implementations.

Rules 1 is what makes the app layered. Rule 2 enables loose coupling as makes it possible to change the implementation or mock the layer or its parts. Rule 3 is the model-per-layer rule. It keeps layers loosely coupled because changes to a model in one layer have no effect on its counterpart in another layer. Just as importantly, it makes the models clean by only containing data that its layer needs. For example, in the storage layer, `Item` has a field `Id  uint64` which only makes sense in this layer since it's the primary key in MySQL. Rule 4 allows for modularity.

## Design

The below is a graphical representation of the layering design.

![layers](doc/layers.drawio.png)

## Setup

### Launch the DB
```shell
docker-compose up
```

If the port clashes,
`cp docker-compose.yml docker-compose.local.yml`, change the port mapping and run
`docker-compose -f docker-compose.local.yml up`.

### Run DB migrations

Install the `migrate` tool [golang-migrate/migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation).

```shell
migrate -source file:./migrations -database mysql://root:root@tcp\(localhost:3306\)/app up
```

### Run the app

Run
```shell
make run
```
