def print_grid_with_borders():
    grid = [
        [0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
        [0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
        [0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
    ]
    rows = len(grid)
    cols = len(grid[0])

    for i in range(rows):
        if i == 0:
            print("\u250c" + "\u2500" * (4 * cols - 1) + "\u2510")  # Top border
        else:
            print("\u251c" + "\u2500" * (4 * cols - 1) + "\u2524")  # Middle borders

        for j in range(cols):
            if grid[i][j] == 1:
                print(
                    "\u2502 \033[91m1\033[0m", end=" "
                )  # Red 1 within vertical border
            else:
                print(
                    "\u2502", grid[i][j], end=" "
                )  # Normal value within vertical border
        print("\u2502")  # Right vertical border

    print("\u2514" + "\u2500" * (4 * cols - 1) + "\u2518")  # Bottom border
