docker build -t firemanm/go-lab001:latest .
docker login
// docker tag firemanm/go-lab001:latest
docker push firemanm/go-lab001:latest

git config --global user.name "firemanm-CO9VM"
git config --global user.email firemanm@gmail.com

kubectl apply -f k8s-dp.yaml
kubectl get pod -o wide
kubectl get pod  --show-labels

kubectl apply -f k8s-svc.yaml
kubectl get svc -o wide
kubectl get svc  --show-labels

//describe service
kubectl describe svc lab001v2-svc

//change the name of POD
kubectl exec firemanm-deployment-55f84c6784-l2mqn -it -- curl http://lab001v2-svc 

//healthcheck
kubectl exec firemanm-deployment-55f84c6784-l2mqn -it -- curl http://lab001v2-svc/health 

//kill the app and POD will be restarted
kubectl exec firemanm-deployment-55f84c6784-l2mqn -it -- curl http://lab001v2-svc/stop 

//check RESTARTS field
kubectl get pod -o wide 

//apply config for ingress
minikube addons enable ingress
kubectl apply -f k8s-ingress.yaml
kubectl get ingress -o wide
kubectl describe ingress firemanm-ingress

//apply config for loadbalancer
kubectl apply -f k8s-loadbalancer.yaml
minikube tunnel&

/get LB external IP - change it in command futher
kubectl get svc -A
curl http://10.102.128.79/health

//kill LB
kubectl delete -f k8s-loadbalancer.yaml
Kill <tunnel pid>

//etc
kubectl get pod -A
minikube addons list
minikube addons enable ingress
echo "$(minikube ip) arch.homework" | sudo tee -a /etc/hosts

kubectl get pods -n ingress-nginx

kubectl logs -f deployment/firemanm-deployment --all-containers=true --prefix --timestamps
