use rs_08::*;

fn main() {
    let (i, m) = parse( "input.txt");
    println!("{}", star1( &i, &m));

    println!("{}", star2( &i, &m));
}
