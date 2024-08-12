# Pure dapr cli
To view KEDA it was better to use normal `kubectl`. But we can also use pure dapr cli for this.


1. Build images

You need local container registry to pull images from localhost

```sh
docker run -d -p 6000:5000 --name local-registry registry:2
```

```sh
cd pub
docker build -t localhost:6000/pub-image .; docker push localhost:6000/pub-image:latest
cd ../sub
docker build -t localhost:6000/sub-image .; docker push localhost:6000/sub-image:latest
```

2. Deploy with Dapr

Dapr run:

```sh
dapr run -k -f dapr.yaml
```
