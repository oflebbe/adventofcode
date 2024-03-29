use std::collections::HashMap;
use std::fs;
use std::cmp::Ordering;

#[derive(Clone)]
pub struct Cards {
    cards: String,
    bid: i32,
}

enum Type {
    FiveOfAKind= 7,
    FourOfAKind = 6,
    FullHouse = 5,
    ThreeOfAKind = 4,
    TwoPair = 3,
    Pair = 2,
    HighCard = 1,
}

fn determine_type(cards: &str, joker_mode: bool) -> Type {
    let mut counter = HashMap::new();
    let mut jokers = 0;
    for c in cards.as_bytes() {
        if joker_mode && *c == b'J' {
            jokers += 1;
            continue;
        }
        if counter.contains_key(c) {
            counter.insert(c, counter.get(c).unwrap() + 1);
        } else {
            counter.insert(c, 1);
        }
    }
    let mut s : Vec<_> = counter.values().map(|x| *x).collect();
    s.sort();
    s.reverse();
    if jokers == 5 {
        return  Type::FiveOfAKind
    }
    match s[0] + jokers {
        5 => Type::FiveOfAKind,
        4 => Type::FourOfAKind,
        3 => if s[1] == 2 {
             Type::FullHouse
        } else {
             Type::ThreeOfAKind
        },
        2 => if s[1] == 2 {
                Type::TwoPair
        } else {
            Type::Pair
        },
        _ => Type::HighCard
    }
}

fn value( val : u8, joker_mode: bool) -> i32 {
    match val {
        b'A' => 14,
        b'K' => 13,
        b'Q' => 12, 
        b'J' => if joker_mode { 1} else { 11},
        b'T' => 10,
        b'9' => 9,
        b'8' => 8,
        b'7' => 7, 
        b'6' => 6,
        b'5' => 5,
        b'4' => 4,
        b'3' => 3, 
        b'2' => 2,
        _ => panic!("should not happen")
    }
}


fn higher( a : &Cards, b: &Cards, joker_mode: bool) -> Ordering {
    let a_t = determine_type( &a.cards, joker_mode) as i32;
    let b_t = determine_type( &b.cards, joker_mode) as i32;
    let s = a_t.cmp(&b_t);
    if s != Ordering::Equal {
        return  s
    } 
    let a_s = a.cards.as_bytes();
    let b_s = b.cards.as_bytes();
    for i in 0..a.cards.len() {
        let s = value(a_s[i], joker_mode).cmp(&value(b_s[i], joker_mode));
        if s == Ordering::Equal {
            continue;
        }
        return s
    }
    Ordering::Equal
}

pub fn parse(file_name: &str) -> Vec<Cards> {
    let contents: Vec<String> = fs::read_to_string(file_name)
        .expect("Should have been able to read the file")
        .split("\n").map(|x| x.to_string())
        .collect();
    let mut list = Vec::new();
    for line in contents {
        if line.len() < 1 {
            continue;
        }
        let toks: Vec<&str> = line.split(' ').collect();
        list.push(Cards {
            cards: toks[0].to_string(),
            bid: toks[1].parse::<i32>().unwrap(),
        });
    }
    list
}



pub fn eval( list : &Vec<Cards>, joker_mode: bool) -> i32 {
    let mut v = list.clone();
    v.sort_by( |a, b| higher( a, b, joker_mode));
    
    let mut sum = 0;
    for i in 0..list.len() {
        sum += ( i + 1)as i32 * v[i].bid;
    }
    sum
}


#[cfg(test)]


#[test]
fn test1() {
    let a = Cards{ cards: "KK677".to_string(), bid: 1};
    let b = Cards{ cards: "KTJJT".to_string(), bid: 2};
    assert_eq!( higher( &a, &b, false), Ordering::Greater);

    let c = Cards{ cards: "KK677".to_string(), bid: 3};
    assert_eq!( higher( &a, &c, false), Ordering::Equal);

    let d = Cards{ cards: "32T3K".to_string(), bid: 4};
    assert_eq!( higher( &a, &d, false), Ordering::Greater);

    let e = Cards{ cards: "T55J5".to_string(), bid: 5};
    assert_eq!( higher( &a, &e, false), Ordering::Less);

    let f = Cards{ cards: "QQQJA".to_string(), bid: 6};
    assert_eq!( higher( &f, &e, false), Ordering::Greater);
}

#[test]
fn test2() {
    let a = Cards{ cards: "KK677".to_string(), bid: 1};
    let b = Cards{ cards: "KTJJT".to_string(), bid: 2};
    assert_eq!( higher( &a, &b, true), Ordering::Less);

    let c = Cards{ cards: "KK677".to_string(), bid: 3};
    assert_eq!( higher( &a, &c, true), Ordering::Equal);

    let d = Cards{ cards: "32T3K".to_string(), bid: 4};
    assert_eq!( higher( &a, &d, true), Ordering::Greater);

    let e = Cards{ cards: "T55J5".to_string(), bid: 5};
    assert_eq!( higher( &b, &e, true), Ordering::Greater);

    let f = Cards{ cards: "QQQJA".to_string(), bid: 6};
    assert_eq!( higher( &f, &e, true), Ordering::Greater);
}



#[test]
fn test_parse() {
    let limit = eval(&parse( "test.txt"), false);

    assert_eq!( limit, 6440);
}


#[test]
fn test_parse2() {

    let limit =eval( &parse( "test.txt"), true);
    

    assert_eq!(limit, 5905);
}

