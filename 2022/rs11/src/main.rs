

fn part1() {
    let mut monkies: Vec<fn(i32) -> (i32, i32)> = Vec::new();

    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x * 19) / 3;
        if (x % 17) == 0 {
            (2, x)
        } else {
            (7, x)
        }
    }); // 0
    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x + 2) / 3;
        if (x % 19) == 0 {
            (7, x)
        } else {
            (0, x)
        }
    }); // 1

    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x + 7) / 3;
        if (x % 7) == 0 {
            (4, x)
        } else {
            (3, x)
        }
    }); // 2
    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x + 1) / 3;
        if (x % 11) == 0 {
            (6, x)
        } else {
            (4, x)
        }
    }); // 3
    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x * 5) / 3;
        if (x % 13) == 0 {
            (6, x)
        } else {
            (5, x)
        }
    }); // 4
    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x + 5) / 3;
        if (x % 3) == 0 {
            (1, x)
        } else {
            (0, x)
        }
    }); // 5
    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x * x) / 3;
        if (x % 5) == 0 {
            (5, x)
        } else {
            (1, x)
        }
    }); // 6
    monkies.push(|x: i32| -> (i32, i32) {
        let x = (x + 3) / 3;
        if (x % 2) == 0 {
            (2, x)
        } else {
            (3, x)
        }
    }); // 7

    let mut business: Vec< Vec<i32>> = Vec::new();
    let mut counter = [0; 8];
    business.push( vec![83, 97, 95, 67]);
    business.push( vec![71, 70, 79, 88, 56, 70]);
    business.push( vec![98, 51, 51, 63, 80, 85, 84, 95]);
    business.push( vec![77, 90, 82, 80, 79]);
    business.push( vec![68]);
    business.push( vec![60, 94]);
    business.push( vec![81, 51, 85]);
    business.push( vec![98, 81, 63, 65, 84, 71, 84]);
    for _ in 0..20 {
        for m in 0..8 {
            let monkey = monkies[m];
            counter[m] += business[m].len();
            for b in business[m].clone() {
                let (new_m, worry) = monkey( b);
                business[ new_m as usize].push( worry);
            }
            business[m] = vec![];
        }
    }
    counter.sort();
    let ans = counter[6]* counter[7];

    println!("{}", ans);
}

fn part2() {
    let mut monkies: Vec<fn(i64) -> (i32, i64)> = Vec::new();
 
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x * 19) % (17*19*7*11*13*3*5*2) ;
        if (x % 17) == 0 {
            (2, x)
        } else {
            (7, x)
        }
    }); // 0
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x + 2) % (17*19*7*11*13*3*5*2) ;
        if (x % 19) == 0 {
            (7, x)
        } else {
            (0, x)
        }
    }); // 1

    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x + 7)% (17*19*7*11*13*3*5*2) ;
        if (x % 7) == 0 {
            (4, x)
        } else {
            (3, x)
        }
    }); // 2
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x + 1)% (17*19*7*11*13*3*5*2) ;
        if (x % 11) == 0 {
            (6, x)
        } else {
            (4, x)
        }
    }); // 3
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x * 5)% (17*19*7*11*13*3*5*2) ;
        if (x % 13) == 0 {
            (6, x)
        } else {
            (5, x)
        }
    }); // 4
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x + 5)% (17*19*7*11*13*3*5*2) ;
        if (x % 3) == 0 {
            (1, x)
        } else {
            (0, x)
        }
    }); // 5
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x * x)% (17*19*7*11*13*3*5*2) ;
        if (x % 5) == 0 {
            (5, x)
        } else {
            (1, x)
        }
    }); // 6
    monkies.push(|x: i64| -> (i32, i64) {
        let x = (x + 3)% (17*19*7*11*13*3*5*2) ;
        if (x % 2) == 0 {
            (2, x)
        } else {
            (3, x)
        }
    }); // 7

    let mut business: Vec< Vec<i64>> = Vec::new();
    let mut counter = [0; 8];
    business.push( vec![83, 97, 95, 67]);
    business.push( vec![71, 70, 79, 88, 56, 70]);
    business.push( vec![98, 51, 51, 63, 80, 85, 84, 95]);
    business.push( vec![77, 90, 82, 80, 79]);
    business.push( vec![68]);
    business.push( vec![60, 94]);
    business.push( vec![81, 51, 85]);
    business.push( vec![98, 81, 63, 65, 84, 71, 84]);
    for _ in 0..10000 {
        for m in 0..8 {
            let monkey = monkies[m];
            counter[m] += business[m].len();
            for b in business[m].clone() {
                let (new_m, worry) = monkey( b);
                business[ new_m as usize].push( worry);
            }
            business[m] = vec![];
        }
    }
    counter.sort();
    let ans = counter[6]* counter[7];

    println!("{}", ans);
}

fn main() {
    part1();
    part2();
}
