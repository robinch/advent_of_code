use std::fs;

fn main() {
    println!("Hello, world!");
    let red_tiles = read_input("input/real.txt");
    let part1 = part1(&red_tiles);
    println!("Part 1: {}", part1);
}

fn part1(red_tiles: &Vec<(u64, u64)>) -> u64 {
    let mut max_area = 0;
    for i in 0..red_tiles.len() {
        for j in i..red_tiles.len() {
            let col = red_tiles[i].0.abs_diff(red_tiles[j].0) + 1;
            let row = red_tiles[i].1.abs_diff(red_tiles[j].1) + 1;
            let area = col * row;
            max_area = max_area.max(area);
        }
    }
    max_area
}

fn read_input(file_path: &str) -> Vec<(u64, u64)> {
    let content = fs::read_to_string(file_path).unwrap();

    content
        .lines()
        .map(|s| {
            let mut split = s.split(",");
            let col = split.next().unwrap().parse::<u64>().unwrap();
            let row = split.next().unwrap().parse::<u64>().unwrap();
            (col, row)
        })
        .collect()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn part1_test() {
        let red_tiles = read_input("input/example.txt");
        let res = part1(&red_tiles);
        println!("part 1: {}", res);
        assert_eq!(res, 50);
    }
}
