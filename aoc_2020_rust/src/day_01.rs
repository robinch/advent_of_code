use std::collections::HashSet;
use std::fs;

pub fn part_1(input_path: &str, target: i32) -> i32 {
    let inputs = fs::read_to_string(input_path).expect("Could not read file");

    let mut numbers = HashSet::new();

    for input in inputs.lines() {
        let n = input.parse::<i32>().unwrap();
        numbers.insert(n);
    }

    for n in numbers.iter() {
        let m = target - n;
        if numbers.contains(&m) {
            return n * m;
        }
    }

    -1
}

pub fn part_2(input_path: &str, target: i32) -> i32 {
    let inputs = fs::read_to_string(input_path).expect("Could not read file");

    let mut numbers = HashSet::new();

    for input in inputs.lines() {
        let n = input.parse::<i32>().unwrap();
        numbers.insert(n);
    }

    for x in numbers.iter() {
        for y in numbers.iter() {
            let z = target - y - x;
            if numbers.contains(&z) {
                return x * y * z;
            }
        }
    }

    -1
}
