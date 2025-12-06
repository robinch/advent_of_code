use std::fs;

fn main() {
    let input = read_input("input/real.txt");
    let part1 = part1(&input);
    println!("Day03 part 1: {}", part1);

    let part2 = part2(&input);
    println!("Day03 part 2: {}", part2);
}

fn part1(banks: &Vec<Vec<u8>>) -> u32 {
    banks
        .iter()
        .map(|battery_bank| {
            let (max_left_index, max_left) =
                max_val_and_index(&battery_bank, 0, battery_bank.len() - 1);
            let (_, max_right) =
                max_val_and_index(&battery_bank, max_left_index + 1, battery_bank.len());

            (10 * max_left + max_right) as u32
        })
        .sum()
}

fn part2(banks: &Vec<Vec<u8>>) -> u64 {
    banks
        .iter()
        .map(|battery_bank| {
            let mut sum: u64 = 0;
            let mut max_index: isize = -1;
            let len = battery_bank.len();
            for i in (0..=11).rev() {
                let (index, max) =
                    max_val_and_index(&battery_bank, (max_index + 1) as usize, len - i);
                sum += (max as u64) * u64::pow(10, i as u32);
                max_index = index as isize;
            }
            sum
        })
        .sum()
}

fn max_val_and_index(battery: &Vec<u8>, from: usize, to: usize) -> (usize, u8) {
    let mut max_pos: usize = 0;
    let mut max: u8 = 0;
    for i in from..to {
        let val = battery[i];
        if val > max {
            max = val;
            max_pos = i;
        }
    }

    (max_pos, max)
}

fn read_input(file_path: &str) -> Vec<Vec<u8>> {
    let contents =
        fs::read_to_string(file_path).expect(&format!("Could not read file {}", file_path));
    let inputs: Vec<_> = contents
        .lines()
        .map(|s| {
            let digits: Vec<u8> = s.bytes().map(|b| b - b'0').collect();
            digits
        })
        .collect();

    inputs
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        let input = read_input("input/example.txt");
        let result = part1(&input);
        assert_eq!(result, 357);
    }

    #[test]
    fn part2_test() {
        let input = read_input("input/example.txt");
        let result = part2(&input);
        assert_eq!(result, 3121910778619);
    }
}
