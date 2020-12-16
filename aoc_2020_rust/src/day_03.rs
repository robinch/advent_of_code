use std::collections::HashSet;
use std::fs;

pub fn part_1(input_path: &str, right: i32, down: i32) -> i32 {
    let tree_map = tree_map(input_path);
    tree_collisions(&tree_map, right, down)
}

pub fn part_2(input_path: &str, steps: Vec<(i32, i32)>) -> i64 {
    let tree_map = tree_map(input_path);
    let mut res: i64 = 1;
    for (right, down) in steps {
        res *= tree_collisions(&tree_map, right, down) as i64;
        println!("RES: {}", res)
    }
    res
}

fn tree_collisions(tree_map: &TreeMap, right: i32, down: i32) -> i32 {
    let mut tree_counter = 0;
    let mut x = 0;
    let mut y = 0;
    loop {
        x = (x + right) % tree_map.width;
        y = y + down;
        if y >= tree_map.height {
            break;
        } else {
            if tree_map.tree_exists_at_coordinate(x, y) {
                tree_counter += 1;
            } else {
            }
        }
    }

    tree_counter
}
fn tree_map(input_path: &str) -> TreeMap {
    let input = fs::read_to_string(input_path).expect("Could not open file");
    let mut tree_map = TreeMap::new();
    let mut y = 0;
    let mut x = 0;
    for row in input.lines() {
        x = 0;
        let row_chars: Vec<char> = row.chars().collect();
        for c in row_chars {
            if c == '#' {
                tree_map.set_tree_at(x, y);
            }
            x += 1;
        }
        y += 1;
    }

    tree_map.set_width(x);
    tree_map.set_height(y);
    tree_map
}

#[derive(Debug)]
struct TreeMap {
    tree_coordinates: HashSet<(i32, i32)>,
    width: i32,
    height: i32,
}

impl TreeMap {
    pub fn new() -> TreeMap {
        TreeMap {
            tree_coordinates: HashSet::new(),
            width: 0,
            height: 0,
        }
    }

    pub fn tree_exists_at_coordinate(&self, x: i32, y: i32) -> bool {
        self.tree_coordinates.contains(&(x, y))
    }

    pub fn set_tree_at(&mut self, x: i32, y: i32) {
        self.tree_coordinates.insert((x, y));
    }

    pub fn set_width(&mut self, width: i32) {
        self.width = width;
    }

    pub fn set_height(&mut self, height: i32) {
        self.height = height;
    }
}
