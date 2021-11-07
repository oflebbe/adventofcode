#[repr(i32)]
#[derive(Clone,Copy,Debug)]
enum Direction {
    East = 0,
    South = 1,
    West = 2,
    North = 3,
}

fn solve( input: &str) -> i32 {
    let mut posx = 0;
    let mut posy = 0;
    let mut d = Direction::East;

    for line in input.lines() {
        
        let (ch, r) = line.split_at(1);
        let num = r.parse::<i32>().unwrap();
        match ch {
            "E" => posx += num,
            "W" => posx -= num,
            "S" => posy += num,
            "N" => posy -= num,
            "R" => {
                let n2 = num / 90;
                let k = (d as i32 + n2) % 4;
                d = match k {
                    0 => Direction::East,
                    1 => Direction::South,
                    2 => Direction::West,
                    3 => Direction::North,
                    _ => panic!("Direction")
                }
            }
            "L" => {
                let n2 = num / 90;
                let k = (d as i32 + 4 - n2) % 4;
                d = match k {
                    0 => Direction::East,
                    1 => Direction::South,
                    2 => Direction::West,
                    3 => Direction::North,
                    _ => panic!("Direction")
                };
               // println!("{:?}", d);
               // d = unsafe { ::std::mem::transmute(k) }; // Hack
            }
            "F" => {
                match d {
                    Direction::East => posx += num,
                    Direction::West => posx -= num,
                    Direction::South => posy += num,
                    Direction::North => posy -= num,
                };
            }
            _ => panic!("Bug")
        }
       // println!("{}, {}, {}, {:?}", line, posx, posy, d);
    }
    posx.abs() + posy.abs()
}

fn solve2( input: &str) -> i32 {
    let mut posx = 0;
    let mut posy = 0;
    
    let mut wayx = 10;
    let mut wayy = -1;

    for line in input.lines() {
        
        let (ch, r) = line.split_at(1);
        let num = r.parse::<i32>().unwrap();
        match ch {
            "E" => wayx += num,
            "W" => wayx -= num,
            "S" => wayy += num,
            "N" => wayy -= num,
            "R" => {
                let n2 = num / 90;
                for _ in 0..n2 {
                    let tmp = -wayy;
                    wayy = wayx;
                    wayx = tmp;
                    // ??? (wayx, wayy) = (-wayy, wayx);
                }
            }
            "L" => {
                let n2 = num / 90;
                for _ in 0..n2 {
                    let tmp = wayy;
                    wayy = -wayx;
                    wayx = tmp;
                    // ??? (wayx, wayy) = (-wayy, wayx);
                }
               // println!("{:?}", d);
               // d = unsafe { ::std::mem::transmute(k) }; // Hack
            }
            "F" => {
                posx += num * wayx;
                posy += num * wayy;
            }
            _ => panic!("Bug")
        }
       // println!("{}, {}, {}, {:?}", line, posx, posy, d);
    }
    posx.abs() + posy.abs()
}
fn main() {
    let i = "F10
N3
F7
R90
F11";
    assert_eq!(25, solve( i));

    let input = include_str!("input.txt");
    println!("{}", solve(input));
    println!("{}", solve2(input));

}