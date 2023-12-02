
def get_numbers_from_cube_set(cube_set: str) -> [int, int, int]:
    bgr = [0, 0, 0] # bluegreenred
    for i, cube in enumerate(cube_set.strip().split(",")):
        cube_color = cube.strip().split(" ")[1]
        match cube_color:
            case "blue":
                bgr[0] = int(cube.strip().split(" ")[0])
            case "green":
                bgr[1] = int(cube.strip().split(" ")[0])
            case "red":
                bgr[2] = int(cube.strip().split(" ")[0])
    return bgr

def min_number_of_cubes_per_color_in_game(game: str) -> [int, int, int]:
    min_cubes_bgr = [0, 0, 0] # bluegreenred
    for cube_set in game.strip().split(";"):
        result = get_numbers_from_cube_set(cube_set)
        for i, color_count in enumerate(result):
            if color_count > min_cubes_bgr[i]:
                min_cubes_bgr[i] = color_count
    return min_cubes_bgr

def calc_cube_power(cube: [int, int, int]) -> int:
    return cube[0] * cube[1] * cube[2]

with open('./input.txt') as f:
    sum_of_cube_powers = 0
    for line in f:
        game = line.strip().split(":", 1).pop(1)
        min_cubes_bgr = min_number_of_cubes_per_color_in_game(game)
        print(f"min_cubes_bgr: {min_cubes_bgr}")
        sum_of_cube_powers += calc_cube_power(min_cubes_bgr)

print(f"Sum of cube powers: {sum_of_cube_powers}")


