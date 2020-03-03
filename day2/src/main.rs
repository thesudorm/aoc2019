//use std::fs::File;
//use std::io::{BufRead, BufReader};

fn main() {

    //let filename = "input/day2.txt";
    //let file = File::open(filename).unwrap();

    let mut data = [1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50];
    let mut index = 0;

    loop {
        let opcode = data[index];

        if opcode == 1 { // Addition
            //println!("{} + {} goes into position {}", data[index + 1], data[index + 2], data[index + 3]);
            data[data[index + 3]] = data[data[index + 1]] + data[data[index + 2]];
            //println!("{} + {} goes into position {}", data[index + 1], data[index + 2], data[index + 3]);
        } else if opcode == 2 { // Multiplication
            data[data[index + 3]] = data[data[index + 1]] * data[data[index + 2]];
        } else if opcode == 99 {
            println!("Finished succesfully.");
            break;
        } else {
            println!("Task Failed");
            break;
        }

        index += 4;
    }

    println!("{:?}", data);
}
