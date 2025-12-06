use std::fs;

fn main() {
    let input = read_input("input/real.txt");
    let part1 = part1(&input);
    println!("Day03 part 1: {}", part1);
}

fn part1(batteries: &Vec<Vec<u8>>) -> u32 {
    batteries
        .iter()
        .map(|battery| max_joltage(battery) as u32)
        .sum()
}

fn max_joltage(battery: &Vec<u8>) -> u8 {
    let mut left_max_pos: usize = 0;
    let mut left_max: u8 = 0;
    for i in 0..battery.len() - 1 {
        let val = battery[i];
        if val > left_max {
            left_max = val;
            left_max_pos = i;
        }
    }

    let right_max = battery[left_max_pos + 1..]
        .iter()
        .max()
        .expect("Could not find right max");
    10 * left_max + right_max
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
        print!("input: {:?}", input);
        let result = part1(&input);
        assert_eq!(result, 357);
    }
}
