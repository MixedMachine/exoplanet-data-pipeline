from dataclasses import dataclass

@dataclass
class Payload:
    _id: str
    status: str
    message: str

