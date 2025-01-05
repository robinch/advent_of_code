use std::collections::HashSet;
use std::fs;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
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

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
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
    let map = parse_input("input/input");

    let guard = Guard {
        pos: map.starting_pos,
        direction: Direction::Up,
    };

    let unique_positions = get_all_unique_guard_positions(&guard, &map);

    println!("Day 6, part 1: {}", unique_positions.len());
}

fn part2() {
    let map = parse_input("input/input");

    let guard = Guard {
        pos: map.starting_pos,
        direction: Direction::Up,
    };

    let mut positions = get_all_unique_guard_positions(&guard, &map);
    positions.remove(&guard.pos);

    let mut count = 0;

    for extra_obstacle in positions {
        if has_loop(&extra_obstacle, &guard, &map) {
            count += 1;
        }
    }

    println!("Day 6, part 2: {}", count);
}

fn has_loop(extra_obstacle: &(i32, i32), current_guard: &Guard, map: &Map) -> bool {
    let mut visited: HashSet<Guard> = HashSet::new();

    let mut obstacles = map.obstacles.clone();
    obstacles.insert(*extra_obstacle);

    let mut guard = current_guard.clone();

    loop {
        let inserted = visited.insert(guard.clone());

        if !inserted {
            return true;
        }

        let next_pos = guard.forward_pos();

        if map.out_of_bounds(&next_pos) {
            break;
        } else if obstacles.contains(&next_pos) {
            guard.turn();
        } else {
            guard.move_forward();
        }
    }

    false
}

fn get_all_unique_guard_positions(start_guard: &Guard, map: &Map) -> HashSet<(i32, i32)> {
    let mut positions: Vec<Guard> = Vec::new();
    let mut unique_positions: HashSet<(i32, i32)> = HashSet::new();

    let mut guard = start_guard.clone();

    loop {
        positions.push(guard.clone());

        let next_pos = guard.forward_pos();

        if map.out_of_bounds(&next_pos) {
            break;
        } else if map.obstacles.contains(&next_pos) {
            guard.turn();
        } else {
            guard.move_forward();
        }
    }
    for p in positions {
        unique_positions.insert(p.pos);
    }

    unique_positions
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

fn main() {
    part1();
    part2();
}
