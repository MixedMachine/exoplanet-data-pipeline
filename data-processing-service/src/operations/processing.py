from src.messaging.constants import (
    MSG_STATUS_COMPLETE,
    MSG_STATUS_FAILED,
)
from src.messaging.models import Payload
from src.operations.data import load_data, save_data


def run_message_processing(message) -> Payload:
    status: str
    message: str
    data: str  = message.data.decode('utf-8')

    print(f"message received: {message.subject} - {data}")

    try:
        process(data)
        status = MSG_STATUS_COMPLETE    
        message = ""

    except Exception as e:
        print("ERROR: ", e)
        status=MSG_STATUS_FAILED
        message=str(e)

    print("---")
    return Payload(_id=data, status=status, message=message)

def process(exoplanet_id: str) -> str | None:
    print("processing data...")
    planet = load_data(exoplanet_id)
    save_data(planet)

