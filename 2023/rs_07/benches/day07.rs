use criterion::{black_box, criterion_group, criterion_main, Criterion};
use rs_07::*;

pub fn criterion_benchmark(c: &mut Criterion) {
    let data = parse("input.txt");
    c.bench_function("star_1", |b| b.iter(|| eval(black_box(&data), false)));
    c.bench_function("star_2", |b| b.iter(|| eval(black_box(&data), true)));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);