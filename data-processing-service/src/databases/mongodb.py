import pymongo
from bson.objectid import ObjectId


mongo_client = pymongo.MongoClient("mongodb://root:root@localhost:27017/")
exoplanets_db = mongo_client["exoplanets"]
k2pandc_coll = exoplanets_db["k2pandc"]
