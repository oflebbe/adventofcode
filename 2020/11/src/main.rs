#[derive(PartialEq,Clone,Copy)]
enum Position {
    Nothing,
    Empty,
    Occupied,
}

use std::fmt;

struct Field {
    // store linear, leave 1 cell free on either side
    grid: Vec<Position>,
    width: usize,
    height: usize
}

impl fmt::Display for Field {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for y in 1..self.height-1 {
            for x in 1..self.width-1 {
                let ch = match self.get(x,y) {
                    Position::Nothing => '.',
                    Position::Empty => 'L',
                    Position::Occupied => '#',
                };
                write!(f, "{}", ch).unwrap();
            }
            write!(f,"\n").unwrap();
        }
        write!(f,"\n")
    }
} 

impl Field {
    fn get_mut(&mut self, x : usize, y : usize) -> & mut Position {
        &mut self.grid[ y * self.width + x] 
    }

    fn get(& self, x : usize, y : usize) -> Position {
        self.grid[ y * self.width + x] 
    }

    fn occupied(& self) -> usize {
        self.grid.iter().filter( |x| **x == Position::Occupied).count()
    } 
}

fn readin(input: &str) -> Field {
    let lines : Vec<&str> = input.lines().collect();
    let width = lines[0].len()+2;
    let height = lines.len()+2 ;
    let mut field = Field { grid: vec!(Position::Nothing; width*height), width : width, height: height };

    for y in 1..height-1 {
        let line = lines[y-1].as_bytes();
        for x in 1..width-1  {
            let pos = match line[x-1] {
                b'L' => Position::Empty,
                b'#' => Position::Occupied,
                b'.' =>  Position::Nothing,
                _ => panic!("wrong input"),
            };
            *field.get_mut( x, y) = pos;
        }
    }

    field
}

fn step(field: &Field) -> (Field, bool) {
    let mut res = Field { grid: vec!(Position::Nothing; field.width*field.height), width : field.width, height: field.height };
    let mut changes = false;
    
    for y in 1..field.height-1 {
        for x in 1..field.width-1 {
            if field.get(x,y) == Position::Nothing {
                continue;
            }
            let mut next_to = 0;
            for ys in -1..=1 {
                for xs in -1..=1 {
                    let xx = (x as i32 + xs) as usize;
                    let yy = (y as i32 + ys) as usize;
                    {
                        if field.get(xx, yy) == Position::Occupied {
                            next_to += 1;
                        }
                    }
                }
            }
            if field.get(x, y)  == Position::Occupied {
                // next_to one too large!
                if next_to >= 5 {
                    *res.get_mut(x,y) = Position::Empty; 
                    changes = true;
                } else {
                    *res.get_mut(x,y) = Position::Occupied;
                }
            } else {
                if next_to == 0 {
                    *res.get_mut(x,y) = Position::Occupied;
                    changes = true;
                } else {
                    *res.get_mut(x,y) = Position::Empty; 
                }
            }
        }
    }
    (res, changes)
}

fn main() {

    let t = "L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL";
    let mut field = readin(t);
    let mut count = 0;
    loop {
        count += 1;
        let (new, changed) = step(&field);
        if !changed {
            break;
        }
        field = new;
    }
    assert_eq!(6, count);
    assert_eq!(37, field.occupied());

    let i = include_str!("input.txt");
    let mut count = 0;
    let mut field = readin(i);
    loop {
        count += 1;
        let (new, changed) = step(&field);
        // print!("{}", new);
        if !changed {
            break;
        }
        field = new;
    }
    println!("{} {}", count, field.occupied());
}
