use std::collections::HashMap;

#[derive(Copy, Clone, PartialEq, Hash)]
pub struct Player {
    pub pos: i32,
    pub score: i32,
}

impl Eq for Player {}

impl Player {
    pub fn advance(&mut self, howmuch: i32) {
        self.pos += howmuch;
        self.pos = ((self.pos - 1) % 10) + 1;
        self.score += self.pos;
    }
}

pub struct Universes {
    pub u: HashMap<(Player, Player), i64>,
    pub wins: (i64, i64),
}

pub fn advance1(u: &Universes, mul: &Vec<i64>) -> Universes {
    let mut newUniverses = Universes {
        u: HashMap::new(),
        wins: u.wins,
    };
    for ((p1, p2), num) in &u.u {
        for i in 3usize..=9 {
            let mut pp1 = *p1;
            pp1.advance(i as i32);
            if pp1.score >= 21 {
                newUniverses.wins.0 += num * mul[i];
            } else {
                if let Some(c) = newUniverses.u.get(&(pp1, *p2)) {
                    let new_c = *c;
                    newUniverses.u.insert((pp1, *p2), new_c + num * mul[i]);
                } else {
                    newUniverses.u.insert((pp1, *p2), *num * mul[i]);
                }
            }
        }
    }
    newUniverses
}

pub fn advance2(u: &Universes, mul: &Vec<i64>) -> Universes {
    let mut newUniverses = Universes {
        u: HashMap::new(),
        wins: u.wins,
    };
    for ((p1, p2), num) in &u.u {
        for i in 3usize..=9 {
            let mut pp2 = *p2;
            pp2.advance(i as i32);
            if pp2.score >= 21 {
                newUniverses.wins.1 += num * mul[i];
            } else {
                if let Some(c) = newUniverses.u.get(&(*p1, pp2)) {
                    let new_c = *c;
                    newUniverses.u.insert((*p1, pp2), new_c + num * mul[i]);
                } else {
                    newUniverses.u.insert((*p1, pp2), *num * mul[i]);
                }
            }
        }
    }
    newUniverses
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example() {
        // i am lazy
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
            (Player { pos: 4, score: 0 }, Player { pos: 8, score: 0 }),
            1,
        );
        let mut universes = Universes { u: u, wins: (0, 0) };

        while universes.u.len() > 0 {
            universes = advance1(&universes, &mul);
            universes = advance2(&universes, &mul);
        }

        assert_eq!( universes.wins.0, 444356092776315);
        assert_eq!( universes.wins.1, 341960390180808);
    }
}
