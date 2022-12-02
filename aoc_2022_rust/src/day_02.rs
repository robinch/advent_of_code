use std::fs;

enum Hand {
    Rock,
    Paper,
    Scissor,
}

enum Outcome {
    Win,
    Draw,
    Lose,
}

pub fn part_1(file_path: &str) -> u32 {
    let inputs = fs::read_to_string(file_path).expect("Could not read file");

    let mut tot_points: u32 = 0;

    for input in inputs.lines() {
        let split_input: Vec<&str> = input.split(" ").collect();

        let opponents_hand = input_to_hand(split_input[0]);
        let your_hand = input_to_hand(split_input[1]);
        let outcome = outcome(&opponents_hand, &your_hand);
        tot_points += outcome_to_points(outcome) + hand_to_points(your_hand);
    }

    tot_points
}

pub fn part_2(file_path: &str) -> u32 {
    let inputs = fs::read_to_string(file_path).expect("Could not read file");

    let mut tot_points: u32 = 0;

    for input in inputs.lines() {
        let split_input: Vec<&str> = input.split(" ").collect();

        let opponents_hand = input_to_hand(split_input[0]);
        let outcome = input_to_outcome(split_input[1]);
        let your_hand = outcome_to_hand(&opponents_hand, &outcome);
        tot_points += outcome_to_points(outcome) + hand_to_points(your_hand);
    }

    tot_points
}

fn input_to_hand(input: &str) -> Hand {
    match input {
        "A" | "X" => Hand::Rock,
        "B" | "Y" => Hand::Paper,
        "C" | "Z" => Hand::Scissor,
        input => panic!("Could not convert input '{}' to Hand", input),
    }
}

fn input_to_outcome(input: &str) -> Outcome {
    match input {
        "X" => Outcome::Lose,
        "Y" => Outcome::Draw,
        "Z" => Outcome::Win,
        input => panic!("Could not convert input '{}' to Outcome", input),
    }
}

fn outcome(opponent: &Hand, you: &Hand) -> Outcome {
    match (opponent, you) {
        (Hand::Rock, Hand::Rock) => Outcome::Draw,
        (Hand::Rock, Hand::Paper) => Outcome::Win,
        (Hand::Rock, Hand::Scissor) => Outcome::Lose,
        (Hand::Paper, Hand::Rock) => Outcome::Lose,
        (Hand::Paper, Hand::Paper) => Outcome::Draw,
        (Hand::Paper, Hand::Scissor) => Outcome::Win,
        (Hand::Scissor, Hand::Rock) => Outcome::Win,
        (Hand::Scissor, Hand::Paper) => Outcome::Lose,
        (Hand::Scissor, Hand::Scissor) => Outcome::Draw,
    }
}

fn outcome_to_hand(opponent: &Hand, outcome: &Outcome) -> Hand {
    match (opponent, outcome) {
        (Hand::Rock, Outcome::Win) => Hand::Paper,
        (Hand::Rock, Outcome::Draw) => Hand::Rock,
        (Hand::Rock, Outcome::Lose) => Hand::Scissor,
        (Hand::Paper, Outcome::Win) => Hand::Scissor,
        (Hand::Paper, Outcome::Draw) => Hand::Paper,
        (Hand::Paper, Outcome::Lose) => Hand::Rock,
        (Hand::Scissor, Outcome::Win) => Hand::Rock,
        (Hand::Scissor, Outcome::Draw) => Hand::Scissor,
        (Hand::Scissor, Outcome::Lose) => Hand::Paper,
    }
}

fn outcome_to_points(outcome: Outcome) -> u32 {
    match outcome {
        Outcome::Win => 6,
        Outcome::Draw => 3,
        Outcome::Lose => 0,
    }
}

fn hand_to_points(hand: Hand) -> u32 {
    match hand {
        Hand::Rock => 1,
        Hand::Paper => 2,
        Hand::Scissor => 3,
    }
}
