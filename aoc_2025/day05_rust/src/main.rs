use std::fs;

fn main() {
    println!("Hello, world!");
    let input = read_input("input/real.txt");
    let part1 = part1(&input);
    println!("Day05 part 1: {}", part1);

    let part2 = part2(&input.0);
    println!("Day05 part 2: {}", part2);
}

fn part1(input: &(Vec<(i64, i64)>, Vec<i64>)) -> usize {
    let (ranges, ids) = input;

    ids.iter()
        .filter(|id| ranges.iter().any(|(from, to)| *id >= from && *id <= to))
        .count()
}

fn part2(ranges: &Vec<(i64, i64)>) -> i64 {
    let ranges_clone = ranges.clone();
    let merged_ranges = merge_ranges(ranges_clone);

    merged_ranges.iter().map(|(from, to)| to - from + 1).sum()
}

fn merge_ranges(mut ranges: Vec<(i64, i64)>) -> Vec<(i64, i64)> {
    ranges.sort_by_key(|(from, _)| *from);

    let mut merged = vec![];

    ranges
        .iter()
        .for_each(|(from, to)| match merged.last_mut() {
            Some((_last_from, last_to)) => {
                if from <= last_to {
                    *last_to = (*last_to).max(*to);
                } else {
                    merged.push((*from, *to))
                }
            }
            None => merged.push((*from, *to)),
        });

    merged
}

fn read_input(file_path: &str) -> (Vec<(i64, i64)>, Vec<i64>) {
    let content =
        fs::read_to_string(file_path).expect(&format!("Could not read file {}", file_path));

    let mut split = content.split("\n\n");
    let ranges_string = split.next().expect("Can't get ranges");
    let ids_string = split.next().expect("Can't get ids");

    let ranges = parse_ranges(ranges_string);
    let ids = parse_ids(ids_string);

    (ranges, ids)
}

fn parse_ranges(ranges_string: &str) -> Vec<(i64, i64)> {
    ranges_string
        .lines()
        .flat_map(|s| s.split(","))
        .map(|s| {
            let mut split = s.split("-");
            let s_from = split.next().expect("Could not get next in split");
            let s_to = split.next().expect("Could not get next in split");
            let from = s_from.parse::<i64>().expect("Could not parse number");
            let to = s_to.parse::<i64>().expect("Could not parse number");

            (from, to)
        })
        .collect()
}

fn parse_ids(ids_string: &str) -> Vec<i64> {
    ids_string
        .lines()
        .map(|s| s.parse::<i64>().expect("Could not parse id"))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        let input = read_input("input/example.txt");
        let result = part1(&input);
        assert_eq!(result, 3)
    }

    #[test]
    fn part2_test() {
        let (ranges, _ids) = read_input("input/example.txt");
        let result = part2(&ranges);
        assert_eq!(result, 14)
    }
}
