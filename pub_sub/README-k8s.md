# Dapr setup

```sh
dapr init -k --dev --wait
```

Show status and components:

```sh
dapr status -k
kubectl get all -o wide
```

Here two dapr components should be added by default:

```sh
  NAMESPACE  NAME        TYPE          VERSION  SCOPES  CREATED              AGE  
  default    pubsub      pubsub.redis  v1               2024-08-02 15:18.20  12m  
  default    statestore  state.redis   v1               2024-08-02 15:18.20  12m  
```

## Pub/sub

1. Build images

```sh
cd pub
docker build -t pub-image .
cd ../sub
docker build -t sub-image .
```

2. Run k8s deployments

```sh
dapr run -k -f dapr-k8s.yaml
```

If `imagepullbackoff` error in cluster. Change `imagePullPolicy` for both apps in the `.dapr/deploy/deployment.yaml`.

```sh
dapr stop -k -f dapr-k8s.yaml
```

imagePullPolicy: Always -> imagePullPolicy: Never

```sh
kubectl apply -f sub/.dapr/deploy/.; kubectl apply -f pub/.dapr/deploy/.
```

```sh
kubectl logs --tail=10 <sub pod name>
```

or just view it in something like `k9s`.

### List dapr apps

```sh
dapr list -k
```

## Cleanup

dapr stop -k -f dapr-k8s.yaml
dapr uninstall -k --dev

Remove persistent volumes:

```sh
kubectl get pv
kubectl delete pv <PV-NAME>
kubectl delete pv <PV-NAME>
kubectl delete pv <PV-NAME>
```
