use std::fs;

fn main() {
    let input = read_input("input/real.txt");
    let part1 = part1(&input);
    let part2 = part2(&input);
    println!("Day 2 part 1: {}", part1);
    println!("Day 2 part 2: {}", part2);
}

fn part1(input: &Vec<(i64, i64)>) -> i64 {
    let mut sum_invalid = 0;

    for (from, to) in input {
        for n in *from..=*to {
            if is_invalid(n) {
                sum_invalid += n;
            }
        }
    }

    return sum_invalid;
}

fn part2(input: &Vec<(i64, i64)>) -> i64 {
    let mut sum_invalid = 0;

    for (from, to) in input {
        for n in *from..=*to {
            if is_repeating(n) {
                sum_invalid += n;
            }
        }
    }

    return sum_invalid;
}

fn is_invalid(n: i64) -> bool {
    let s = n.to_string();
    let number = s.as_bytes();

    if number.len() % 2 == 1 {
        return false;
    }

    let half = number.len() / 2;

    for i in 0..half {
        if number[i] != number[i + half] {
            return false;
        }
    }

    true
}

fn is_repeating(n: i64) -> bool {
    let digits: Vec<u8> = n.to_string().bytes().map(|b| b - b'0').collect();

    let len = digits.len();

    for pattern_len in 1..=len / 2 {
        if len % pattern_len != 0 {
            continue;
        }

        let reps = len / pattern_len;
        let pattern = &digits[0..pattern_len];

        let mut is_repeating = true;

        for chunk in digits.chunks_exact(pattern_len).skip(1).take(reps - 1) {
            if chunk != pattern {
                is_repeating = false;
                break;
            }
        }

        if is_repeating {
            return true;
        }
    }

    false
}

fn read_input(file_path: &str) -> Vec<(i64, i64)> {
    let contents =
        fs::read_to_string(file_path).expect(&format!("Could not read file {}", file_path));
    let inputs: Vec<_> = contents
        .lines()
        .flat_map(|s| s.split(","))
        .map(|s| {
            let mut split = s.split("-");
            let s_from = split.next().expect("Could not get next in split");
            let s_to = split.next().expect("Could not get next in split");
            let from = s_from.parse::<i64>().expect("Could not parse number");
            let to = s_to.parse::<i64>().expect("Could not parse number");

            (from, to)
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
        assert_eq!(result, 1227775554);
    }

    #[test]
    fn part2_test() {
        let input = read_input("input/example.txt");
        let result = part2(&input);
        assert_eq!(result, 4174379265);
    }
}
