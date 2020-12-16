use std::collections::HashSet;
use std::fs;

pub fn part_1(input_path: &str) -> i32 {
    let required_fields = required_fields();
    let input = fs::read_to_string(input_path).expect("Could not read file");
    let inputs: Vec<&str> = input.split("\n\n").map(|s| str::replace(s, "\n", ""));
    println!("{:?}", inputs);
    0
}

fn required_fields() -> HashSet<&'static str> {
    let required_fields: HashSet<&'static str> =
        ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"]
            .iter()
            .cloned()
            .collect();

    required_fields
}
