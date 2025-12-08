use std::fs;

fn main() {
    println!("Hello, world!");
    let (numbers, operators) = read_input_for_part1("input/real.txt");
    let part1 = part1(numbers, operators);
    println!("Day06 part 1: {}", part1);

    let (numbers2, operators2) = read_input_for_part2("input/real.txt");
    let part2 = part2(numbers2, operators2);
    println!("Day06 part 2: {}", part2);
}

fn part1(numbers: Vec<Vec<i32>>, operators: Vec<char>) -> i64 {
    let mut result: i64 = 0;

    for i in 0..operators.len() {
        let mut res: i64 = if operators[i] == '+' { 0 } else { 1 };
        for j in 0..numbers.len() {
            let n = numbers[j][i] as i64;
            if operators[i] == '+' {
                res += n
            } else {
                res *= n
            };
        }

        result += res;
    }

    result
}

fn part2(numbers: Vec<Vec<i32>>, operators: Vec<char>) -> i64 {
    let mut result = 0;

    for row in 0..numbers.len() {
        let res: i64 = if operators[row] == '+' {
            numbers[row].iter().map(|n| *n as i64).sum()
        } else {
            numbers[row].iter().map(|n| *n as i64).product()
        };

        result += res;
    }

    result
}

fn read_input_for_part1(file_path: &str) -> (Vec<Vec<i32>>, Vec<char>) {
    let content = fs::read_to_string(file_path).expect("Could not read file");

    let input: Vec<Vec<String>> = content
        .lines()
        .filter(|s| !s.is_empty())
        .map(|s| {
            s.split(' ')
                .filter(|x| !x.is_empty())
                .map(|x| x.to_string())
                .collect()
        })
        .collect();

    let numbers: Vec<Vec<i32>> = input[..input.len() - 1]
        .iter()
        .map(|numbers| numbers.iter().map(|n| n.parse::<i32>().unwrap()).collect())
        .collect();

    let operators: Vec<char> = input
        .iter()
        .last()
        .unwrap()
        .iter()
        .map(|s| s.chars().next().unwrap())
        .collect();

    (numbers, operators)
}

fn read_input_for_part2(file_path: &str) -> (Vec<Vec<i32>>, Vec<char>) {
    // fn read_input2(file_path: &str) -> (Vec<Vec<char>>, Vec<char>, Vec<usize>) {
    let content = fs::read_to_string(file_path).expect("Could not read file");

    let input: Vec<String> = content
        .lines()
        .filter(|s| !s.is_empty())
        .map(|s| s.to_string())
        .collect();

    let char_splitted: Vec<Vec<char>> = input[..input.len() - 1]
        .iter()
        .map(|s| s.chars().collect())
        .collect();

    let mut numbers: Vec<Vec<i32>> = vec![];
    let mut col_numbers: Vec<i32> = vec![];

    let rows = char_splitted.len();
    let cols = char_splitted[1].len();

    for c in 0..cols {
        let mut col: Vec<char> = vec![];

        for r in 0..rows {
            col.push(char_splitted[r][c]);
        }

        if col.iter().all(|ch| *ch == ' ') {
            numbers.push(col_numbers);
            col_numbers = vec![];
        } else {
            let n: i32 = col
                .iter()
                .filter(|ch| **ch != ' ')
                .collect::<String>()
                .parse()
                .unwrap();
            col_numbers.push(n);
        }
    }

    numbers.push(col_numbers);

    let operators: Vec<char> = input
        .iter()
        .last()
        .unwrap()
        .chars()
        .filter(|c| *c != ' ')
        .collect();

    (numbers, operators)
}
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        let (numbers, operators) = read_input_for_part1("input/example.txt");
        let result = part1(numbers, operators);
        assert_eq!(result, 4277556)
    }

    #[test]
    fn part2_test() {
        let (numbers, operators) = read_input_for_part2("input/example.txt");
        let result = part2(numbers, operators);
        assert_eq!(result, 3263827)
    }
}
