use twentyone::Player;
use twentyone::Dice;
fn main() {
    let mut p1 = Player{ pos: 5, score:0};
    let mut p2 = Player{ pos: 9, score:0};
    let mut d = Dice{last:0, count:0};

    loop {
        let v = d.roll3times();
        p1.advance( v);
        if p1.score >= 1000 {
            println!("Player 1 wins");
            println!("value {}", p2.score * d.count );
         
            break
        }
        let v = d.roll3times();
        p2.advance( v);
        if p2.score >= 1000 {
            println!("Player 2 winds");
            println!("value {}", p1.score * d.count );
           
            break
        }
       
    }

    println!("End")
}
