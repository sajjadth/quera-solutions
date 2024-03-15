inp = list(input())
board = [[0, 0, 0, 0, 0, 0, 0, 0], [1, 0, 0, 0, 0, 0, 0, 0]]
dead = False

x, y = 0, 1


def print_board(board):
    for row in board:
        print(''.join(str(num) for num in row))


for move in inp:
    x += 1
    if x == 8:
        dead = True
        break
    if move == "L":
        if y == 0:
            dead = True
            break
        else:
            y = 0
    elif move == "R":
        if y == 1:
            dead = True
            break
        else:
            y = 1

    board[y][x] = 1

if dead:
    print("DEATH")
else:
    print_board(board)
