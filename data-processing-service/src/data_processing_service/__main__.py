import asyncio
import os

from src.messaging.client import set_up_messaging


async def main():
    await set_up_messaging()

    while True:
        await asyncio.sleep(1)
        print("...", flush=True, end="\r")


if __name__ == "__main__":
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        os._exit(1)
