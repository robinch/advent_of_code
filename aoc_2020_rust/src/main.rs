mod day_01;
mod day_02;
mod day_03;
mod day_04;

fn main() {
    day_01();
    day_02();
    day_03();
    day_04();
}

fn day_01() {
    let answer1 = day_01::part_1("./input/day_01.txt", 2020);
    println!("Day 1 Part 2: {}", answer1);

    let answer2 = day_01::part_2("./input/day_01.txt", 2020);
    println!("Day 1 Part 2: {}", answer2);
}

fn day_02() {
    let (answer1, answer2) = day_02::solve("./input/day_02.txt");
    println!("Day 2 Part 1: {}", answer1);
    println!("Day 2 Part 2: {}", answer2);
}

fn day_03() {
    let answer_test = day_03::part_1("./input/day_03_test.txt", 3, 1);
    let answer1 = day_03::part_1("./input/day_03.txt", 3, 1);
    let mut steps: Vec<(i32, i32)> = Vec::new();
    steps.push((1, 1));
    steps.push((3, 1));
    steps.push((5, 1));
    steps.push((7, 1));
    steps.push((1, 2));
    let answer2 = day_03::part_2("./input/day_03.txt", steps);
    println!("Day 3 Part 1 Test: {}", answer_test);
    println!("Day 3 Part 1: {}", answer1);
    println!("Day 3 Part 2: {}", answer2);
}

fn day_04() {
    let answer_test = day_04::part_1("./input/day_04_test.txt");
    println!("Day 4 Part 1 Test: {}", answer_test);
}
