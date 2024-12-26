use std::collections::HashMap;
use std::collections::HashSet;
use std::fs;

fn main() {
    part01();
    part02();
}

fn part01() {
    let (is_before_map, updates) = parse_input("input/input");

    let valid_updates: Vec<Vec<i32>> = updates
        .iter()
        .filter(|update| valid_order(&update, &is_before_map))
        .cloned()
        .collect();

    let sum_of_middle_pages: i32 = valid_updates
        .iter()
        .map(|update| update[update.len() / 2])
        .sum();

    println!("Day 5, part 1: : {}", sum_of_middle_pages);
}

fn part02() {
    let (is_before_map, updates) = parse_input("input/input");

    let sum_of_middle_pages: i32 = updates
        .into_iter()
        .filter(|update| !valid_order(&update, &is_before_map))
        .map(|mut update| {
            if !valid_order(&update, &is_before_map) {
                update.sort_by(|a, b| cmp_valid_order(a, b, &is_before_map))
            }

            update[update.len() / 2]
        })
        .sum();

    println!("Day 5, part 2: : {}", sum_of_middle_pages);
}

fn valid_order(update: &Vec<i32>, is_before_map: &HashMap<i32, Vec<i32>>) -> bool {
    let mut has_encountered: HashSet<&i32> = HashSet::new();

    for page in update {
        if let Some(before) = is_before_map.get(page) {
            for b in before {
                if has_encountered.contains(b) {
                    return false;
                }
            }
        }

        has_encountered.insert(page);
    }

    true
}

fn cmp_valid_order(a: &i32, b: &i32, is_before_map: &HashMap<i32, Vec<i32>>) -> std::cmp::Ordering {
    if a == b {
        return std::cmp::Ordering::Equal;
    }

    if let Some(before) = is_before_map.get(&a) {
        if before.contains(&b) {
            return std::cmp::Ordering::Less;
        }
    }

    std::cmp::Ordering::Greater
}

fn parse_input(path: &str) -> (HashMap<i32, Vec<i32>>, Vec<Vec<i32>>) {
    let input = fs::read_to_string(path).expect("Error reading file");
    let content = input.lines().map(|line| line.to_string()).collect();
    let (order_input, updates_input) = split_at_empty_string(content);
    let is_before_map = to_is_before_map(order_input);
    let updates = to_int_vec(updates_input);

    (is_before_map, updates)
}

fn to_int_vec(order: Vec<String>) -> Vec<Vec<i32>> {
    order
        .iter()
        .map(|line| line.split(",").map(|s| s.parse::<i32>().unwrap()).collect())
        .collect()
}

fn to_is_before_map(order: Vec<String>) -> HashMap<i32, Vec<i32>> {
    let mut order_map: HashMap<i32, Vec<i32>> = HashMap::new();

    order.iter().for_each(|line| {
        let o: Vec<i32> = line.split("|").map(|s| s.parse::<i32>().unwrap()).collect();
        order_map.entry(o[0]).or_insert(vec![]).push(o[1]);
    });

    order_map
}

fn split_at_empty_string(input: Vec<String>) -> (Vec<String>, Vec<String>) {
    if let Some(index) = input.iter().position(|line| line.is_empty()) {
        let order = input[..index].to_vec();
        let updates = input[index + 1..].to_vec();
        (order, updates)
    } else {
        (input.to_vec(), vec![])
    }
}
