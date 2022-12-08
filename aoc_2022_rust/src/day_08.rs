use std::fs;

pub fn part_1(file_path: &str) -> u32 {
    let forrest = generate_forrest(file_path);
    let size = forrest.size();

    let mut visible_trees: u32 = 0;

    for y in 1..size - 1 {
        for x in 1..size - 1 {
            let height = forrest.get_height(x, y);

            let row = forrest.get_row(y);
            let column = forrest.get_column(x);

            let left = &row[0..x];
            let right = &row[x + 1..size];

            let up = &column[0..y];
            let down = &column[y + 1..size];

            if higher_then_rest(height, left) {
                visible_trees += 1;
            } else if higher_then_rest(height, right) {
                visible_trees += 1;
            } else if higher_then_rest(height, up) {
                visible_trees += 1;
            } else if higher_then_rest(height, down) {
                visible_trees += 1;
            }
        }
    }

    visible_trees + 4 * size as u32 - 4
}

pub fn part_2(file_path: &str) -> u32 {
    let forrest = generate_forrest(file_path);
    let size = forrest.size();

    let mut max_scenic_score: u32 = 0;

    for y in 1..size - 1 {
        for x in 1..size - 1 {
            let height = forrest.get_height(x, y);
            let row = forrest.get_row(y);
            let column = forrest.get_column(x);

            let mut left = row[0..x].to_vec();
            left.reverse();

            let right = &row[x + 1..size].to_vec();

            let mut up = column[0..y].to_vec();
            up.reverse();

            let down = &column[y + 1..size].to_vec();

            let scenic_score = visible_trees(&height, &left)
                * visible_trees(&height, &right)
                * visible_trees(&height, &up)
                * visible_trees(&height, &down);

            if scenic_score > max_scenic_score {
                max_scenic_score = scenic_score;
            }
        }
    }

    max_scenic_score
}

#[derive(Debug)]
struct Forrest {
    grid: Vec<Vec<u8>>,
}

impl Forrest {
    fn new(size: usize) -> Forrest {
        let mut grid: Vec<Vec<u8>> = Vec::new();

        for _ in 0..size {
            let row: Vec<u8> = vec![0; size];
            grid.push(row);
        }

        Forrest { grid: grid }
    }

    fn set_height(self: &mut Forrest, x: usize, y: usize, height: u8) {
        self.grid[y][x] = height;
    }

    fn get_height(self: &Forrest, x: usize, y: usize) -> &u8 {
        &self.grid[y][x]
    }

    fn get_row(self: &Forrest, y: usize) -> Vec<u8> {
        self.grid[y].to_vec()
    }

    fn get_column(self: &Forrest, x: usize) -> Vec<u8> {
        let len = self.grid.len();
        let mut column: Vec<u8> = vec![0; len];
        for y in 0..len {
            let height = &self.grid[y][x];
            column[y] = *height;
        }
        column
    }

    fn size(self: &Forrest) -> usize {
        self.grid.len()
    }
}
fn higher_then_rest(height: &u8, rest: &[u8]) -> bool {
    !rest.contains(height) && (*rest.iter().max().unwrap() as u8) < *height
}

fn visible_trees(height: &u8, trees: &Vec<u8>) -> u32 {
    let mut visible_trees: u32 = 0;

    for i in 0..trees.len() {
        visible_trees += 1;

        if trees[i] >= *height {
            break;
        }
    }

    visible_trees
}

fn generate_forrest(file_path: &str) -> Forrest {
    let input = fs::read_to_string(file_path).expect("Could not read file!");

    let size = input.lines().count();

    let mut forrest = Forrest::new(size);

    for (y, row) in input.lines().enumerate() {
        for (x, height) in row.chars().enumerate() {
            forrest.set_height(x, y, height.to_digit(10).unwrap() as u8);
        }
    }

    forrest
}
