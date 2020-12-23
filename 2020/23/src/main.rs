
use std::iter::Iterator;
use std::vec::Vec;

// convention: current is always first
fn round(input: Vec<u8>) -> Vec<u8> {
    let a = input[0];
    let three: Vec<u8> = input.iter().skip(1).take(3).map(|&x| x).collect();
    let mut insertVal = ((a as usize -1  + input.len() - 1) % input.len()) as u8 + 1;
    while three.contains(&insertVal) {
        insertVal = ((insertVal as usize -1 + input.len() - 1) % input.len()) as u8 + 1;
    }
    let remaining: Vec<u8> = input.iter().skip(4).map(|&x| x).collect();
    assert_eq!(remaining.contains(&insertVal), true);

    let c = remaining.clone();
    let mut iter = c.split(|&num| num == insertVal);

    let mut new = vec![
        vec![a],
        iter.next().unwrap().to_vec(),
        vec![insertVal],
        three,
        iter.next().unwrap().to_vec(),
    ]
    .concat();
    new.rotate_left(1);
    new
}

fn result( mut v : Vec<u8>) -> String {
    while v[0] != 1 {
        v.rotate_left(1)
    }
    let w : Vec<u8> = v.iter().skip(1).map(|&x| x + '0' as u8 ).collect();
    String::from_utf8( w).unwrap()
}

fn main() {

    let mut v = vec![3, 8, 9, 1, 2, 5, 4, 6, 7];

    for _ in 0..10 {
        v = round(v);
    }
    assert_eq!( result(v), "92658374");
    v = vec![3, 8, 9, 1, 2, 5, 4, 6, 7];
    for _ in 0..100 {
        v = round(v);
    }
    assert_eq!( result(v), "67384529");
    v = vec![7,1,6,8,9,2,5,4,3];
    for _ in 0..100 {
        v = round(v);
    }
    print!("{}\n", result(v));
}
