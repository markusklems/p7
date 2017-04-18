# debug
kubectl --v=8 version

kubectl create -f namespace-p7.json

### database
helm install --name p7 --namespace p7 stable/mysql
# edit pvc yaml and remove storage class annotation

# get root password
printf $(printf '\%o' `kubectl get secret --namespace p7 p7-mysql -o jsonpath="{.data.mysql-root-password[*]}"`);echo

# debug database
docker exec -ti mysql mysql -u root --password=rootpwd

### p7
kubectl -n p7 port-forward p7-3170047332-qtsn6 8888

