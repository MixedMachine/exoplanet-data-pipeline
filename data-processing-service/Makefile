.PHONY: install run test lint

install:
	poetry install

run:
	poetry run python -m src.data_processing_service

test:
	poetry run pytest tests/

lint:
	poetry run flake8 src/