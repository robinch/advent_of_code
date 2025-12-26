use std::cmp::Reverse;
use std::fs;

fn main() {
    println!("Hello, world!");
    let red_tiles = read_input("input/real.txt");
    let part1 = part1(&red_tiles);
    println!("Part 1: {}", part1);

    let part2 = part2(&red_tiles);
    println!("Part 2: {}", part2);
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

fn part2(red_tiles: &Vec<(u64, u64)>) -> u64 {
    let mut boarders = vec![];

    for i in 0..red_tiles.len() {
        let start = red_tiles[i];
        let end = red_tiles[(i + 1) % red_tiles.len()];
        boarders.push((start, end));
    }

    let mut red_tiles_pair_by_area = vec![];

    for i in 0..red_tiles.len() {
        for j in i..red_tiles.len() {
            let col = red_tiles[i].0.abs_diff(red_tiles[j].0) + 1;
            let row = red_tiles[i].1.abs_diff(red_tiles[j].1) + 1;
            let area = col * row;

            red_tiles_pair_by_area.push((red_tiles[i], red_tiles[j], area));
        }
    }

    red_tiles_pair_by_area.sort_by_key(|(_, _, area)| Reverse(*area));

    for (tile1, tile2, area) in red_tiles_pair_by_area {
        let mut inside = true;
        for (start, end) in &boarders {
            let left_of_rect = start.0.max(end.0) <= tile1.0.min(tile2.0);
            let right_of_rect = start.0.min(end.0) >= tile1.0.max(tile2.0);
            let above_rect = start.1.max(end.1) <= tile1.1.min(tile2.1);
            let below_rect = start.1.min(end.1) >= tile1.1.max(tile2.1);

            if !(left_of_rect || right_of_rect || above_rect || below_rect) {
                inside = false;
                break;
            }
        }

        if inside {
            return area;
        }
    }

    panic!("You should have found an answer!")
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

    #[test]
    fn part2_test() {
        let red_tiles = read_input("input/example.txt");
        let res = part2(&red_tiles);
        assert_eq!(res, 24);
    }
}
