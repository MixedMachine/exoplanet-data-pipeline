from dataclasses import asdict
import json
import nats
import os

from src.messaging.constants import (
    NATS_URL,
    NATS_CHANNEL_EXO_INGESTED,
    NATS_CHANNEL_EXO_PROCESSED,
)
from src.operations.processing import run_message_processing


async def set_up_messaging():
    try:
        client = await nats.connect(NATS_URL)
    except Exception as e:
        print(e)
        os.exit(1)

    async def ingest_handler(message):
        payload = run_message_processing(message)
        publish_message = json.dumps(asdict(payload))
        await client.publish(NATS_CHANNEL_EXO_PROCESSED, publish_message.encode('utf-8'))

    await client.subscribe(NATS_CHANNEL_EXO_INGESTED, cb=ingest_handler)

