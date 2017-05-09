all: p7


TAG = 0.0.3
PREFIX = p7
ID = k8s-$(PREFIX)
REGISTRY=gcr.io
KEY = /tmp/nginx.key
CERT = /tmp/nginx.crt
SECRET = /tmp/secret.json
DBROOTPW = $(shell tr -cd '[:alnum:]' < /dev/urandom | fold -w30 | head -n1 | base64 -)
DBPW = $(shell tr -cd '[:alnum:]' < /dev/urandom | fold -w30 | head -n1 | base64 -)
APIDESIGNPATH = github.com/markusklems/p7/cmd/api/design
IMAGEDESIGNPATH = github.com/markusklems/p7/cmd/image/design
KUBEDESIGNPATH = github.com/markusklems/p7/cmd/kube/design
#DBSECRET = $(cat k8s/p7-secret-template.yaml | sed -e "s,{{ password }},$(DBPW),g;" | kubectl create -f --namespace=p7 -)

api:
	cd cmd/api
	tar -cvf api.tar public
	-rm api
	export GO_EXTLINK_ENABLED=0
	export CGO_ENABLED=0
	go build --ldflags '-extldflags "-static"' -o api .

image:
	cd cmd/image
	tar -cvf Dockerfile.tar -C dockerfiles .
	tar -cvf image.tar public Dockerfile.tar
	export GO_EXTLINK_ENABLED=0
	export CGO_ENABLED=0
	go build --ldflags '-extldflags "-static"' -o image .

kube:
	cd cmd/kube
	tar -cvf kube.tar public
	export GO_EXTLINK_ENABLED=0
	export CGO_ENABLED=0
	go build --ldflags '-extldflags "-static"' -o kube .

clean:
	rm $(KEY)
	rm $(CERT)

containers: api
	#docker build -t $(PREFIX)/image:$(TAG) cmd/image
	#docker tag $(PREFIX)/image:$(TAG) $(REGISTRY)/$(ID)/$(PREFIX)/image:$(TAG)
	#docker build -t $(PREFIX)/kube cmd/kube
	#docker tag $(PREFIX)/kube localhost:5000/p7/kube
	#docker build -t $(PREFIX)/api:$(TAG) cmd/api
	#docker tag $(PREFIX)/api:$(TAG) localhost:5000/$(PREFIX)/api:$(TAG)
	docker build -t $(PREFIX)/api:$(TAG) cmd/api
	docker tag $(PREFIX)/api:$(TAG) $(REGISTRY)/$(ID)/api:$(TAG)

push: containers
	#docker push localhost:5000/$(PREFIX)/kube
	gcloud docker -- push $(REGISTRY)/$(ID)/api:$(TAG)

delete:
	gcloud beta container images delete $(REGISTRY)/$(ID)/api:$(TAG)

start:
	kubectl create -f k8s/p7.yaml

status:
	kubectl rollout status deployment/p7 -n p7

update:
	kubectl set image deployment/p7-deployment p7=$(ID)/$(PREFIX)/api:$(TAG) --namespace=p7

keys:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $(KEY) -out $(CERT) -subj "/CN=nginxsvc/O=nginxsvc"

secret: keys
	# go run make_secret.go -crt $(CERT) -key $(KEY) > $(SECRET)
	#https-nginx -crt $(CERT) -key $(KEY) > $(SECRET)
	# kubectl create -f /tmp/secret.json
	# kubectl create configmap nginxconfigmap --from-file=examples/https-nginx/default.conf
	#@echo DBROOTPW is $(DBROOTPW), DBPW is $(DBPW) and db secret $(DBSECRET) was created.
	#kubectl --namespace=p7 create secret tls tmrtn.de --cert=fullchain1.pem --key=privkey1.pem
	#kubectl --namespace=p7 create secret generic basic-auth --from-file=auth

db:
	helm install --name p7-db01 --namespace p7 --set mysqlLRootPassword=$(DBROOTPW),mysqlUser=user,mysqlPassword=$(DBPW),mysqlDatabase=p7 stable/mysql

get-db-secret:
	kubectl get secret --namespace p7 p7-db01-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo
	kubectl get secret --namespace p7 p7-db01-mysql -o jsonpath="{.data.mysql-password}" | base64 --decode; echo

db-connect:
	kubectl run -i --tty ubuntu --image=ubuntu:16.04 --restart=Never -- bash -il
	apt-get update && apt-get install mysql-client -y
	mysql -h p7-db01-mysql -p

tmp:
	@echo DBROOTPW is $(DBROOTPW) and DBPW is $(DBPW)

goagen_api:
	goagen app     -d $(APIDESIGNPATH) -o cmd/api
	goagen client  -d $(APIDESIGNPATH) -o cmd/api
	goagen swagger -d $(APIDESIGNPATH) -o cmd/api/public
	goagen js      -d $(APIDESIGNPATH) -o cmd/api/public
	goagen schema  -d $(APIDESIGNPATH) -o cmd/api/public
	goagen gen     -d $(APIDESIGNPATH) -o cmd/api --pkg-path=github.com/goadesign/gorma

goagen_image:
	goagen app        -d $(IMAGEDESIGNPATH) -o cmd/image
	goagen client     -d $(IMAGEDESIGNPATH) -o cmd/image
	goagen controller -d $(IMAGEDESIGNPATH) -o cmd/image/controllers
	goagen js         -d $(IMAGEDESIGNPATH) -o cmd/image/public
	goagen schema     -d $(IMAGEDESIGNPATH) -o cmd/image/public
	goagen swagger    -d $(IMAGEDESIGNPATH) -o cmd/image/public

goagen_kube:
	goagen app        -d $(KUBEDESIGNPATH) -o cmd/kube
	goagen client     -d $(KUBEDESIGNPATH) -o cmd/kube
	goagen controller -d $(KUBEDESIGNPATH) -o cmd/kube/controllers
	goagen js         -d $(KUBEDESIGNPATH) -o cmd/kube/public
	goagen schema     -d $(KUBEDESIGNPATH) -o cmd/kube/public
	goagen swagger    -d $(KUBEDESIGNPATH) -o cmd/kube/public

run:
	#docker run --rm -p 8888:8888 -ti $(PREFIX)/api:$(TAG)
	docker run -v /run/docker.sock:/var/run/docker.sock --rm -p 8890:8890 -ti $(PREFIX)/image:$(TAG)

setup:
	glide install
	ln -s vendor/ src/

define tests =
	for pkg in $(go list ./... | grep -v '/vendor/'); do
	  go test ${pkg} -v;
	done;
endef
test: ; @ $(value tests)

.ONESHELL:
