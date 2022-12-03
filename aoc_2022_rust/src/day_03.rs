use std::{collections::HashSet, fs};

pub fn part_1(file_path: &str) -> u32 {
    let backpacks = fs::read_to_string(file_path).expect("Could not read file");

    let mut duplicates: Vec<char> = vec![];

    for backpack in backpacks.lines() {
        duplicates.push(get_duplicate(&backpack));
    }

    let mut sum_of_priorities: u32 = 0;

    for dup in duplicates {
        sum_of_priorities += get_priority(&dup)
    }

    sum_of_priorities
}

pub fn part_2(file_path: &str) -> u32 {
    let inputs = fs::read_to_string(file_path).expect("Could not read file");

    let backpacks: Vec<&str> = inputs.lines().collect();

    let mut sum_of_priorities: u32 = 0;

    for i in 0..(backpacks.len() / 3) {
        let backpack_a = backpacks.get(i * 3).unwrap();
        let backpack_b = backpacks.get(i * 3 + 1).unwrap();
        let backpack_c = backpacks.get(i * 3 + 2).unwrap();

        let badge = get_badge(backpack_a, backpack_b, backpack_c);

        sum_of_priorities += get_priority(&badge);
    }

    sum_of_priorities
}

// item that exists in both compartments of a backpack
fn get_duplicate(backpack: &str) -> char {
    let (compartment_a, compartment_b) = backpack.split_at(backpack.len() / 2);
    let mut items_in_a = HashSet::new();

    let mut duplicate: Option<char> = None;

    for item_a in compartment_a.chars() {
        items_in_a.insert(item_a);
    }

    for item_b in compartment_b.chars() {
        if items_in_a.contains(&item_b) {
            duplicate = Some(item_b);
            break;
        }
    }

    duplicate.unwrap()
}

// badge exists in all three backpacks
fn get_badge(backpack_a: &str, backpack_b: &str, backpack_c: &str) -> char {
    let mut items_in_a = HashSet::new();
    let mut items_in_a_and_b = HashSet::new();

    let mut badge: Option<char> = None;

    for item_a in backpack_a.chars() {
        items_in_a.insert(item_a);
    }

    for item_b in backpack_b.chars() {
        if items_in_a.contains(&item_b) {
            items_in_a_and_b.insert(item_b);
        }
    }

    for item_c in backpack_c.chars() {
        if items_in_a_and_b.contains(&item_c) {
            badge = Some(item_c);
            break;
        }
    }

    badge.unwrap()
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
fn get_priority(item: &char) -> u32 {
    if item.is_lowercase() {
        (*item as u32) - ('a' as u32) + 1
    } else {
        (*item as u32) - ('A' as u32) + 27
    }
}
