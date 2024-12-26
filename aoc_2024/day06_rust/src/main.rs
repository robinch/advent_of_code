use std::collections::HashSet;
use std::fs;

fn main() {
    part1();
}

#[derive(Debug)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[derive(Debug)]
struct Map {
    starting_pos: (i32, i32),
    obstacles: HashSet<(i32, i32)>,
    height: i32,
    width: i32,
}

impl Map {
    fn out_of_bounds(&self, pos: &(i32, i32)) -> bool {
        pos.0 < 0 || pos.1 < 0 || pos.0 >= self.height || pos.1 >= self.width
    }
}

#[derive(Debug)]
struct Guard {
    pos: (i32, i32),
    direction: Direction,
}

impl Guard {
    fn forward_pos(&self) -> (i32, i32) {
        match self.direction {
            Direction::Up => (self.pos.0 - 1, self.pos.1),
            Direction::Right => (self.pos.0, self.pos.1 + 1),
            Direction::Down => (self.pos.0 + 1, self.pos.1),
            Direction::Left => (self.pos.0, self.pos.1 - 1),
        }
    }

    fn move_forward(&mut self) -> () {
        self.pos = self.forward_pos();
    }

    fn turn(&mut self) -> () {
        self.direction = match self.direction {
            Direction::Up => Direction::Right,
            Direction::Right => Direction::Down,
            Direction::Down => Direction::Left,
            Direction::Left => Direction::Up,
        };
    }
}

fn part1() {
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    let map = parse_input("input/input");

    let mut guard = Guard {
        pos: map.starting_pos,
        direction: Direction::Up,
    };

    loop {
        visited.insert(guard.pos);

        let next_pos = guard.forward_pos();

        if map.out_of_bounds(&next_pos) {
            break;
        } else if map.obstacles.contains(&next_pos) {
            guard.turn();
        } else {
            guard.move_forward();
        }
    }

    println!("Day 6, part 1: {}", visited.len());
}

// The origo is at the top left corner, (row, col)
fn parse_input(path: &str) -> Map {
    let mut starting_pos: (i32, i32) = (-1, -1);
    let mut obstacles: HashSet<(i32, i32)> = HashSet::new();

    let input = fs::read_to_string(path).unwrap();
    let content: Vec<String> = input.lines().map(|line| line.to_string()).collect();

    let width = content[0].len() as i32;
    let height = content.len() as i32;

    for (i, line) in content.iter().enumerate() {
        for (j, c) in line.chars().enumerate() {
            match c {
                '^' => starting_pos = (i as i32, j as i32),
                '#' => {
                    obstacles.insert((i as i32, j as i32));
                }
                _ => (),
            }
        }
    }

    Map {
        starting_pos,
        obstacles,
        height,
        width,
    }
}
