use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "src/day1.txt";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);
    let mut sum = 0;

    for line in reader.lines() {
        let line = line.expect("Failed to read line");
        let mut fuel: u32 = line.trim().parse()
            .expect("Didn't read a number");

        loop {
            if fuel >= 6 {
                fuel = (fuel / 3) - 2;
                sum += fuel;
            } else {
                break;
            }
        }
    }

    println!("Sum: {}", sum);
}
