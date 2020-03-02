use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "day1.txt";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    for line in reader.lines() {
        println!("{}", line);
    }
}
