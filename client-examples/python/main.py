import asyncio
import websockets
import json

def handle_message(message):
    js = json.loads(message)
    print(js["Source"], js["User"], js["Comment"])


async def echo(websocket, path):
    async for message in websocket:
        handle_message(message)
        await websocket.send(message)

async def main():
    async with websockets.serve(echo, 'localhost', 8839):
        await asyncio.Future()  # run forever

asyncio.run(main())