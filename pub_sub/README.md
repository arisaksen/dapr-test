# Dapr setup

README in repo root

## Pub/sub

![Screenshot](https://docs.dapr.io/images/pubsub-quickstart/pubsub-diagram.png)

run with dapr:

```sh
dapr run -f .
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

### Redis tips - Read the key based on its type

Check the key type:

```
docker exec -it dapr_redis redis-cli
TYPE orders
```

#### If it’s a stream:

`XRANGE orders - +`

#### If it’s a list:

`LRANGE orders 0 -1`

#### If it’s a set:

`SMEMBERS orders`

#### If it’s a sorted set:

`ZRANGE orders 0 -1 WITHSCORES`

#### If it’s a hash:

`HGETALL orders`