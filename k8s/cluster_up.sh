# debug
kubectl --v=8 version

kubectl create ns p7

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /tmp/tls.key -out /tmp/tls.crt -subj "/CN=jenkins/O=jenkins"
kubectl create secret tls tls-secret --cert=path/to/tls.cert --key=path/to/tls.key

### database
helm install --name database stable/mysql
# get root password - helm list # get release
printf $(printf '\%o' `kubectl get secret [YOUR_RELEASE_NAME]-mysql -o jsonpath="{.data.mysql-root-password[*]}"`)

