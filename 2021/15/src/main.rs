use fifteen::*;
use priority_queue::PriorityQueue;
use std::collections::HashMap;

fn main() {

    let input = include_str!("input.txt");
    let map = parse(&input);
    let mut queue = PriorityQueue::new();
    let point = Point { x: 0, y: 0 };

    let maxc = max_cost(&map) as i32;
    queue.push(point, -metric(&map, &point));
    let mut done_list = HashMap::new();
    done_list.insert(point, 0);
    loop {
        let res = advance(&map, &mut queue, &mut done_list, maxc);
        if let Some(val) = res {
            println!("{}", val);
            break;
        }
    }

    let map = five_times(&parse(&input));
    let mut queue = PriorityQueue::new();
    let point = Point { x: 0, y: 0 };

    let maxc = max_cost(&map) as i32;
    queue.push(point, -metric(&map, &point));
    let mut done_list = HashMap::new();
    done_list.insert(point, 0);
    loop {
        let res = advance(&map, &mut queue, &mut done_list, maxc);
        if let Some(val) = res {
            println!("{}", val);
            break;
        }
    }
}
