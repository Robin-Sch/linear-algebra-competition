extern crate nalgebra;
extern crate rand;

use nalgebra::{DMatrix, SymmetricEigen};
use std::time::Instant;
use rand::{thread_rng, Rng};

fn main() {
	let mut rng = thread_rng();

	let n = 1000;
	let A = DMatrix::<f64>::from_fn(n, n, |_i, _j| rng.gen::<f64>());

	let start = Instant::now();
	let _eig = SymmetricEigen::new(A);
	let end = Instant::now();
	
	let time = end - start;
	let sec = time.as_secs();
	let ms = time.subsec_millis();
	println!("{}.{:03}", sec, ms);
}