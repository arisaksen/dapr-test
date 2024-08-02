# Dapr setup
README in repo root

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

## User Dapr API
Launch a Dapr sidecar that will listen on port 3500 for a blank application named `myapp`:
```bash
dapr run --app-id myapp --dapr-http-port 3500
dapr list
dapr stop --app-id myapp
```

Since no custom component folder was defined with the above command, Dapr uses the default component definitions created during the dapr init flow.