use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "src/day1.txt";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);
    let mut sum = 0;

    for line in reader.lines() {
        let line = line.expect("Failed to read line");
        let line: u32 = line.trim().parse()
            .expect("Didn't read a number");
        sum += (line / 3) - 2;
    }

    println!("Sum: {}", sum);
}
