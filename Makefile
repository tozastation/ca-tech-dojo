DEV_DOCKER_COMPOSE_FILE='./deployments/docker-compose.dev.yaml'
up-dev:
	docker-compose -f ${DEV_DOCKER_COMPOSE_FILE} build
	docker-compose -f ${DEV_DOCKER_COMPOSE_FILE} up -d mysql
	sleep 10
	docker-compose -f ${DEV_DOCKER_COMPOSE_FILE} up game