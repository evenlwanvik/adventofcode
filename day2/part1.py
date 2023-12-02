
def cube_set_is_possible(cube_set: str) -> bool:
    for i, cube in enumerate(cube_set.strip().split(",")):
        cube_color = cube.strip().split(" ")[1]
        n_cubes = int(cube.strip().split(" ")[0])
        match cube_color:
            case "red":
                if n_cubes > 12:
                    return False
            case "green":
                if n_cubes > 13:
                    return False
            case "blue":
                if n_cubes > 14:
                    return False
    return True

def game_is_possible(game: str) -> bool:
    for cube_set in game.strip().split(";"):
        if not cube_set_is_possible(cube_set):
            return False
    return True

with open('./input.txt') as f:

    valid_game_id_sum = 0
    for i, line in enumerate(f):
        game = line.strip().split(":", 1).pop(1)
        if game_is_possible(game):
            valid_game_id_sum += i + 1
            print(f"Game {i + 1} is valid")
        else:
            print(f"Game {i + 1} is not valid")

print(f"Sum of valid game IDs: {valid_game_id_sum}")



