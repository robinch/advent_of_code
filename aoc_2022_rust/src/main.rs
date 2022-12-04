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
    let answer1 = day_01::part_1("./input/day_01.txt");
    println!("Day 1 Part 1: {}", answer1);

    let answer2 = day_01::part_2("./input/day_01.txt", 3);
    println!("Day 1 Part 2: {}", answer2);
}

fn day_02() {
    let answer1 = day_02::part_1("./input/day_02.txt");
    println!("Day 2 Part 1: {}", answer1);

    let answer2 = day_02::part_2("./input/day_02.txt");
    println!("Day 2 Part 2: {}", answer2);
}

fn day_03() {
    let answer1 = day_03::part_1("./input/day_03.txt");
    println!("Day 3 Part 1: {}", answer1);

    let answer2 = day_03::part_2("./input/day_03.txt");
    println!("Day 3 Part 2: {}", answer2);
}

fn day_04() {
    let answer1 = day_04::part_1("./input/day_04.txt");
    println!("Day 4 Part 1: {}", answer1);

    let answer2 = day_04::part_2("./input/day_04.txt");
    println!("Day 4 Part 2: {}", answer2);
}
