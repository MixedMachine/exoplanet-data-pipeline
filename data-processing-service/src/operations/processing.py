from src.operations.data import load_data, save_data

def process(exoplanet_id: str) -> str | None:
    print("processing data...")
    planet = load_data(exoplanet_id)
    save_data(planet)

