import requests
from board import *

def get_random_board():
    url = "http://localhost:8000/startRandom"
    response = requests.get(url)
    
    if response.status_code == 200:
        board_data = response.json()
        
        return Board(board_data)
    else:
        print(f"Failed to retrieve board: {response.status_code}")

def new_game():
    url = "http://localhost:8000/start"