
minikube start
minikube dashboard

create test server

eval $(minikube docker-env)

docker build -t test-server .
kubectl -n kube-system apply -f metric-server/metricserver-0.3.7.yaml

kubectl -n kube-system get pods
kubectl apply -f example/

after success deployments:

kubectl port-forward service/client-service 4000:4000 or minikube service client-service





kubectl top pods
kubectl top nodes


watch -n 1 kubectl get hpa

