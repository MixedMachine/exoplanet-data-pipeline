import datetime
from bson.objectid import ObjectId

from src.databases.models import Exoplanet
from src.databases.sqlite import sqlite_client, sqlite_cursor, insert_into_table
from src.databases.mongodb import k2pandc_coll


def load_data(exoplanet_id: str) -> dict:
    print("loading data...")
    planet = k2pandc_coll.find_one({"_id": ObjectId(exoplanet_id)})
    return planet

def save_data(planet: dict):
    print("saving data...")
    planetObj = Exoplanet(
        planet.get("_id", "NULL"),
        planet.get("pl_name", "NULL"),
        planet.get("pl_bmasse", "NULL"),
        planet.get("pl_rade", "NULL"),
        planet.get("pl_orbper", "NULL"),
        planet.get("sy_dist", "NULL"),
        planet.get("disc_year", "NULL"),
        planet.get("discoverymethod", "NULL"),
        planet.get("disposition", "NULL") == "CONFIRMED",
        planet.get("rowupdate", "NULL"),
        datetime.datetime.now(),
        datetime.datetime.now(),
    )
    planet_data = planetObj.to_dict()
    
    insert_into_table("exoplanets", **planet_data)
