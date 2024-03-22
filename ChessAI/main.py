from fastapi import FastAPI, HTTPException
from typing import Dict, List
from pydantic import BaseModel
import httpx

class ChessGameState(BaseModel):
    width: int
    height: int
    pieces: Dict[str, str]
    history: List[str]

class Coord:
    row: int
    col: int

class Position:
    file: Coord
    rank: Coord

class Move:
    from_: Position
    to: Position

app = FastAPI()

@app.get("/start")
async def start():
    async with httpx.AsyncClient() as client:
        response = await client.get("http://localhost:8000/start")
        response.raise_for_status()
    return {"message": "Game started successfully"}

@app.get("/state")
async def state():
    async with httpx.AsyncClient() as client:
        response = await client.get("http://localhost:8000/state")
        response.raise_for_status()
        game_state = ChessGameState.model_validate(response.json())
    return game_state.model_dump()

@app.post("/move")
async def move(move: Move):
    async with httpx.AsyncClient() as client:
        response = await client.post("http://localhost:8000/move", json=dict(move))
        response.raise_for_status()
    return {"message": "Move successfully made"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080)