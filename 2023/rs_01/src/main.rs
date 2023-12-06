use std::fs;

const DIGITS: [&str; 9] = ["1", "2", "3", "4", "5", "6", "7", "8", "9"];
const NUMBERS: [&str; 9] = [
    "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
];

fn eval1(line: &str) -> usize {
    let mut first_i = 10;
    let mut last_i = 10;
    let mut first_pos = line.len();
    let mut last_pos = 0;

    for i in 0..9 {
        if let Some(found) = line.find(DIGITS[i]) {
            if first_pos > found {
                first_pos = found;
                first_i = i + 1;
            }
        }
        if let Some(found) = line.rfind(DIGITS[i]) {
            if last_pos <= found {
                last_pos = found;
                last_i = i + 1;
            }
        }
    }

    first_i * 10 + last_i
}

fn eval2(line: &str) -> usize {
    let mut first_i = 10;
    let mut last_i = 10;
    let mut first_pos = line.len();
    let mut last_pos = 0;

    for i in 0..9 {
        if let Some(found) = line.find(DIGITS[i]) {
            if first_pos > found {
                first_pos = found;
                first_i = i + 1;
            }
        }
        if let Some(found) = line.find(NUMBERS[i]) {
            if first_pos > found {
                first_pos = found;
                first_i = i + 1;
            }
        }
        if let Some(found) = line.rfind(DIGITS[i]) {
            if last_pos <= found {
                last_pos = found;
                last_i = i + 1;
            }
        }
        if let Some(found) = line.rfind(NUMBERS[i]) {
            if last_pos <= found {
                last_pos = found;
                last_i = i + 1;
            }
        }
    }
    first_i * 10 + last_i
}

fn eval1_file( file_name: &str)-> usize {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let mut sum = 0;    
    for line in contents.split("\n") {
        if line.len() < 1 {
            continue;
        }
        sum += eval1( line);
    }
    sum
}

fn eval2_file( file_name: &str)-> usize {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let mut sum = 0;    
    for line in contents.split("\n") {
        if line.len() < 1 {
            continue;
        }
        sum += eval2( line);
    }
    sum
}


#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn simple() {
        assert_eq!( eval1("pqr3stu8vwx"), 38);
    
        assert_eq!( eval1("treb7uchet"), 77);
    }

    #[test]
    fn simple2() {
        assert_eq!( eval2("zoneight234"), 14);
    
        assert_eq!( eval2("7pqrstsixteen"), 76);
    }

    #[test]
    fn part1() {
        assert_eq!( eval1_file("test1.txt"), 142);
    }


    #[test]

    fn part2() {
        assert_eq!( eval2_file("test2.txt"), 281);
    }
}


fn main() {

    println!("part1: {}", eval1_file("input.txt"));
    println!("part2: {}", eval2_file("input.txt"));

}
