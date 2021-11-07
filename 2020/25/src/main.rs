
fn loopsize( subject: u64, target: u64) -> u64 {
    let mut pivot = 1;
    let mut count = 0;
    while pivot != target {
        pivot = (pivot * subject) % 20201227;
        count+=1;
    }
    count
}

fn do_loop( subject: u64, loopsize: u64) -> u64  {
    let mut pivot = 1;
    for _ in 0..loopsize {
        pivot = (pivot * subject) % 20201227;
    }
    pivot
}

fn main() {

    println!("{}", loopsize(7,5764801));
    println!("{}", loopsize(7,17807724));
//14222596
//4057428
    let loopsize1 = loopsize(7,14222596);
    let loopsize2 = loopsize(7,4057428);
    println!("{}", do_loop( 14222596, loopsize2));
    println!("{}", do_loop( 4057428, loopsize1));

    
}
