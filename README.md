# Dockerfile api2
Moved outside of dir api2 to include dependency from api1 folder

# kubernetes portforward to localhost
kubectl config get-contexts
kubectl config use-context CONTEXT_NAME

kubectl port-forward api1-deploy-6657895cd5-h2vhn --address 0.0.0.0 8080:8080
