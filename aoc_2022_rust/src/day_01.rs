use std::fs;

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
pub fn part_1(input_path: &str) -> u32 {
    let inputs = fs::read_to_string(input_path).expect("Could not read file");

    let mut max: u32 = 0;
    let mut local_max: u32 = 0;

    for calorie in inputs.lines() {
        if calorie == "" {
            if local_max > max {
                max = local_max;
            }

            local_max = 0;
        } else {
            local_max += calorie.parse::<u32>().unwrap();
        }
    }

    max
}

pub fn part_2(input_path: &str, nr_of_top_elves: u32) -> u32 {
    let inputs = fs::read_to_string(input_path).expect("Could not read file");

    let mut calories = Vec::new();
    let mut calorie_count: u32 = 0;

    for calorie in inputs.lines() {
        if calorie == "" {
            calories.push(calorie_count);

            calorie_count = 0;
        } else {
            calorie_count += calorie.parse::<u32>().unwrap();
        }
    }

    calories.push(calorie_count);

    calories.sort();
    calories.reverse();

    let mut calorie_sum = 0;

    for i in 0..nr_of_top_elves as usize {
        calorie_sum += calories[i];
    }

    calorie_sum
}
