use core::panic;
use std::cmp::Reverse;
use std::collections::HashSet;
use std::fs;
use std::hash::Hash;
use std::vec;

#[derive(Debug, Eq, Hash, PartialEq)]
struct Pos {
    x: i32,
    y: i32,
    z: i32,
}

#[derive(Debug)]
struct Pair<'a> {
    dist: f64,
    pos1: &'a Pos,
    pos2: &'a Pos,
}

fn main() {
    let positions = read_input("input/real.txt");
    let part1 = part1(&positions, 1000);
    println!("Day08 part1: {}", part1);
    let part2 = part2(&positions);
    println!("Day08 part2: {}", part2);
}

fn part1(positions: &Vec<Pos>, pairs_to_check: usize) -> u32 {
    let sorted_pairs_by_dist = sorted_pairs_by_dist(positions);
    let circuits_to_multiply = 3;

    let mut circuits: Vec<HashSet<&Pos>> = vec![];

    for pair in sorted_pairs_by_dist.iter().take(pairs_to_check) {
        add_pair_to_circuits(&mut circuits, pair);
    }

    let mut circuit_sizes: Vec<_> = circuits
        .iter()
        .map(|circuit| circuit.len() as u32)
        .collect();
    circuit_sizes.sort_by_key(|size| Reverse(*size));

    circuit_sizes.iter().take(circuits_to_multiply).product()
}

fn part2(positions: &Vec<Pos>) -> i64 {
    let mut unconnected: HashSet<&Pos> = positions.iter().collect();
    let mut circuits: Vec<HashSet<&Pos>> = vec![];
    let sorted_pairs_by_dist = sorted_pairs_by_dist(positions);

    let mut last_pair: Option<&Pair> = None;

    for pair in sorted_pairs_by_dist.iter() {
        add_pair_to_circuits(&mut circuits, pair);
        unconnected.remove(pair.pos1);
        unconnected.remove(pair.pos2);

        if unconnected.is_empty() && circuits.len() == 1 {
            last_pair = Some(pair);
            break;
        }
    }

    match last_pair {
        Some(pair) => (pair.pos1.x as i64 * pair.pos2.x as i64)
            .try_into()
            .unwrap(),
        None => panic!("No last pair found"),
    }
}

fn add_pair_to_circuits<'a>(circuits: &mut Vec<HashSet<&'a Pos>>, pair: &Pair<'a>) {
    let pos1_idx = circuits
        .iter()
        .position(|circuit| circuit.contains(pair.pos1));

    let pos2_idx = circuits
        .iter()
        .position(|circuit| circuit.contains(pair.pos2));

    match (pos1_idx, pos2_idx) {
        (None, None) => {
            let mut circuit = HashSet::new();
            circuit.insert(pair.pos1);
            circuit.insert(pair.pos2);

            circuits.push(circuit);
        }

        (Some(i), None) => {
            circuits[i].insert(pair.pos2);
        }
        (None, Some(i)) => {
            circuits[i].insert(pair.pos1);
        }
        (Some(i), Some(j)) if i != j => {
            let min_index = i.min(j);
            let max_index = i.max(j);
            let circuit = circuits.swap_remove(max_index);
            circuits[min_index].extend(circuit);
        }

        _ => {}
    }
}

fn sorted_pairs_by_dist(positions: &Vec<Pos>) -> Vec<Pair<'_>> {
    let mut sorted_pairs_by_dist: Vec<Pair> = vec![];

    for (i, pos1) in positions.iter().enumerate() {
        for pos2 in &positions[i + 1..] {
            let dist = euclidean_distance(pos1, pos2);
            sorted_pairs_by_dist.push(Pair {
                dist: dist,
                pos1: pos1,
                pos2: pos2,
            });
        }
    }
    sorted_pairs_by_dist.sort_by(|pair1, pair2| pair1.dist.partial_cmp(&pair2.dist).unwrap());
    sorted_pairs_by_dist
}

fn read_input(file_path: &str) -> Vec<Pos> {
    let content = fs::read_to_string(file_path).expect("Could not read file");

    content
        .lines()
        .map(|s| {
            let numbers: Vec<i32> = s
                .split(",")
                .map(|n| n.parse::<i32>().expect("Could not parse number"))
                .collect();
            Pos {
                x: numbers[0],
                y: numbers[1],
                z: numbers[2],
            }
        })
        .collect()
}

fn euclidean_distance(pos1: &Pos, pos2: &Pos) -> f64 {
    let dx = (pos1.x - pos2.x) as i64;
    let dy = (pos1.y - pos2.y) as i64;
    let dz = (pos1.z - pos2.z) as i64;

    ((dx * dx + dy * dy + dz * dz) as f64).sqrt()
}
#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn part1_test() {
        let input = read_input("input/example.txt");
        let res = part1(&input, 10);
        assert_eq!(res, 40);
    }

    #[test]
    fn part2_test() {
        let input = read_input("input/example.txt");
        let res = part2(&input);
        assert_eq!(res, 25272);
    }
}
