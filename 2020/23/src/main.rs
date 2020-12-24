
use std::iter::Iterator;
use std::vec::Vec;

fn round( input: & mut [usize], current : usize) -> usize {
    let next = input[current];
    let length = input.len();

    let mut destination = (current + length - 1) % length;
    let mut ptr = next;
    loop {
        let mut collision = false;
        ptr = current; 
        for _ in 0..3 {
            ptr = input[ptr];
            if ptr == destination {
                // try next destination
                collision = true;
                destination = (destination + length - 1) % length;
                break;
            }
            
        }
        if !collision {
            // collision free destination found
            break;
        }
    }
    // destination now valid
    // now we can rearrange references
    // next : index to first element of triple
    // ptr : index to last element of triple

    // now ptr to element after triple
    input[current] = input[ptr]; // 
    let tmp = input[destination]; // 4
    input[destination] = next; //
    input[ptr] = tmp; // 

    /*
    current = 2
    next = 7 
    destination = 1
    ptr = 0
    0 1 2 3 4 5 6 7 8    0 1 2 3 4 5 6 7 8   
    
    2 7 8 0 1 4 3 5 6    2 1 7 8 0 4 3 5 6

    1 4 7 5 3 6 2 8 0    4 7 1 5 3 6 2 8 0
                         X X X
    */

    input[current] 
}

fn result( v : &[usize]) -> String {
    let mut ptr = v[0];
    let mut result = Vec::new();
    for _ in 1..v.len() {
        result.push( ptr as u8 + '1' as u8);
        ptr = v[ptr];
    }
    String::from_utf8( result).unwrap()
}

fn transform( v :&[usize]) -> Vec<usize>{
    let mut result = Vec::with_capacity(v.len());
    result.resize(v.len(), 0);
    for i in 0..v.len() {
        let j = v[i]-1;
        result[j] =  v[(i+1)% v.len()]-1; 
    }
    result
}



fn main() {
    let mut v = vec![3, 8, 9, 1, 2, 5, 4, 6, 7];
    let mut current = v[0]-1;
    v = transform( &v);
    for _ in 0..10 {
        current = round(&mut v, current);
    }
    assert_eq!( result(&v), "92658374");
    v = vec![3, 8, 9, 1, 2, 5, 4, 6, 7];
    current = v[0]-1;
    v = transform( &v);
    for _ in 0..100 {
        current = round(&mut v, current);
    }
    assert_eq!( result(&v), "67384529");
    v = vec![7,1,6,8,9,2,5,4,3];
    current = v[0]-1;
    v = transform( &v);
    for _ in 0..100 {
        current = round(&mut v, current);
    }
    print!("{}\n", result(&v));
    //// part2
    let mut v = vec![3, 8, 9, 1, 2, 5, 4, 6, 7];
    for i in v.len()..1_000_000{
        v.push(i+1)
    }
    let mut current = v[0]-1;
    v = transform( &v);
    for _ in 0..10_000_000 {
        current = round(&mut v, current);
    }
    assert_eq!( v[0]+1, 934001);
    assert_eq!( v[v[0]]+1, 159792);

    v = vec![7,1,6,8,9,2,5,4,3];
    for i in v.len()..1_000_000{
        v.push(i+1)
    }
    current = v[0]-1;
    v = transform( &v);
    
    for _ in 0..10_000_000 {
        current = round(&mut v, current);
    }
    let one = v[0];
    let two = v[one];
    print!("{}\n", (one+1)*(two+1));
}
