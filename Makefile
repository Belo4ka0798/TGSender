###########################################
# MakeFile for Deploy all project "Anima"
###########################################

# Docker command for docker compose
DC=docker-compose

# Path to root Server
SERVER_PATH=./Server

# Path to root Deploy
DEPLOY_PATH=./Deploy

.PHONY: all
all: compose-up

.PHONY: compose-up
compose-up:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml up

.PHONY: compose-up-silent
compose-up-silent: proto-cp
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml up -d

.PHONY: compose-start
compose-start:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml start

.PHONY: compose-stop
compose-stop:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml stop

.PHONY: compose-down
compose-down:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml down
	docker rmi deploy-server


