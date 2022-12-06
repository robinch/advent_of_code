use std::collections::HashSet;
use std::fs;

pub fn part_1(file_path: &str) -> usize {
    let input = fs::read_to_string(file_path).expect("Could not read file");
    let signals: Vec<char> = input.chars().collect();

    find_marker(&signals, 4)
}

pub fn part_2(file_path: &str) -> usize {
    let input = fs::read_to_string(file_path).expect("Could not read file");
    let signals: Vec<char> = input.chars().collect();

    find_marker(&signals, 14)
}

fn find_marker(signals: &Vec<char>, nr_of_unique_values: usize) -> usize {
    let mut set: HashSet<char> = HashSet::new();
    let mut first_marker: Option<usize> = None;

    for i in 0..(signals.len() - nr_of_unique_values) {
        let mut dup_found = false;
        set.clear();

        for j in i..(i + nr_of_unique_values) {
            if !set.insert(signals[j]) {
                dup_found = true;
                break;
            }
        }

        if !dup_found {
            first_marker = Some(i + nr_of_unique_values);
            break;
        }
    }

    match first_marker {
        None => panic!("No marker found!"),
        Some(first_marker) => first_marker,
    }
}
