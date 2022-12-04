use std::fs;

struct Pair(Section, Section);

struct Section {
    from: u32,
    to: u32,
}

pub fn part_1(file_path: &str) -> u32 {
    let inputs = fs::read_to_string(file_path).expect("Could not read file!");

    let mut fully_overlapping_pairs: u32 = 0;

    for pair in parse_to_pairs(&inputs) {
        if fully_overlapping(&pair) {
            fully_overlapping_pairs += 1;
        }
    }

    fully_overlapping_pairs
}

pub fn part_2(file_path: &str) -> u32 {
    let inputs = fs::read_to_string(file_path).expect("Could not read file!");
    let mut overlapping_pairs: u32 = 0;

    for pair in parse_to_pairs(&inputs) {
        if partially_overlapping(&pair) {
            overlapping_pairs += 1;
        }
    }

    overlapping_pairs
}

fn fully_overlapping(pair: &Pair) -> bool {
    (pair.0.from <= pair.1.from && pair.0.to >= pair.1.to)
        || (pair.0.from >= pair.1.from && pair.0.to <= pair.1.to)
}

fn partially_overlapping(pair: &Pair) -> bool {
    (pair.0.from >= pair.1.from && pair.0.from <= pair.1.to)
        || (pair.0.to >= pair.1.from && pair.0.to <= pair.1.to)
        || (pair.1.from >= pair.0.from && pair.1.from <= pair.0.to)
        || (pair.1.to >= pair.0.from && pair.1.to <= pair.0.to)
}

fn parse_to_pairs(inputs: &str) -> Vec<Pair> {
    let mut pairs: Vec<Pair> = Vec::new();

    for input in inputs.lines() {
        let pair: Vec<&str> = input.split(",").collect();
        let sections_for_a: Vec<&str> = pair[0].split("-").collect();
        let sections_for_b: Vec<&str> = pair[1].split("-").collect();

        let section_a = Section {
            from: sections_for_a[0].parse().unwrap(),
            to: sections_for_a[1].parse().unwrap(),
        };

        let section_b = Section {
            from: sections_for_b[0].parse().unwrap(),
            to: sections_for_b[1].parse().unwrap(),
        };

        pairs.push(Pair(section_a, section_b));
    }

    pairs
}
