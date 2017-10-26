GOCMD=go
GOBUILD=$(GOCMD) build
GOBUILDPROD=$(GOCMD) build -ldflags "-linkmode external -extldflags -static" 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
SODA=buffalo db
BUFFALO=buffalo
APPNAME=gotimefm

build:
	$(GOBUILD) -v -o $APPNAME

buildprod:
	$(GOBUILDPROD) -v -o $APPNAME

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/$APPNAME
	rm -rf $APPNAME

test:
	$(GOTEST) -v ./grifts -race
	$(GOTEST) -v ./models -race
	$(GOTEST) -v ./actions -race

db-up: 
	docker run --name=$APPNAME_db -d -p 5432:5432 -e POSTGRES_DB=$APPNAME_development postgres
	sleep 10
	$(SODA) create 
	$(SODA) migrate up
	docker ps | grep $APPNAME_db

db-down:
	docker stop $APPNAME_db
	docker rm $APPNAME_db


teardown-dev: clean
	$(DOCKERCOMPOSE) down

run-dev: 
	$(BUFFALO) dev

