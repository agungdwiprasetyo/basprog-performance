
fn main() {
    let mut sum: i64 = 0;

    let n: i64 = 1000000000;
    println!("n: {}", n);
    for i in 0..n {
        sum = sum + i;
    }

    println!("sum = {}", sum);
}