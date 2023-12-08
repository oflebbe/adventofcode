use std::collections::HashMap;
use std::fs::read_to_string;


pub fn parse( file_name : &str) -> (String, HashMap< String, ( String, String)>) {
    let mut res = HashMap::new();

    let contents = read_to_string(file_name)
        .expect("Should have been able to read the file");

    let parts : Vec<&str> = contents.split("\n\n").collect();

    // JHK = (LRD, DDR)
    // 0123456789012345
    for line in parts[1].split("\n") {
        let point = line[0..3].to_string();
        let l = line[7..10].to_string();
        let r = line[12..15].to_string();
        res.insert( point, (l, r));
    
    }

    (parts[0].to_string(), res)
}

pub fn star1( instruction : &str, map : &HashMap< String, ( String, String)>) -> i32{
    determine_cycle( instruction, map, "AAA")
}

fn determine_cycle(instruction : &str, map : &HashMap< String, ( String, String)>, start: &str) -> i32 { 
    let inst_list = instruction.as_bytes();
    let mut inst_ptr = 0;
    let mut current = start;
    let mut count = 0;
    while !current.ends_with("Z") {
        let next = map.get( current).unwrap();
        if inst_list[inst_ptr] == b'L' {
            current =&next.0;
        } else {
            current = &next.1;
        }
        inst_ptr = (inst_ptr + 1) % inst_list.len();
        count+=1;
 
    }
    count
}


fn lcm(first: usize, second: usize) -> usize {
    first * second / gcd(first, second)
}

fn gcd(first: usize, second: usize) -> usize {
    let mut max = first;
    let mut min = second;
    if min > max {
        let val = max;
        max = min;
        min = val;
    }

    loop {
        let res = max % min;
        if res == 0 {
            return min;
        }

        max = min;
        min = res;
    }
}					

pub fn star2(  instruction : &str, map : &HashMap< String, ( String, String)>) -> usize {

    // extract all start nodes
    let current: Vec<&str> = map.keys().filter( |x|  x.ends_with('A')).map(|x| x.as_str()).collect();

    let mut prod = 1;
    for i in current {
        let z =  determine_cycle( instruction, map, i) as usize;
        println!("{}", z);
        prod = lcm(prod, z);
        
    }
    prod
}

#[test]
fn test( ) {
    let (i, m) = parse( "test2.txt");
    assert_eq!( star1( &i, &m),6);
}

#[test]
fn test2() {
    let (i, m) = parse( "test3.txt");
     assert_eq!(star2( &i, &m), 6);
}