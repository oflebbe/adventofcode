use twentyone::Player;
use twentyone::Universes;
use twentyone::advance1;
use twentyone::advance2;
use std::collections::HashMap;
use std::time::Instant;

fn main() {
    let timer = Instant::now();
    let mut mul = vec![0;10];
    for i in 1..=3 {
        for j in 1..=3 {
            for k in 1..=3 {
                mul[i+j+k as usize]+=1;
            }
        }
    } 
    let mut u = HashMap::new();
    u.insert(
        (Player { pos: 5, score: 0 }, Player { pos: 9, score: 0 }),
        1,
    );
    let mut universes = Universes { u: u, wins: (0, 0) };

    while universes.u.len() > 0 {
        universes = advance1(&universes, &mul);
        universes = advance2(&universes, &mul);
    }


    println!("player1: {}", universes.wins.0 );
    println!("player2: {}", universes.wins.1 );

    println!("-> Solved day 21 in {:?}\n", timer.elapsed());
}
