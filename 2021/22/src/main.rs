use regex::Regex;
use std::collections::HashSet;

#[derive(Copy, Clone, Debug)]
struct Range(i32, i32);

#[derive(Copy, Clone, Debug)]
struct Block(Range,Range,Range);

#[derive(Copy, Clone, Debug,Hash, PartialEq)]
struct Coord(i32,i32,i32);

impl Eq for Coord {}

struct Instruction {
    on : bool,
    block : [Range;3]
}


/* on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10 */

fn parse(input: &str) -> Vec<Instruction> {
    let mut ret = Vec::new();
    let lines: Vec<_> = input.split('\n').collect();

    let r = Regex::new(r"^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)")
        .unwrap();

    for line in lines {
        let caps = r.captures(line).unwrap();
        let on = caps.get(1).unwrap().as_str() == "on";
        let mut i = Instruction {
            on: on,
            block: [Range(0, 0); 3],
        };
        for axes in 0..3 {
            let lower: i32 = caps.get(2 + axes * 2).unwrap().as_str().parse().unwrap();
            let upper: i32 = caps.get(3 + axes * 2).unwrap().as_str().parse().unwrap();
            i.block[axes].0 = lower;
            i.block[axes].1 = upper;
        }
        ret.push(i);
    }

    ret
}

fn splitPoints( instructions: & Vec<Instruction> )  -> [Vec<i32>;3] {
    let mut sets = vec![HashSet::new();3];
    for i in instructions {
        for d in 0..3 {
            sets[d].insert( i.block[d].0);
            sets[d].insert( i.block[d].1+1);
        }
    }
    let mut r = vec![Vec::new();3];
    for d in 0..3 {
        for i in &sets[d] {
            r[d].push( *i);
        }
    }
    for d in 0..3 {
        r[d].sort();
  /*      let len = r[d].len();
        let max = r[d][len-1];
        r[d].push(max+1);*/
    }
    
    r.try_into().unwrap()
}

fn findIndex( r: Range, indices : &Vec<i32>) -> (usize, usize) {
    let mut start = 0;
    while  r.0 != indices[start] {
        start+=1;
    }
    let mut end = start;
    while  r.1 + 1 != indices[end] {
        end+=1;
    }
    (start, end)
}

fn process( world: &mut HashSet<[usize;3]>, instructions: &Vec<Instruction>, split_points : &[Vec<i32>;3]) {
    let all = instructions.len();
    let mut count = 0;
    for is in instructions {
        println!("{}/{}", count, all);
        count+=1;
        let mut v = [0usize;3];
        let (xs, xe) = findIndex( is.block[0], &split_points[0]);
        for x in xs..xe {
            v[0]=x;
            let (ys, ye) = findIndex( is.block[1], &split_points[1]);
            for y in ys..ye {
                v[1]=y;
                let (zs, ze) = findIndex( is.block[2], &split_points[2]);
                for z in zs..ze {
                    v[2]=z;
                    if is.on {
                        world.insert( v);
                    } else {
                        world.remove( &v);
                    }
                }
            }
        }
    }
}

fn score( world: &HashSet<[usize;3]>, split_points : &[Vec<i32>;3]) -> u64 {
    let mut sum = 0;
    for w in world {
        let mut size = 1;
        for d in 0..3 {
            let index = w[d];
            size *= (split_points[d][index+1] - split_points[d][index]) as u64 ;
        }
        sum += size;
    }
    sum
}


fn main() {
    let input = "on x=10..12,y=10..12,z=10..12\n\
        on x=11..13,y=11..13,z=11..13\n\
        off x=9..11,y=9..11,z=9..11\n\
        on x=10..10,y=10..10,z=10..10";

    let input_real = include_str!("../input.txt");

    let is = parse(input_real);
    let sp = splitPoints(&is);
    let mut world = HashSet::new();
    process( &mut world, &is, &sp);
    let s = score( &world, &sp);
    println!( "{}", s);
}
