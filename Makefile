GOCMD=go
    GORUN=$(GOCMD) run
    GOCLEAN=$(GOCMD) clean

run_server:
	$(GORUN) simple-server/main/main.go

run_client:
	$(GORUN) main.go

clean:
	$(GOCLEAN)