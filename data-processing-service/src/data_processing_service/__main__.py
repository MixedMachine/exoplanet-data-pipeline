

import asyncio
import json
import os
from dataclasses import asdict, dataclass
import nats
from nats import errors

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

def load_data():
    print("loading data...")

def save_data():
    print("saving data...")

def process():
    print("processing data...")
    load_data()
    save_data()

async def main():

    is_done = asyncio.Future()

    client = await nats.connect(NATS_URL)
    async def ingest_handler(message):
        print("message received:")
        print("channel|" + message.subject)
        print("message|" + message.data.decode('utf-8'))
        print("---")

        try:
            process()
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
