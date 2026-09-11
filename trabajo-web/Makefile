# indica a la herramienta make que "test" es el nombre de una tarea y no el nombre de un archivo 
.PHONY: test

test:
	# limpia ejec ant 
	docker compose down -v

	# levanta postgres en 2do plano 
	docker compose up -d postgresql

	# esperar a que se incie base de datos
	@until docker exec postgres_container pg_isready -U postgres > /dev/null 2>&1; do \
		sleep 1; \
	done

	# generar cod go (a partir de consultas sql)
	sqlc generate 

	# ejecuta test de go 
	go test -v ./...

	# limpia al finalizar
	docker compose down -v
