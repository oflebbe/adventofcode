use std::fs;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    let buf = fs::read_to_string(&args[1]).expect("reading");
    let lines : Vec<&str> = buf.split("\n").collect();
    let mut col1 = Vec::<u32>::new();
    let mut col2 = Vec::<u32>::new();
    for line in lines {
        let toks : Vec<u32> = line.split_whitespace().map( |x| x.parse::<u32>().unwrap()).collect();
        if toks.len() == 0 {
            break
        }
        col1.push( toks[0]);
        col2.push( toks[1]);
    }
    col1.sort();
    col2.sort();
    let mut sum = 0;
    for i in 0..col1.len() {
        sum += (col1[i] as i32 - col2[i]as i32).abs();
    }
    println!("Task 1 {}", sum);

    let mut last_score = 0;
    let mut last_value = 0;
    let mut ptr2 = 0;
    let mut total_score = 0;
    for v in col1 {
        if v == last_value {
            total_score += last_score;
            continue;
        }
        if ptr2 >= col2.len() {
            break
        }
        
        while v > col2[ptr2] {
            ptr2+=1;
            if ptr2 >= col2.len() {
                break
            }
        }
        if ptr2 >= col2.len() {
            break
        }
        let mut score = 0;
        while v == col2[ptr2] {
            score += 1;
            ptr2 +=1;
            if ptr2 > col2.len() {
                break
            }
        }
        total_score += score * v;
        last_score = score*v;
        last_value = v;
    }

    println!("Task2: {}", total_score);
}
