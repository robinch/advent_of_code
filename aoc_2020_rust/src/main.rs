mod day_01;
mod day_02;

fn main() {
    day_01();
    day_02();
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
