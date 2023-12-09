
fn solve2( n: i64, s: i64) -> i64 {
    // (N-x)x > S
    // -x2 + NX - S > 0
    // x*x - n*x + s = 0
    // v n - sqrt(n^2-4s)/2 - ^sqrt(N^2-4S)/2 
    let n_ = n as f64;
    let s_ = s as f64;
    
    let t =  ((n_*n_-4.0*s_) as f64).sqrt();
    let lower_v = (n_  - t) / 2.0;
    let lower = lower_v.ceil() as i64;
    let upper = ((n_  + t) / 2.0).trunc() as i64;
    if (lower_v -lower_v.ceil()).abs() == 0.0 {
         return upper - lower - 1;
    }    
    
    return upper - lower + 1; 
 }   

#[cfg(test)]

#[test]
fn test1( ) {
    assert_eq!(solve2( 7, 9), 4);
    assert_eq!(solve2( 15, 40), 8);
    assert_eq!(solve2( 30, 200), 9);
}

fn main() {
    let mut m = solve2( 44, 277);
    m *= solve2( 89, 1136);
    m *= solve2( 96, 1890);
    m *= solve2( 91, 1768);
    println!("{}", m);

    println!("{}", solve2( 44_89_96_91, 277_1136_1890_1768));
}
