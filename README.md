# Dockerfile api2
Moved outside of dir api2 to include dependency from api1 folder

# kubernetes portforward to localhost
```bash
kubectl config get-contexts
kubectl config use-context <context_name>

kubectl port-forward <deployment_name> --address 0.0.0.0 8080:8080
```
