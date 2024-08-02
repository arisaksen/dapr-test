# Dapr setup

```bash
dapr init
dapr --version
docker ps
```
On dapr init, the CLI also creates a default components folder that contains several YAML files with definitions for a state store, Pub/sub, and Zipkin. The Dapr sidecar will read these components and use:

The Redis container for state management and messaging.
The Zipkin container for collecting traces.

