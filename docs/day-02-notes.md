# Day 2 - Docker Compose  

## Docker Commands/Theory

- docker compose = use docker-compose.yml
- up = create/start services
- -d = run in the background
- docker compose down = stop and remove the container
- docker compose ps = check
- docker compose up -d = start it again
- docker compose stop = just stops the container but keeps it
- images = class; container = object;
- Ports thoery -> host/computer port : container port

## PostgreSQL configuration 

- Image: postgres:17 
- Database: wallet 
- User: postgres 
- Password: password 
- Port: 5432 
- Volume: postgres_data 

## Commands
- Validate configuration: docker compose config
- Start PostgreSQL: docker compose up -d  
- Check running containers: docker compose ps
- Stop PostgreSQL: docker compose stop 
- Stop and remove container: docker compose down
- Stops everything and removes the named volume too: docker compose down -v

## Understanding
- An image is: A class, like a blueprint, for creating containers.  
- A container is: An object, made from image template.  
- A volume is: Storage that keepsdata outside the containers filesystem 
- `-d` means: Run on background 

## Result

- PostgreSQL container is running: Yes 
- I can stop and start it: Yes
 