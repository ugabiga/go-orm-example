bo-create:
	@sqlboiler psql --config examples/boiler/sqlboiler.toml

up:
	@docker-compose -f docker-compose.yml up -d

down:
	@docker-compose -f docker-compose.yml down
