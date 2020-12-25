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

fn problem( input : &str) ->  HashMap<(i32,i32),bool> {
    let lines : Vec<&str> = input.lines().collect();
    let mut tiles = HashMap::new();
    for line in lines {
            let (x, y) = handle( line );
            let c = tiles.entry((x,y)).or_insert(false);
            // flip
            *c = !*c;
    }
    tiles
}



// true == black
fn day( tiles : &HashMap<(i32,i32),bool>) -> HashMap<(i32,i32),bool>    {

    let maxx = tiles.iter().map( |(&key,_)| key.0).max().unwrap()+1;
    let maxy = tiles.iter().map( |(&key,_)| key.1).max().unwrap()+1;
    let minx = tiles.iter().map( |(&key,_)| key.0).min().unwrap()-1;
    let miny = tiles.iter().map( |(&key,_)| key.1).min().unwrap()-1;
   
    let mut newDay : HashMap<(i32,i32),bool> = HashMap::new();

    for x in minx..=maxx {
        for y in miny..=maxy {
            let mut blacks = 0;
            let coord = (x,y);
            for d in 0..6 {
                let (dx, dy) = match d {
                    0 => ( 1, 0),
                    1 => ( -1, 0),
                    2 => ( 0, 1),
                    3 => ( 0, -1),
                    4 => (-1, 1),
                    5 => ( 1, -1),
                    _ => panic!("Problem") 
                };
                let v = (x + dx, y+dy);
                blacks += match tiles.get(&v) {
                    Some(&true) => 1,
                    _ => 0,
                }
            }
            let c = newDay.entry(coord).or_insert(false);
            *c = match tiles.get(&coord) {
                Some(&true) => match blacks {
                        0|3|4|5|6 => false,
                        _ => true,
                    },
                Some(&false)| None => match blacks {
                    2 => true,
                    _ => false,
                },
            }
        }
    }
    newDay
}

fn counter( tiles : &HashMap<(i32,i32),bool>) -> u32 {
    let mut count = 0;
    for (_, v) in tiles {
        if v == &true {
            count+=1;
        }
    }
    count
}


fn main() {
    let input = include_str!("test.txt");
    let mut test = problem( input);
    let black = counter(&test);
    
    assert_eq!(black ,10);

    for i in 0..100 {
        test = day( &test);
    }
    let black = counter(&test);
    assert_eq!(black ,2208);

    let input = include_str!("input.txt");
    let mut tiles = problem( input);
  
    print!("{}\n",black);
 
    for i in 0..100 {
        tiles =  day( &tiles);
       
    }
    println!("{}",counter(&tiles));
}
