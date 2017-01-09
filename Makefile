all: p7


TAG = 0.0.1
PREFIX = p7
KEY = /tmp/nginx.key
CERT = /tmp/nginx.crt
SECRET = /tmp/secret.json
DBROOTPW = $(shell tr -cd '[:alnum:]' < /dev/urandom | fold -w30 | head -n1 | base64 -)
DBPW = $(shell tr -cd '[:alnum:]' < /dev/urandom | fold -w30 | head -n1 | base64 -)
APIDESIGNPATH = github.com/markusklems/p7/cmd/api/design
IMAGEDESIGNPATH = github.com/markusklems/p7/cmd/image/design
#DBSECRET = $(cat k8s/p7-secret-template.yaml | sed -e "s,{{ password }},$(DBPW),g;" | kubectl create -f --namespace=p7 -)

p7:
	cd cmd/api
	tar -cvf p7.tar js swagger swagger-ui
	-rm p7
	go build -o p7 .

image:
	cd cmd/image
	tar -cvf Dockerfile.tar -C dockerfiles .
	tar -cvf p7.tar js swagger swagger-ui Dockerfile.tar
	go build -o image .

clean:
	rm $(KEY)
	rm $(CERT)

containers: p7
	#docker build -t $(PREFIX)/kube cmd/kube
	#docker tag $(PREFIX)/kube localhost:5000/p7/kube
	docker build -t $(PREFIX)/api:$(TAG) cmd/api
	docker tag $(PREFIX)/api:$(TAG) localhost:5000/$(PREFIX)/api:$(TAG)

push: containers
	#docker push localhost:5000/$(PREFIX)/kube
	docker push localhost:5000/$(PREFIX)/api:$(TAG)

start:
	kubectl create -f k8s/p7.yaml

status:
	kubectl rollout status deployment/p7 -n p7

update:
	kubectl set image deployment/p7-deployment p7=localhost:5000/$(PREFIX)/api:$(TAG) --namespace=p7

keys:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $(KEY) -out $(CERT) -subj "/CN=nginxsvc/O=nginxsvc"


secret: keys
	# go run make_secret.go -crt $(CERT) -key $(KEY) > $(SECRET)
	#https-nginx -crt $(CERT) -key $(KEY) > $(SECRET)
	# kubectl create -f /tmp/secret.json
	# kubectl create configmap nginxconfigmap --from-file=examples/https-nginx/default.conf
	#@echo DBROOTPW is $(DBROOTPW), DBPW is $(DBPW) and db secret $(DBSECRET) was created.

db: secret
	helm install --name p7-db01 --namespace p7 --set mysqlLRootPassword=$(DBROOTPW),mysqlUser=user,mysqlPassword=$(DBPW),mysqlDatabase=p7,persistence.enabled=false \
    stable/mysql

tmp:
	@echo DBROOTPW is $(DBROOTPW) and DBPW is $(DBPW)

goagen_api:
	goagen app     -d $(APIDESIGNPATH) -o cmd/api
	goagen client  -d $(APIDESIGNPATH) -o cmd/api
	goagen swagger -d $(APIDESIGNPATH) -o cmd/api
	goagen js      -d $(APIDESIGNPATH) -o cmd/api
	goagen schema  -d $(APIDESIGNPATH) -o cmd/api
	goagen gen     -d $(APIDESIGNPATH) -o cmd/api --pkg-path=github.com/goadesign/gorma

goagen_image:
	goagen bootstrap -d $(IMAGEDESIGNPATH) -o cmd/image

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
