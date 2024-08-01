# Dapr setup

```bash
dapr init
dapr --version
docker ps
```
On dapr init, the CLI also creates a default components folder that contains several YAML files with definitions for a state store, Pub/sub, and Zipkin. The Dapr sidecar will read these components and use:

The Redis container for state management and messaging.
The Zipkin container for collecting traces.

## User Dapr API
```bash
dapr run --app-id myapp --dapr-http-port 3500
```

### Test with curl
```bash
curl -X POST -H "Content-Type: application/json" -d '[{ "key": "J.R.R Tolkien", "value": "1892"}]' http://localhost:3500/v1.0/state/statestore
curl http://localhost:3500/v1.0/state/statestore/name
curl -v -X DELETE -H "Content-Type: application/json" http://localhost:3500/v1.0/state/statestore/name
```

### See how the state is stored in Redis
```bash
docker exec -it dapr_redis redis-cli
keys *
```

# Docker

```bash
cd api1
docker run -p 8080:8080 -it $(docker build -q .)

cd api2
docker run -p 8081:8081 -it $(docker build -q .)
```

# kubernetes portforward to localhost

```bash
kubectl config get-contexts
kubectl config use-context <context_name>

kubectl port-forward <deployment_name> --address 0.0.0.0 8080:8080
```

# Source Dapr Docs
https://docs.dapr.io/getting-started/get-started-api/