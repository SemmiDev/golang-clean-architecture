cup:
	sudo docker-compose --compatibility up -d --build
cup-prod:
	sudo docker-compose --compatibility -f docker-compose.prod.yml up -d  --build
cdown:
	sudo docker-compose down