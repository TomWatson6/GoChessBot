import matplotlib.pyplot as plt
import numpy as np
import matplotlib.image as mpimg
from matplotlib.offsetbox import OffsetImage, AnnotationBbox
from api import *
from board import *

def plot_chessboard(board_data):
    # Create a grid
    board = np.zeros((8, 8))
    
    # Color the squares with green and buff colors
    for row in range(board_data.height):
        for col in range(board_data.width):
            if (row + col) % 2 == 0:
                board[row, col] = 1
    
    fig, ax = plt.subplots(figsize=(8, 8))
    # Use a custom color map for green and buff
    ax.imshow(board, cmap='coolwarm', extent=[0, 8, 0, 8])
    
    # Add the grid lines
    ax.set_xticks(np.arange(8) + 0.5)
    ax.set_xticklabels(['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'])
    ax.set_yticks(np.arange(8) + 0.5)
    y_labels = ['1', '2', '3', '4', '5', '6', '7', '8']
    if board_data.turn == "White":
        y_labels = y_labels[::-1]
    ax.set_yticklabels(y_labels)
    ax.grid(False)
    
    # Invert y-axis to match chessboard coordinates
    ax.invert_yaxis()
    
    # Add the pieces to the board
    for position, piece in board_data.pieces.items():
        img = mpimg.imread(f'images/{piece}.png')
        imagebox = OffsetImage(img, zoom=0.8)  # Adjust the zoom to fit the piece correctly
        ab = AnnotationBbox(imagebox, (position[1]+0.5, position[0]+0.5), frameon=False)
        ax.add_artist(ab)
    
    # List to keep track of highlighted rectangles
    highlighted_rects = []

    # Event handler for clicking on squares
    def on_click(event):
        nonlocal highlighted_rects
        # Clear previous highlights
        while highlighted_rects:
            rect = highlighted_rects.pop()
            rect.remove()
        if event.inaxes is not None:
            x, y = int(event.xdata), int(event.ydata)
            clicked_pos = (y, x)
            if clicked_pos in board_data.pieces and clicked_pos in board_data.power:
                affected_positions = board_data.power[clicked_pos]
                # Highlight the affected positions
                for pos in affected_positions:
                    rect = plt.Rectangle((pos[1], pos[0]), 1, 1, color='yellow', alpha=0.5)
                    ax.add_patch(rect)
                    highlighted_rects.append(rect)
                plt.draw()
    
    fig.canvas.mpl_connect('button_press_event', on_click)
    plt.show()

if __name__ == "__main__":
    b = get_random_board()
    plot_chessboard(b)
