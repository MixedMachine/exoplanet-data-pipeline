import pymongo
from bson.objectid import ObjectId
import sqlite3
import pprint
import asyncio
import json
import os
from dataclasses import asdict, dataclass
import nats
import datetime

NATS_URL = os.getenv("NATS_URL", "nats://localhost:4222")
NATS_CHANNEL_BASE = "exoplanets"
NATS_CHANNEL_EXO_INGESTED = f"{NATS_CHANNEL_BASE}.ingested"
NATS_CHANNEL_EXO_PROCESSED = f"{NATS_CHANNEL_BASE}.processed"
MSG_STATUS_COMPLETE = "complete"
MSG_STATUS_FAILED = "failed"

@dataclass
class Payload:
    status: str
    _id: str

class Exoplanet:
    _id: str
    name: str
    mass: float
    radius: float
    period: float
    distance: float
    year_discovered: int
    method: str
    confirmed: bool
    planet_last_updated: str
    created_at: datetime.datetime
    last_updated: datetime.datetime

    def __init__(self, _id, name, mass, radius, period, distance, year_discovered, method, confirmed, planet_last_updated, created_at, last_updated):
        self._id = str(_id)
        self.name = str(name)
        if mass == None or mass == "NULL" or mass == "NaN":
            self.mass = 0.0
        else:
            self.mass = float(mass)
        if radius == None or radius == "NULL" or radius == "NaN":
            self.radius = 0.0
        else:
            self.radius = float(radius)
        if period == None or period == "NULL" or period == "NaN":
            self.period = 0.0
        else:
            self.period = float(period)
        if distance == None or distance == "NULL" or distance == "NaN":
            self.distance = 0.0
        else:
            self.distance = float(distance)
        if year_discovered == None or year_discovered == "NULL" or year_discovered == "NaN":
            self.year_discovered = 0
        else:
            self.year_discovered = int(year_discovered)
        self.method = str(method)
        self.confirmed = bool(confirmed)
        self.planet_last_updated = planet_last_updated
        self.created_at = created_at
        self.last_updated = last_updated

mongo_client = pymongo.MongoClient("mongodb://root:root@localhost:27017/")
exoplanets_db = mongo_client["exoplanets"]
k2pandc_coll = exoplanets_db["k2pandc"]

sqlite_client = sqlite3.connect("../data/sqlite/exoplanets.db")
sqlite_cursor = sqlite_client.cursor()
res = sqlite_cursor.execute("""
CREATE TABLE IF NOT EXISTS exoplanets (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE,
    mass REAL,
    radius REAL,
    period REAL,
    distance REAL,
    year_discovered INTEGER,
    method TEXT,
    confirmed BOOLEAN,
    planet_last_updated INTEGER,
    created_at TEXT,
    last_updated TEXT
);
""")

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
    sqlite_cursor.execute("""
    INSERT INTO exoplanets (
        id,
        name,
        mass,
        radius,
        period,
        distance,
        year_discovered,
        method,
        confirmed,
        planet_last_updated,
        created_at,
        last_updated
    ) VALUES (
        ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    );
    """,(
        planetObj._id,
        planetObj.name,
        planetObj.mass,
        planetObj.radius,
        planetObj.period,
        planetObj.distance,
        planetObj.year_discovered,
        planetObj.method,
        planetObj.confirmed,
        planetObj.planet_last_updated,
        planetObj.created_at.strftime("%Y-%m-%d"),
        planetObj.last_updated.strftime("%Y-%m-%d")
    ))
    sqlite_client.commit()

def process(exoplanet_id: str):
    print("processing data...")
    planet = load_data(exoplanet_id)
    save_data(planet)

async def main():
    client = await nats.connect(NATS_URL)
    async def ingest_handler(message):
        print("message received:")
        print("channel|" + message.subject)
        data: str  = message.data.decode('utf-8')
        print("message|" + data)
        print("---")

        try:
            process(data)
            payload = Payload(status=MSG_STATUS_COMPLETE, _id=message.data.decode('utf-8'))
            publish_message = json.dumps(asdict(payload))
        except Exception as e:
            print(e)
            payload = Payload(status=MSG_STATUS_FAILED, _id=message.data.decode('utf-8'))
            publish_message = json.dumps(asdict(payload))
        finally:
            await client.publish(NATS_CHANNEL_EXO_PROCESSED, publish_message.encode('utf-8'))

    async def proccess_handler(message):
        print("message received:")
        print("channel|" + message.subject)
        print("message|" + message.data.decode('utf-8'))
        print("---")

    await client.subscribe(NATS_CHANNEL_EXO_INGESTED, cb=ingest_handler)
    await client.subscribe(NATS_CHANNEL_EXO_PROCESSED, cb=proccess_handler)


    while True:
        await asyncio.sleep(1)
        print("waiting for messages...", flush=True, end="\r")


if __name__ == "__main__":
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        os._exit(1)
