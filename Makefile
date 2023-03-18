# Crear imagen de Docker
docker-build:
	docker build -t gcs-upload .

# Crear contenedor de Docker
docker-create:
	docker run --name gcs-upload_prod -dp 8080:8080 gcs-upload

# Ejecutar el contenedor ya creado
docker-run:
	docker start gcs-upload_prod

# Detener contenedor de Docker
docker-stop:
	docker stop gcs-upload_prod

# Eliminar contenedor de Docker
docker-rm:
	docker rm gcs-upload_prod

# Eliminar imagen de Docker
docker-rmi:
	docker rmi gcs-upload

# PELIGROSO: Purgar cache de Docker. Úsalo solo si sabes lo que haces
docker-purge:
	docker system prune -af