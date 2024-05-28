import requests
import ast
from board import *

def get_random_board():
    url = "http://localhost:8000/randomboard"
    response = requests.get(url)
    
    if response.status_code == 200:
        board_data = response.json()
        
        # Convert pieces to a map of tuple to string
        pieces = {ast.literal_eval(k): v for k, v in board_data['pieces'].items()}
        pieces = {(k[1], k[0]): v for k, v in pieces.items()}
        
        # Convert power to a map of tuple to array of tuples
        power = {ast.literal_eval(k): [ast.literal_eval(t) for t in v.strip('[]').split()] for k, v in board_data['power'].items()}
        power = {(k[1], k[0]): v for k, v in power.items()}

        for k, v in power.items():
            temp = []
            for i in v:
                temp.append((i[1], i[0]))

            power[k] = temp
        
        board = Board(
            board_data['width'], 
            board_data['height'],
            pieces,
            board_data['history'],
            power
        )

        return board
        
    else:
        print(f"Failed to retrieve board: {response.status_code}")