use std::collections::HashMap;
use std::collections::HashSet;
use std::fs;

fn main() {
    let input = read_input("input/real.txt");
    let part1 = part1(&input);
    println!("Day 07, part1: {}", part1);

    let part2 = part2(&input);
    println!("Day 07, part2: {}", part2);
}

fn part1(input: &Vec<Vec<char>>) -> i32 {
    let mut splits = 0;
    let mut active_beams: HashSet<usize> = HashSet::new();
    for col in 0..input[0].len() {
        if input[0][col] == 'S' {
            active_beams.insert(col);
        }
    }

    for row in 0..input.len() - 1 {
        let beams: Vec<usize> = active_beams.iter().cloned().collect();
        beams.iter().for_each(|beam_col| {
            if input[row + 1][*beam_col] == '^' {
                splits += 1;
                active_beams.remove(beam_col);
                active_beams.insert(beam_col - 1);
                active_beams.insert(beam_col + 1);
            }
        });
    }

    splits
}

fn part2(input: &Vec<Vec<char>>) -> i64 {
    let mut splits: HashMap<(usize, usize), i64> = HashMap::new();

    let (start, _) = input[0]
        .iter()
        .enumerate()
        .find(|(_, c)| **c == 'S')
        .expect("Could not find start");

    calc_splits(&input, 1, start, &mut splits)
}

fn calc_splits(
    input: &Vec<Vec<char>>,
    row: usize,
    col: usize,
    splits: &mut HashMap<(usize, usize), i64>,
) -> i64 {
    if row == input.len() - 1 {
        return 1;
    }

    if input[row][col] == '.' {
        return calc_splits(&input, row + 1, col, splits);
    }

    match splits.get(&(row, col)) {
        Some(splits) => *splits,
        None => {
            let nr_of_splits = calc_splits(&input, row + 1, col - 1, splits)
                + calc_splits(&input, row + 1, col + 1, splits);
            splits.insert((row, col), nr_of_splits);
            nr_of_splits
        }
    }
}

fn read_input(file_path: &str) -> Vec<Vec<char>> {
    let content = fs::read_to_string(file_path).expect("Could not read file");
    content.lines().map(|s| s.chars().collect()).collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    pub fn part1_test() {
        let input = read_input("input/example.txt");
        let res = part1(&input);
        assert_eq!(res, 21);
    }

    #[test]
    pub fn part2_test() {
        let input = read_input("input/example.txt");
        let res = part2(&input);
        assert_eq!(res, 40);
    }
}
