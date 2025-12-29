use std::fs;

#[derive(Debug)]
struct Machine {
    light_diagram: u16, // binary representation of lights
    button_wiring: Vec<u16>, // binary rep of light toggles for buttons
    joltage: Vec<u32>,
}

fn main() {
    let machines = read_input("input/real.txt");
    let part1 = part1(&machines);
    println!("Day 10 part 1: {}", part1);
}

fn part1(machines: &Vec<Machine>) -> u32 {
    machines.iter().map(|m| min_presses(m) as u32).sum()
}

fn min_presses(machine: &Machine) -> u16 {
    let mut min_presses: u16 = u16::MAX;
    let buttons = machine.button_wiring.len();

    for button_combo in 0..(1u16 << buttons) {
        let mut light_state: u16 = 0;
        let mut button_presses = 0;

        for b in 0..buttons {
            if (button_combo >> b) & 1 == 1 {
                light_state ^= machine.button_wiring[b];
                button_presses += 1;
            }
        }

        if light_state == machine.light_diagram {
            min_presses = min_presses.min(button_presses);
        }
    }

    min_presses
}

fn read_input(file_path: &str) -> Vec<Machine> {
    let content = fs::read_to_string(file_path).expect("Could not read file");
    content
        .lines()
        .map(|s| {
            let split: Vec<&str> = s.split(" ").collect();

            let light_diagram = parse_light_diagram(&split[0]);
            let button_wiring: Vec<u16> = split[1..split.len() - 1]
                .iter()
                .map(|button| parse_button_wiring(button))
                .collect();

            let joltage = parse_joltage(&split[split.len() - 1]);

            Machine {
                light_diagram: light_diagram,
                button_wiring: button_wiring,
                joltage: joltage,
            }
        })
        .collect()
}

fn parse_light_diagram(s: &str) -> u16 {
    let inner = &s[1..s.len() - 1];

    let mut mask: u16 = 0;

    for (i, ch) in inner.chars().enumerate() {
        if ch == '#' {
            mask |= 1 << i;
        }
    }

    mask
}

fn parse_button_wiring(s: &str) -> u16 {
    let inner = &s[1..s.len() - 1];

    let mut mask: u16 = 0;

    if inner.is_empty() {
        return mask;
    }

    for idx in inner.split(',') {
        let i: usize = idx.parse().unwrap();
        mask |= 1 << i;
    }

    mask
}

fn parse_joltage(s: &str) -> Vec<u32> {
    s.trim_start_matches("{")
        .trim_end_matches("}")
        .split(",")
        .map(|n| n.parse::<u32>().unwrap())
        .collect()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn part1_test() {
        let machines = read_input("input/example.txt");
        let res = part1(&machines);
        println!("Machines: {:?}", machines);
        assert_eq!(res, 7);
    }
}
