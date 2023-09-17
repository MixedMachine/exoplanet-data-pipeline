import os

NATS_URL = os.getenv("NATS_URL", "nats://localhost:4222")
NATS_CHANNEL_BASE = "exoplanets"
NATS_CHANNEL_EXO_INGESTED = f"{NATS_CHANNEL_BASE}.ingested"
NATS_CHANNEL_EXO_PROCESSED = f"{NATS_CHANNEL_BASE}.processed"
MSG_STATUS_COMPLETE = "complete"
MSG_STATUS_FAILED = "failed"
