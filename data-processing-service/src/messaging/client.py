from dataclasses import asdict
import json
import nats
import os

from src.messaging.constants import (
    NATS_URL,
    NATS_CHANNEL_EXO_INGESTED,
    NATS_CHANNEL_EXO_PROCESSED,
    MSG_STATUS_COMPLETE,
    MSG_STATUS_FAILED,
)
from src.messaging.models import Payload
from src.operations.processing import process

async def set_up_messaging():
    try:
        client = await nats.connect(NATS_URL)
    except Exception as e:
        print(e)
        os.exit(1)

    async def ingest_handler(message):
        data: str  = message.data.decode('utf-8')
        print(f"message received: {message.subject} - {data}")
        print("---")

        try:
            process(data)
            payload = Payload(status=MSG_STATUS_COMPLETE, _id=message.data.decode('utf-8'), message="")
            publish_message = json.dumps(asdict(payload))
        except Exception as e:
            print("ERROR: ", e)
            payload = Payload(status=MSG_STATUS_FAILED, _id=message.data.decode('utf-8'), message=str(e))
            publish_message = json.dumps(asdict(payload))
        finally:
            await client.publish(NATS_CHANNEL_EXO_PROCESSED, publish_message.encode('utf-8'))

    await client.subscribe(NATS_CHANNEL_EXO_INGESTED, cb=ingest_handler)
