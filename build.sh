POD=$(kubectl get pods --namespace kube-system -l k8s-app=kube-registry \
            -o template --template '{{range .items}}{{.metadata.name}} {{.status.phase}}{{"\n"}}{{end}}' \
            | grep Running | head -1 | cut -f1 -d' ')
kubectl port-forward --namespace kube-system $POD 5000:5000 &
ssh -L 5000:127.0.0.1:5000 server
