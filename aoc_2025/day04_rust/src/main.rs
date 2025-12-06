use std::collections::HashSet;
use std::fs;

fn main() {
    let input = read_input("input/real.txt");
    let part1 = part1(&input);
    println!("Day04 Part 1: {}", part1);

    let part2 = part2(&input);
    println!("Day04 Part 2: {}", part2);
}

fn part1(input: &Vec<Vec<char>>) -> i32 {
    let (count, _) = papers_that_can_be_removed(&input);
    count
}

fn part2(input: &Vec<Vec<char>>) -> i32 {
    let mut map = input.clone();
    let mut total_papers_removed = 0;

    loop {
        match papers_that_can_be_removed(&map) {
            (0, _) => break,
            (count, can_remove) => {
                total_papers_removed += count;

                can_remove.iter().for_each(|(row, col)| {
                    map[*row][*col] = '.';
                })
            }
        }
    }

    total_papers_removed
}

fn papers_that_can_be_removed(input: &Vec<Vec<char>>) -> (i32, Vec<(usize, usize)>) {
    let mut paper_that_can_be_accessed: HashSet<(usize, usize)> = HashSet::new();

    for row in 0..input.len() {
        for col in 0..input[1].len() {
            if input[row][col] == '@' {
                if nr_of_papers_adjacent(input, row, col) < 4 {
                    paper_that_can_be_accessed.insert((row, col));
                }
            }
        }
    }

    let papers: Vec<_> = paper_that_can_be_accessed
        .iter()
        .map(|coord| *coord)
        .collect();

    (papers.len() as i32, papers)
}

fn nr_of_papers_adjacent(input: &Vec<Vec<char>>, row: usize, col: usize) -> usize {
    adjacent_coords(row, col, input.len(), input[1].len())
        .iter()
        .filter(|(arow, acol)| input[*arow][*acol] == '@')
        .count()
}

fn adjacent_coords(row: usize, col: usize, nr_rows: usize, nr_cols: usize) -> Vec<(usize, usize)> {
    let irow = row as isize;
    let icol = col as isize;
    let rows = nr_rows as isize;
    let cols = nr_cols as isize;
    let adjacent_coords: Vec<(isize, isize)> = vec![
        (irow - 1, icol - 1),
        (irow - 1, icol),
        (irow - 1, icol + 1),
        (irow, icol - 1),
        (irow, icol + 1),
        (irow + 1, icol - 1),
        (irow + 1, icol),
        (irow + 1, icol + 1),
    ];

    adjacent_coords
        .iter()
        .filter_map(|(arow, acol)| {
            if *arow >= 0 && *arow < rows && *acol >= 0 && *acol < cols {
                Some((*arow as usize, *acol as usize))
            } else {
                None
            }
        })
        .collect()
}

fn read_input(file_path: &str) -> Vec<Vec<char>> {
    fs::read_to_string(file_path)
        .expect("could not read file!")
        .split("\n")
        .filter(|s| !s.is_empty())
        .map(|s| {
            s.chars()
                .map(|c| match c {
                    'x' => '@',
                    '@' => '@',
                    '.' => '.',
                    unexpected => panic!("Unexpected char {}", unexpected),
                })
                .collect()
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        let input = read_input("input/example.txt");
        let result = part1(&input);
        assert_eq!(result, 13);
    }

    #[test]
    fn part2_test() {
        let input = read_input("input/example.txt");
        let result = part2(&input);
        assert_eq!(result, 43);
    }
}
