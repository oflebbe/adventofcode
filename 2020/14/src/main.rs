use std::collections::HashMap;
use regex::Regex;
use std::iter::Iterator;
use std::vec::Vec;



fn solve2( input : &str) -> usize {
    let mut orvalue = 0usize;
    let mut andvalue = 0usize;
    let mut memory = HashMap::new();
    let re = Regex::new("mem\\[([0-9]*)\\] = ([01X]*)").unwrap();

    for line in input.lines() {
        if line.starts_with( "mask = ") {
            let rest = line.strip_prefix("mask = ").unwrap();
        } else {
            let caps = re.captures(line).unwrap();
            let key = caps.get(1).map_or("", |m| m.as_str());
            let val : usize = caps.get(2).map_or("", |m| m.as_str()).parse().unwrap();
            val |= orvalue;
            val &= andvalue;
            memory.insert(key, val);
        }
    }

    let mut sum = 0;
    for (_, value) in &memory {
        sum += value;
    } 
    sum
}

fn solve1( input : &str) -> usize {
    let mut orvalue = 0usize;
    let mut andvalue = 0usize;
    let mut memory = HashMap::new();
    for line in input.lines() {
        if line.starts_with( "mask = ") {
            let rest = line.strip_prefix("mask = ").unwrap();
            
        } else {
            let tok : Vec<&str> = line.split(" = ").collect();
            let mut val : usize = tok[1].parse().unwrap();
            let key = tok[0];
            val |= orvalue;
            val &= andvalue;
            memory.insert(key, val);
        }
    }

    let mut sum = 0;
    for (_, value) in &memory {
        sum += value;
    } 
    sum
}
fn main() {
   let input = include_str!("input.txt");
   println!("{}", solve( input));
}
