
use std::collections::HashMap;
use std::fs;
use std::cmp;


// or null if nor possible
fn possible( line: &str, limit : &HashMap<String,i32>) -> i32 {
    let toks : Vec<&str> = line.split(": ").collect();
    let tries: Vec<&str> = toks[1].split("; ").collect();
    for t in tries {
        let dices : Vec<&str> = t.split(", ").collect();
        for d in dices {
            let one : Vec<&str> = d.split(" ").collect();
            if one.len() != 2{
                println!("{}|{}|{}|{}", line, t, d, one.len());
                panic!("bummer");
            }
            if limit[ one[1]] < one[0].parse::<i32>().unwrap() {
                return 0
            }
        }
    }

   let nummer : Vec<&str> =  toks[0].split(" ").collect();
   return nummer[1].parse::<i32>().unwrap();
}

fn power( line: &str) -> i32 {
    let toks : Vec<&str> = line.split(": ").collect();
    let tries: Vec<&str> = toks[1].split("; ").collect();
    let mut nums = HashMap::<String,i32>::new();
    for t in tries {
        let dices : Vec<&str> = t.split(", ").collect();
        for d in dices {
            let one : Vec<&str> = d.split(" ").collect();
            if one.len() != 2{
                println!("{}|{}|{}|{}", line, t, d, one.len());
                panic!("bummer");
            }
            let d = one[0].parse::<i32>().unwrap() ;
             let value = match nums.get(one[1]) {
                Some(v) =>  cmp::max( *v, d),
                None => d
            };
            nums.insert( one[1].to_string(), value);
        }
    }
    let mut power = 1;
    for n in nums.iter() {
        power *= n.1;
    }
   
   return power;
}


fn create_hashmap() -> HashMap<String,i32> {
    let mut ret = HashMap::new();
    // 12 red cubes, 13 green cubes, and 14 blue cubes?
    ret.insert( "red".to_string(), 12);
    ret.insert("green".to_string(), 13);
    ret.insert("blue".to_string(), 14);
    ret
}

#[cfg(test)]
#[test]
fn test() {
    let limit = create_hashmap();
    assert_eq!(possible( "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", &limit), 1);
    assert_eq!(possible( "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", &limit), 0);
}

#[test]
fn test2() {
    assert_eq!(power( "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"), 48);

    assert_eq!(power( "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"), 1560);
    
}


fn file( file_name : &str) -> i32{
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let mut sum = 0;
    let limit = create_hashmap();
    for line in contents.split('\n') {
        if line == "" {
            continue;
        }
        sum += possible( line, &limit);
    }
    sum
}

fn file2( file_name : &str) -> i32{
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let mut sum = 0;
    for line in contents.split('\n') {
        if line == "" {
            continue;
        }
        sum += power( line);
    }
    sum
}
fn main() {
    println!("{}", file("input.txt"));
    println!("{}", file2("input.txt"));
}
