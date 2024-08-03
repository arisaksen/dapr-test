# Dapr setup localhost

README in repo root

## Pub/sub

![Screenshot](https://docs.dapr.io/images/pubsub-quickstart/pubsub-diagram.png)
(Checkout: pub app, Order processor: sub app. Image example with http. Here sdk is used)

run with dapr:

```sh
dapr run -f dapr-localhost.yaml
```

no need run init dapr app beforehand like in `state_management`.

### List dapr apps

```sh
dapr list
```

### See queue in Redis

Connect to Redis:

```sh
docker exec -it dapr_redis redis-cli
```

List keys:

```redis
keys *
TYPE orders
```

this should show `orders` and type `stream`

To show the messages for the stream after running the apps with `dapr run -f`

```redis
XRANGE orders - +
```

## Use RabbitMQ

Comment out radis and comment inn rabbitMq in the `localhostComponents/pubsub.yml` folder the run dapr.

Create docker container:

```sh
docker run -it --rm --name dapr_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management 
```

http://localhost:15672
Username: guest, Password: guest

```sh
cd pub_sub
dapr run -f dapr-localhost.yaml
```
