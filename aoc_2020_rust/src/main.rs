mod day_01;

fn main() {
    day_01();
}

fn day_01() {
    let answer1 = day_01::part_1("./input/day_01.txt", 2020);
    println!("Day 1 Part 2: {}", answer1);

    let answer2 = day_01::part_2("./input/day_01.txt", 2020);
    println!("Day 1 Part 2: {}", answer2);
}
