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
kubectl apply -f k8s-ingress.yaml
kubectl get ingress -o wide

kubectl describe ingress firemanm-ingress


//etc
kubectl get svc -A
kubectl get pod -A

minikube tunnel&


