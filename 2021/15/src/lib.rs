use priority_queue::PriorityQueue;
//
use std::collections::HashMap;

#[derive(Debug, PartialEq, Clone, Copy, Hash)]
pub struct Point {
    pub x: usize,
    pub y: usize,
}

impl Eq for Point {}

pub fn parse(input: &str) -> Vec<Vec<u8>> {
    let lines: Vec<_> = input.split('\n').collect();
    let mut matrix = Vec::new();
    for line in lines {
        let row: Vec<_> = line
            .chars()
            .map(|x| -> u8 { x as u8 - '0' as u8 })
            .collect();
        matrix.push(row);
    }

    matrix
}

pub fn max_cost(input: &Vec<Vec<u8>>) -> u32 {
    let size_x = input[0].len();
    let size_y = input.len();
    let mut sum = 0;
    for i in 0..size_y {
        sum += input[i][0] as u32;
    }
    for i in 0..size_x {
        sum += input[0][i] as u32;
    }
    sum
}

pub fn five_times(input: &Vec<Vec<u8>>) -> Vec<Vec<u8>> {
    let size_x = input[0].len();
    let size_y = input.len();
    let mut ret = Vec::new();
    for _ in 0..size_y * 5 {
        ret.push(vec![0; 5 * size_x]);
    }
    for y in 0..size_y * 5 {
        for x in 0..size_y * 5 {
            ret[y][x] =
                (((input[y % size_y][x % size_x] as i32 + (y / size_y + x / size_x) as i32) - 1)
                    % 9
                    + 1) as u8;
        }
    }
    ret
}

pub fn dist(map: &Vec<Vec<u8>>) -> Vec<u32> {
    let sizey = map.len();
    let sizex = map[0].len();
    let mut metric = vec![9; (sizex - 1) + (sizey - 1) + 1];
    for i in 0..sizey {
        for j in 0..sizex {
            metric[(sizex - j - 1) + (sizey - i - 1)] =
                std::cmp::min(map[i][j], metric[(sizex - j - 1) + (sizey - i - 1)]);
        }
    }
    let mut ret = vec![0; (sizex - 1) + (sizey - 1) + 1];
    ret[0] = 0;
    for i in 1..=(sizex - 1) + (sizey - 1) {
        ret[i] = ret[i - 1] + metric[i - 1] as u32;
    }

    ret
}

pub fn metric(map: &Vec<Vec<u8>>, pos: &Point) -> i32 {
    let sizey = map.len();
    let sizex = map[0].len();

    ((sizex - pos.x - 1) + (sizey - pos.y - 1)) as i32
}

fn succ(map: &Vec<Vec<u8>>, last: &Point) -> Vec<Point> {
    let sizey = map.len();
    let sizex = map[0].len();
    let mut res = Vec::new();
    if last.x > 0 {
        let p = Point {
            x: last.x - 1,
            y: last.y,
        };
        res.push(p);
    }
    if last.x < (sizex - 1) {
        let p = Point {
            x: last.x + 1,
            y: last.y,
        };

        res.push(p);
    }
    if last.y > 0 {
        let p = Point {
            x: last.x,
            y: last.y - 1,
        };

        res.push(p);
    }
    if last.y < (sizey - 1) {
        let p = Point {
            x: last.x,
            y: last.y + 1,
        };

        res.push(p);
    }
    res
}

pub fn advance(
    map: &Vec<Vec<u8>>,
    queue: &mut PriorityQueue<Point, i32>,
    done_list: &mut HashMap<Point, usize>,
    // came_from: &mut HashMap<Point, Point>,
    max_cost: i32,
) -> Option<usize> {
    let sizex = map[0].len();
    let sizey = map.len();
    let curr = queue.pop().unwrap();
    if curr.0.x == (sizex - 1) && curr.0.y == (sizey - 1) {
        return Some(done_list[&curr.0]);
    }
    let list = succ(map, &curr.0);
    let curr_cost = *done_list.get(&curr.0).unwrap();

    for el in list {
        let cost = curr_cost + map[el.y][el.x] as usize;
        let metric = metric(map, &el);
        if let Some(known_costs) = done_list.get(&el) {
            if (*known_costs) < cost {
                continue;
            }
        }
        done_list.insert(el, cost);
        let prio = -(cost as i32 + metric);
        if prio < -max_cost {
            continue;
        }
        if let Some(_) = queue.get(&el) {
            queue.change_priority(&el, prio);
        } else {
            queue.push(el, prio);
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example() {
        let input = include_str!("example.txt");
        let map = parse(input);
        let mut queue = PriorityQueue::new();
        let point = Point { x: 0, y: 0 };

        let maxc = max_cost(&map) as i32;
        queue.push(point, -metric(&map, &point));
        let mut done_list = HashMap::new();
        done_list.insert(point, 0);
        loop {
            let res = advance(&map, &mut queue, &mut done_list, maxc);
            if let Some(val) = res {
                assert_eq!(val, 40);
                return;
            }
        }
    }

    #[test]
    fn part2() {
        let input = include_str!("example.txt");
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
                assert_eq!(val, 315);
                return;
            }
        }
    }
}
