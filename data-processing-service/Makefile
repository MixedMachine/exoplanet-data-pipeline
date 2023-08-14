.PHONY: install run test lint

install:
	poetry install

run:
	poetry run python ./src/data_processing_service/main.py

test:
	poetry run pytest tests/

lint:
	poetry run flake8 src/