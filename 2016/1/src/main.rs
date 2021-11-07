use std::collections::HashMap;

fn rotate( d: [i8; 2], rotate: [i8; 4]) -> [i8; 2] {

    let ret : [i8; 2] = [
    d[0] * rotate[0] + d[1] * rotate[1],
    d[0] * rotate[2] + d[1] * rotate[3] 
    ];
    
    ret
}

fn follow( line : &str) -> u32 {
    let mut d : [i8; 2] = [0, 1];

    let rotate_r : [i8; 4] = [0, 1, -1, 0];
    let rotate_l : [i8; 4] = [0, -1, 1, 0];
    let mut k : [i32; 2] = [0, 0];
   // println!("{:?}", d);
    for s in line.split(", ") {
        match s.chars().nth(0) {
            Some('L') => d = rotate( d, rotate_l),
            Some('R') => d = rotate( d, rotate_r),
            _ => println!("Error")
        }
        let length = s[1..].parse::<i32>();
        match length  {
            Ok(i) => { 
                k[0] += d[0] as i32 * i;
                k[1] += d[1] as i32 * i;
            }
            Err(_e) => println!("Error")

        }
        
     //  println!("{}, {:?}, {:?}",s, d, k);
    }
    ( k[0].abs() + k[1].abs()) as u32

}

fn follow_twice( line : &str) -> u32 {
    let mut grid = HashMap::new();
    let mut d : [i8; 2] = [0, 1];

    let rotate_r : [i8; 4] = [0, 1, -1, 0];
    let rotate_l : [i8; 4] = [0, -1, 1, 0];
    let mut k : [i32; 2] = [0, 0];
   // println!("{:?}", d);
    let st = format!("{:?}", k);
    grid.insert(st, 1);
    for s in line.split(", ") {
        match s.chars().nth(0) {
            Some('L') => d = rotate( d, rotate_l),
            Some('R') => d = rotate( d, rotate_r),
            _ => println!("Error")
        }
        let length = s[1..].parse::<i32>();
        let l = match length  {
            Ok(i) => i,
            Err(_e) => { println!("Error"); 0 }
        };
        for  _i in 0..l {
            k[0] += d[0] as i32;
            k[1] += d[1] as i32;
            let st = format!("{:?}", k);
            let counter = grid.entry(st.clone()).or_insert(0);
            *counter += 1;
            if *counter == 2 {
               return ( k[0].abs() + k[1].abs()) as u32
            }
        }
        
        
    }
     //  println!("{}, {:?}, {:?}",s, d, k);
    
    0

}

fn main() {

    let x = &[0, 1, 2];

    if let Some((last, elements)) = x.split_last() {
        assert_eq!(last, &2);
        assert_eq!(elements, &[0, 1]);
    }

    println!("{} 5", follow("R2, L3"));
    println!("{} 2", follow("R2, R2, R2"));
    println!("{} 12", follow("R5, L5, R5, R3"));
    
   
    let my_str = include_str!("input.txt").trim_end();
    println!("{}", follow( my_str));
    println!("{} 4", follow_twice( "R8, R4, R4, R8"));
    println!("{}", follow_twice( my_str));

}
