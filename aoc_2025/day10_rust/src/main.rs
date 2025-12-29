use std::fs;

#[derive(Debug)]
struct Machine {
    light_diagram: Vec<bool>,
    button_wiring: Vec<Vec<u32>>,
    joltage: Vec<u32>,
}

fn main() {
    println!("Hello, world!");
}

fn read_input(file_path: &str) -> Vec<Machine> {
    let content = fs::read_to_string(file_path).expect("Could not read file");
    content
        .lines()
        .map(|s| {
            let split: Vec<&str> = s.split(" ").collect();

            let light_diagram = parse_light_diagram(&split[0]);
            let button_wiring: Vec<Vec<u32>> = split[1..split.len() - 1]
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

fn parse_light_diagram(s: &str) -> Vec<bool> {
    s.trim_start_matches("[")
        .trim_end_matches("]")
        .chars()
        .map(|c| match c {
            '.' => false,
            '#' => true,
            x => panic!("unexpected char {}", x),
        })
        .collect()
}

fn parse_button_wiring(s: &str) -> Vec<u32> {
    s.trim_start_matches("(")
        .trim_end_matches(")")
        .split(",")
        .map(|n| n.parse::<u32>().unwrap())
        .collect()
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
    fn part1() {
        let machines = read_input("input/example.txt");
        println!("Machines: {:?}", machines);
        assert_eq!(0, 7);
    }
}
