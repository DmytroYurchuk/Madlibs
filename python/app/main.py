import aiohttp
import asyncio
from fastapi import FastAPI
from typing import Union

app : FastAPI = FastAPI()

async def fetch_word(type) -> Union[str, ValueError]:
    async with aiohttp.ClientSession() as session:
        async with session.get(f"https://reminiscent-steady-albertosaurus.glitch.me/{type}") as response:
            response_text = await response.json()
            if response.status != 200:
                raise ValueError(f"Failed to fetch {type}: {response_text}")
            return response_text

@app.get("/madlib")
async def madlib() -> str:
    try:
        adjective, verb, noun = await asyncio.gather(fetch_word("adjective"), fetch_word("verb"), fetch_word("noun"))
    except Exception as e:
        return f"Error: {str(e)}"
    return f"It was a {adjective} day. I went downstairs to see if I could {verb} dinner. I asked, 'Does the stew need fresh {noun}?'"
