use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "src/day1.txt";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);
    let mut total1 = 0;
    let mut total2 = 0;

    for line in reader.lines() {
        let line = line.expect("Failed to read line");
        let mass: usize = line.trim().parse()
            .expect("Didn't read a number");

        total1 += part_one(mass);
        total2 += part_two(mass);
    }

    println!("Total 1: {}   Total 2: {}", total1, total2);
}

fn part_one(x: usize) -> usize {
    match (x / 3).checked_sub(2) {
        Some(y) => y,
        None => 0,
    }
}

fn part_two(x: usize) -> usize {
    let mut total = 0;
    let mut x = x;
    loop {
        let y = match (x / 3).checked_sub(2) {
            Some(y) => y,
            None => {
                break total;
            }
        };
        println!("{}", y);
        x = y;
        total += y;
    }
}
