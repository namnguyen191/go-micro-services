# Docker Swarm

- Start the swarm: `docker swarm init`

- Start all services: `docker stack deploy -c swarm.yml myapp`

- Remove all services: `docker stack rm myapp`

- Leave the swarm (take down manager node): `docker swarm leave --force`

- Scale a service: `docker service scale myapp_logger-service=2`

- Update a service: `docker serivce update --image namng191/logger-service:1.0,1 myapp_logger-service`
