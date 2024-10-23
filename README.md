```bash
make docker
kubectl apply -f k8s-demo-service.yaml
kubectl apply -f k8s-demo-deployment.yaml
kubectl apply -f k8s-demo-mysql-service.yaml
kubectl apply -f k8s-demo-mysql-deployment.yaml
kubectl apply -f k8s-demo-mysql-pvc.yaml
kubectl apply -f k8s-demo-mysql-pv.yaml
kubectl apply -f k8s-demo-redis-service.yaml
kubectl apply -f k8s-demo-redis-deployment.yaml
kubectl apply -f k8s-demo-ingress.yaml
kubectl get service
kubectl get deployment
kubectl get pod
kubectl get persistentvolumeclaim
kubectl get persistentvolume
kubectl log k8s-demo
kubectl delete deployment k8s-demo-mysql
# 修改 host 文件
sudo vim /etc/hosts
# 查看所有日志
kubectl describe pod k8s-demo-57f9499dd4-bq6v9 
kubectl logs -f -l  app=k8s-demo --all-containers
```
```bash
# ingress-nginx 
brew install helm
helm upgrade --install ingress-nginx ingress-nginx \
--repo https://kubernetes.github.io/ingress-nginx \ 
--namespace ingress-nginx --create-namespace
```