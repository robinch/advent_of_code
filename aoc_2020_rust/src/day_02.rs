use regex::Regex;
use std::fs;

#[derive(Debug)]
struct PasswordRow {
    min: i32,
    max: i32,
    policy_character: char,
    password: String,
}

pub fn solve(input_path: &str) -> (i32, i32) {
    let inputs = fs::read_to_string(input_path).expect("Could not read file");

    let mut passwords_that_follows_policy_1_counter = 0;
    let mut passwords_that_follows_policy_2_counter = 0;
    for input in inputs.lines() {
        let password_row = input_to_password_row(input);
        if follows_policy_1(&password_row) {
            passwords_that_follows_policy_1_counter += 1;
        }
        if follows_policy_2(&password_row) {
            passwords_that_follows_policy_2_counter += 1;
        }
    }

    (
        passwords_that_follows_policy_1_counter,
        passwords_that_follows_policy_2_counter,
    )
}

//This is slow
#[allow(dead_code)]
fn input_to_password_row_with_regex(input: &str) -> PasswordRow {
    let re = Regex::new(r"(\d+)-(\d+)\s(\w):\s(\w+)").unwrap();
    let cap = re.captures(input).unwrap();
    PasswordRow {
        min: cap[1].parse::<i32>().unwrap(),
        max: cap[2].parse::<i32>().unwrap(),
        policy_character: cap[3].chars().next().unwrap(),
        password: cap[4].to_string(),
    }
}

// This is faster
fn input_to_password_row(input: &str) -> PasswordRow {
    let split: Vec<&str> = input.split(" ").collect();
    let min_max: Vec<i32> = split[0]
        .split("-")
        .map(|s| s.parse::<i32>().unwrap())
        .collect();
    PasswordRow {
        min: min_max[0],
        max: min_max[1],
        policy_character: split[1].chars().next().unwrap(),
        password: split[2].to_string(),
    }
}

fn follows_policy_1(password_row: &PasswordRow) -> bool {
    let mut match_counter = 0;

    for c in password_row.password.chars() {
        if c == password_row.policy_character {
            match_counter += 1;
            if match_counter > password_row.max {
                return false;
            }
        }
    }

    match_counter >= password_row.min
}

fn follows_policy_2(password_row: &PasswordRow) -> bool {
    let pasword_chars: Vec<char> = password_row.password.chars().collect();
    (pasword_chars[(password_row.min - 1) as usize] == password_row.policy_character)
        ^ (pasword_chars[(password_row.max - 1) as usize] == password_row.policy_character)
}
