use regex::Regex;
use std::collections::HashMap;
use std::cmp::Ordering;


fn rot( c : char, i : i32) -> char {
    let z = 'z' as i32 - 'a' as i32 + 1;

    let c = ((c as i32  - 'a' as i32 ) + (i % z)  ) % z + 'a' as i32 ;
    (c as u8) as char
}

fn decrypt( st: &str, room: i32) -> String {
    let mut v = Vec::new();
    for c in st.chars() {
        if c == '-' {
            v.push( ' ');
        } else {
            v.push( rot(c, room));
        }
    }
    let s: String = v.into_iter().collect();
    s
}

fn calc( input: &str)-> i32 {
    let re = Regex::new(r"^([\D-]*)-(\d*)\[(\D{5})]").unwrap();
    let mut room_sum = 0;
    for l in input.lines() {
        let c = re.captures(l).unwrap();
        let st = c.get(1).unwrap().as_str();
        let room = c.get(2).unwrap().as_str().parse::<i32>().unwrap();
        let chk = c.get(3).unwrap().as_str();
        let mut freq = HashMap::new();
        
        for c in st.chars() {
            if c == '-' {
                continue;
            }
            let count = freq.entry(c).or_insert(0);
            *count += 1;
        }
        let mut vec = Vec::new();
        for (key, value) in freq {
            vec.push( (key, value));
        }
        vec.sort_by( |a, b| { 
            if a.1 > b.1 { 
                return Ordering::Less ;
            } else { 
                if a.1 < b.1 {
                    return Ordering::Greater;
                } else {
                    return a.0.cmp( &b.0 );
                }
            }
        });
        let s: String = vec.into_iter().take(5).map( |x| x.0).collect();
        if s.eq(chk) {
            room_sum += room;
        }
        if decrypt(st, room).eq("northpole object storage") {
            println!("NN: {}", room)
        } 
    }
    room_sum
}

fn main() {
   let input = "aaaaa-bbb-z-y-x-123[abxyz]\n\
                a-b-c-d-e-f-g-h-987[abcde]\n\
                not-a-real-room-404[oarel]\n\
                totally-real-room-200[decoy]\n";
 

    
    println!("{}", calc(input));

    let my_str = include_str!("input.txt");
    println!("{}", calc(my_str));


    assert_eq!( rot( 'a', 1), 'b');
    assert_eq!( rot( 'z', 1), 'a');
    assert_eq!( rot( 'a', 2), 'c');
 
    assert_eq!(decrypt("qzmt-zixmtkozy-ivhz", 343), "very encrypted name");

    
    
}
