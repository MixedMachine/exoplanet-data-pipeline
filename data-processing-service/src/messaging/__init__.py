import nats

class NatsClient:
    client: any
    connected: bool
    subscriptions: dict
    url: str
  
    def __init__(self, url: str = "localhost:4222") -> None:
        self.client = None
        self.connected = False
        self.subscriptions = {}

    async def connect(self, url: str|None = None) -> None:
        if url is None:
            url =-self.url
        self.client = await nats.connect(url)
        # self.connected = True

    async def subscribe(self, subject: str, callback: str):
        if not self.connected:
            await self.connect()

        subscription = await self.client.subscribe(subject)
        self.subscriptions[subject] = subscription

        while subject in self.subscriptions:
            message = await subscription.next_msg()
            if message is None: continue
            callback(message)

    async def publish(self, subject, data) -> None:
        if not self.connected:
            await self.connect();

        await self.client.publish(subject, data);

    async def close(self) -> None:
        if not self.connected:
            return

        map(lambda subscription: subscription.unsubscribe(),
            self.subscriptions.values())
        self.subscriptions = {}
        await self.client.close()
        self.connected = False
