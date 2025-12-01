use std::fs;

fn main() {
    let input = read_input("input/real.txt");
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &Vec<(String, i32)>) -> i32 {
    let mut on_zero = 0;
    let mut dial = 50;

    for (direction, steps) in input {
        if direction == "R" {
            dial = (dial + steps).rem_euclid(100);
        } else {
            dial = (dial - steps).rem_euclid(100);
        }

        if dial == 0 {
            on_zero += 1;
        }
    }

    on_zero
}

fn part2(input: &Vec<(String, i32)>) -> i32 {
    let mut passed_zero = 0;
    let mut dial = 50;

    for (direction, steps) in input {
        let new_dial: i32;
        if direction == "R" {
            new_dial = dial + steps;
            if new_dial > 99 {
                passed_zero += new_dial / 100
            }
        } else {
            new_dial = dial - steps;

            if new_dial <= 0 {
                passed_zero += new_dial / -100;

                if dial != 0 {
                    passed_zero += 1;
                }
            }
        }

        dial = (new_dial).rem_euclid(100);
    }

    passed_zero
}

fn read_input(file_path: &str) -> Vec<(String, i32)> {
    let contents =
        fs::read_to_string(file_path).expect(&format!("Could not read file {}", file_path));
    let inputs: Vec<_> = contents
        .lines()
        .map(|s| {
            let (direction, number) = s.split_at(1);
            let steps = number.parse::<i32>().expect("Could not parse number");

            (direction.to_string(), steps)
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
        assert_eq!(result, 3);
    }

    #[test]
    fn part2_test() {
        let input = read_input("input/example.txt");
        let result = part2(&input);
        assert_eq!(result, 6);
    }
}
