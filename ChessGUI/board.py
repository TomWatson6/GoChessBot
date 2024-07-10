import ast

class Board:
    def __init__(self, board_data):
        # Convert pieces to a map of tuple to string
        pieces = {ast.literal_eval(k): v for k, v in board_data['board']['pieces'].items()}
        pieces = {(k[1] if board_data['turn'] == "Black" else board_data['board']['height'] - k[1] + 1, k[0]): v for k, v in pieces.items()}
        
        # Convert power to a map of tuple to array of tuples
        power = {ast.literal_eval(k): [ast.literal_eval(t) for t in v.strip('[]').split()] for k, v in board_data['board']['power'].items()}
        power = {(k[1], k[0]): v for k, v in power.items()}

        for k, v in power.items():
            temp = []
            for i in v:
                temp.append((i[1] if board_data['turn'] == "Black" else board_data['board']['height'] - i[1] + 1, i[0]))

            power[k] = temp

        self.width = board_data['board']['width']
        self.height = board_data['board']['height']
        self.pieces = pieces
        self.power = power
        self.history = board_data['board']['history']
        self.turn = board_data['turn']