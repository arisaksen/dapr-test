# KEDA. Kubernetes pub and sub jobs triggerd by autoscaler KEDA.

Offical Dapr/KEDA docs:
https://docs.dapr.io/developing-applications/integrations/autoscale-keda/#create-the-kakfa-topic
Instead of running the subscription app continuously. We will now start a `sub` based on the size of a queue with keda.

## Pre-requisite

* Kubernetes cluster
* Helm
* See pub_sub and `README-k8s.yaml` for basic Dapr on kubernetes.

### Setup

1. Dapr & Rabbitmq pre-requisite (Radis queue not supported by KEDA) & KEDA

* Dapr init:

```sh
dapr init -k --dev --wait
```

* Rabbitmq:

```sh
helm install dapr-dev bitnami/rabbitmq --set auth.tls.enabled=false --wait
```

* KEDA. https://keda.sh/docs/2.15/deploy/

```sh
helm install keda kedacore/keda --namespace keda --create-namespace --wait
```

### Setup

2. Build images

```sh
cd pub
docker build -t pub-image .
cd ../sub
docker build -t sub-image .
```

3. Deploy new `components/pubsub.yaml` Rabbitmq component
   (this was added to the cluster with `dapr init -k --dev --wait`)

Apply the `components/pubsub.yaml`:

```sh
kubectl apply -f components/pubsub.yaml
```

4. Deploy `deployment/subDeployment.yaml`

```sh
kubectl apply -f deploy/subDeployment.yaml
```

```sh
kubectl get secret --namespace default dapr-dev-rabbitmq -o jsonpath="{.data.rabbitmq-password}" | base64 -d
kubectl port-forward svc/dapr-dev-rabbitmq 15672:15672
```

Test Rabbitmq login in web for http://localhost:15672

We should now se the sub app with 1 replica. This will change soon.

Test `deploy/pubJob.yaml` to check sub subscription is working ok.

```sh
kubectl apply -f deploy/pubJob.yaml
```

5. Deploy KEDA autoscaler for Rabbitmq

First get and update password in the `deploy/subScaledObject.yaml`:

```sh
kubectl get secret --namespace default dapr-dev-rabbitmq -o jsonpath="{.data.rabbitmq-password}" | base64 -d
kubectl port-forward svc/dapr-dev-rabbitmq 15672:15672
```

(Remember to not include extra characters like `%` at the end.)

```yaml
metadata:
  host: amqp://user:<password>@dapr-dev-rabitmq:5672`.
```

Now apply updated scaledObject:

```sh
kubectl apply -f deploy/subScaledObject.yaml 
```

This should scale down the deployment to 0.

6. Deploy pub and populate the queue

```sh
kubectl apply -f deploy/pubJob.yaml
```

In `deploy/subScaledObject.yaml` activationValue is set to 11.
This is to show that sub app will not scale up the first time.
But when we run it again the queue will be emptied by the sub app.

```yaml
        activationValue: "11"       # Target value for activating the scaler. Learn more about activation here.(Default: 0, Optional, This value can be a float)
```

You can also see jobs run under events:

```sh
kubectl get events
```

You will not find the logs for this in the cluster. If you have log aggregation tool you should find under pod/job name.

# Cleanup

When using Docker desktop. Just go to settings and reset cluster.

We can also do it manually:

```sh
kubectl delete -f  deploy/.; kubectl delete -f components/.;
dapr uninstall -k --dev
kubectl delete component statestore
helm uninstall dapr-dev bitnami/rabbitmq
```

Remove persistent volumes:

```sh
kubectl get pv
kubectl delete pv <PV-NAME>
kubectl delete pv <PV-NAME>
kubectl delete pv <PV-NAME>
```

# Kubernetes debug host

```sh
kubectl run -i --tty --rm debug --image=alpine --restart=Never -- sh
```
