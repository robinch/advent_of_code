use std::fs;

#[derive(Debug)]
struct Stacks {
    stacks: Vec<Vec<char>>,
}

impl Stacks {
    fn new(nr_of_stacks: usize) -> Stacks {
        let mut stacks: Vec<Vec<char>> = Vec::new();
        for _ in 0..nr_of_stacks {
            stacks.push(Vec::new());
        }

        Stacks { stacks: stacks }
    }

    fn push(&mut self, stack_nr: usize, value: char) {
        let _ = &self.stacks[stack_nr].push(value);
    }

    fn pop(&mut self, stack_nr: usize) -> char {
        match &self.stacks[stack_nr].pop() {
            Some(value) => *value,
            None => panic!("Nothing to pop!"),
        }
    }

    fn pop_many(&mut self, stack_nr: usize, nr_to_pop: usize) -> Vec<char> {
        let pop_from = &self.stacks[stack_nr].len() - nr_to_pop;
        let popped_values: Vec<char> = self.stacks[stack_nr].drain(pop_from..).collect();
        popped_values
    }

    fn push_many(&mut self, stack_nr: usize, values: &mut Vec<char>) {
        let _ = &self.stacks[stack_nr].append(values);
    }

    fn put_in_bottom(&mut self, stack_nr: usize, value: char) {
        let _ = &self.stacks[stack_nr].insert(0, value);
    }
}

struct Operation {
    move_amount: usize,
    from_stack: usize,
    to_stack: usize,
}

pub fn part_1(file_path: &str) -> String {
    let inputs = fs::read_to_string(file_path).expect("Could not read file!");

    let split_inputs: Vec<&str> = inputs.split("\n\n").collect();
    let mut stacks = create_stacks(split_inputs[0]);
    let operations = create_operations(split_inputs[1]);

    for op in operations {
        for _ in 0..op.move_amount {
            let val = stacks.pop(op.from_stack - 1);
            stacks.push(op.to_stack - 1, val);
        }
    }

    let mut top_crates = String::from("");

    for mut stack in stacks.stacks {
        match stack.pop() {
            Some(val) => top_crates.push(val),
            None => (),
        }
    }

    top_crates
}

pub fn part_2(file_path: &str) -> String {
    let inputs = fs::read_to_string(file_path).expect("Could not read file!");

    let split_inputs: Vec<&str> = inputs.split("\n\n").collect();
    let mut stacks = create_stacks(split_inputs[0]);
    let operations = create_operations(split_inputs[1]);

    for op in operations {
        let mut val = stacks.pop_many(op.from_stack - 1, op.move_amount);
        stacks.push_many(op.to_stack - 1, &mut val);
    }

    let mut top_crates = String::from("");

    for mut stack in stacks.stacks {
        match stack.pop() {
            Some(val) => top_crates.push(val),
            None => (),
        }
    }

    top_crates
}

fn create_stacks(input: &str) -> Stacks {
    let mut rows: Vec<&str> = input.split("\n").collect();

    // Remove the row with numbers
    rows.pop();

    // Every crate covers 4 chars (3 chars + 1 whitespace)
    // Last crate on each row covers 3 (3 chars, no whitespace)
    let nr_of_stacks: usize = ((rows[0].len() + 1) / 4) as usize;
    let mut stacks = Stacks::new(nr_of_stacks);

    for row in rows {
        let crates_on_row: Vec<&[u8]> = row.as_bytes().chunks(4).collect();
        for (i, c) in crates_on_row.iter().enumerate() {
            // the char on index 1 is a letter or whitespace
            if c[1].is_ascii_alphabetic() {
                stacks.put_in_bottom(i, c[1] as char);
            }
        }
    }

    stacks
}

fn create_operations(input: &str) -> Vec<Operation> {
    let mut operations: Vec<Operation> = Vec::new();
    let rows: Vec<&str> = input.split("\n").collect();
    for row in rows {
        let split_row: Vec<&str> = row.split(" ").collect();
        operations.push(Operation {
            move_amount: split_row[1].parse::<usize>().unwrap(),
            from_stack: split_row[3].parse::<usize>().unwrap(),
            to_stack: split_row[5].parse::<usize>().unwrap(),
        })
    }

    operations
}
