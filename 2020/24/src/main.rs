use std::collections::HashMap;

fn handle( line: &str) -> (i32, i32) {
    let mut ch = line.chars();
    let mut posx = 0i32;
    let mut posy = 0i32;
    loop {
       let (x, y) = match ch.next() {
            Some('e') => (1,0),
            Some('w') => (-1,0),
            Some('n') => match ch.next() {
                Some('e') => (0,1),
                Some('w') => (-1,1),
                _ => panic!("problem")
            },
            Some('s') => match ch.next() {
                Some('e') => (1,-1),
                Some('w') => (0,-1),
                _ => panic!("problem")
            },
            None => break,
            _ => panic!("problem")
        };
        posx += x;
        posy += y;
    }
    (posx, posy)
}

fn problem( input : &str) -> u32 {
    let lines : Vec<&str> = input.lines().collect();
    let mut counter = HashMap::new();
    for line in lines {
        
            let (x, y) = handle( line );
            let c = counter.entry((x,y)).or_insert(0);
            *c+=1;
    }
    let mut black = 0;
    for (k, v) in counter {
        if v % 2 == 1 {
            black +=1;
        }
    }
    black
}

fn main() {
    let input = include_str!("test.txt");
    print!("{}\n", problem( input));

    let input = include_str!("input.txt");
    print!("{}\n", problem( input));
}
