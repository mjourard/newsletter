.PHONY: build clean deploy gomodgen test initdb initdomain inittest remove removedb removetest

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/addemail addemail/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/getemails getemails/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/addarchivedarticle addarchivedarticle/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --config sls_app.yaml --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

test:
	go test

initdb:
	sls deploy --config sls_persistent_resources.yaml --region us-east-1

initdomain:
	sls create_domain --config sls_app.yaml --region us-east-1

inittest:
	sls deploy --config sls_test.yaml --region us-east-1

remove:
	sls remove --config sls_app.yaml

removedb:
	sls remove --config sls_persistent_resources.yaml

removetest:
	sls remove --config sls_test.yaml
