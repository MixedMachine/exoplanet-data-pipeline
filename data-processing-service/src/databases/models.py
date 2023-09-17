import datetime

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
