use rand::Rng;

fn perform_survey() -> Option<String> {
	let mut rng = rand::thread_rng();
	let xr: u8 = rng.gen();
	if xr >= 128 {
		Some("Some data".to_owned())
	} else {
		None
	}
}

fn main() {
	let mut survey: Vector<String> = vec![];
	let mut i = 0;
	while i<100 {
		let res = perform_survey();
		match res {
			Some(value) => survey.add(value),
			None => (),
		}

		i = i+1;
	}
	println!("len is {:?}", survey.len());
}
