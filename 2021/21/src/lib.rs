pub struct Dice {
    pub last : i32,
    pub count : i32
}

impl Dice { 

  pub fn roll3times( &mut self) -> i32 {
      self.count +=3;
      let ret = 3*self.last + 1+2+3;
        self.last +=3;
        ret
  }
}

pub struct Player {
    pub pos: i32,
    pub score: i32
}

impl Player {
    pub fn advance( &mut self, howmuch : i32 ) {
        self.pos += howmuch;
        self.pos =  ((self.pos-1) % 10) +1;
        self.score += self.pos;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example() {
        let mut p1 = Player{ pos: 4, score:0};
        let mut p2 = Player{ pos: 8, score:0};
        let mut d = Dice{last:0, count:0};
        loop {
            let v = d.roll3times();
            p1.advance( v);
            if p1.score >= 1000 {
                println!("Player 1 wins");
            
                assert_eq!( 739785, p2.score * d.count );
                break
            }
            let v = d.roll3times();
            p2.advance( v);
            if p2.score >= 1000 {
                println!("Player 2 winds");
                println!("value {}", p1.score * d.count );
                assert_eq!( 1, 0);
                break
            }
           
        }

        println!("End")
    }

}