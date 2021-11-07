use std::collections::HashMap;

// true == black
fn day( tiles : &HashMap<(i32,i32),bool>, day: i32) -> HashMap<(i32,i32),bool>    {
    let mut newDay :  HashMap<(i32,i32),bool> = HashMap::new();

    for x in -day..=day {
        for y in -day..=day {
            let mut blacks = 0;
            let coord = (x,y);
            for d in 0..6 {
                let (dx, dy ) = match d {
                    0 => ( 1, 0),
                    1 => ( -1, 0),
                    2 => ( 0, 1),
                    3 => ( 0, -1),
                    4 => (-1, 1),
                    5 => ( 1, -1),
                    _ => panic!("Problem") 
                };
                let v = (x + dx, y+dy);
                match tiles.get(&v) {
                    Some(true) => blacks +=1,
                    _ => {},
                }
            }
            let c = newDay.entry(coord).or_insert(false);
            *c = match tiles.get(&coord) {
                Some(true) => match blacks {
                        0|2|3|4|5|6 => true,
                        _ => false
                    },
                Some(false)| None => match blacks {
                    2 => true,
                    _ => false
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
    let mut val : HashMap<(i32,i32),bool> = HashMap::new();
    for i in 0..10 {
        val =  day( &val, i+1);
        println!("{}",counter(&val));
    }
}