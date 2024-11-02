
fn maybe_word() ->  Option<String> {
	Some("Ich gehe nach Haus".to_owned())
}


fn main() {
	let word_len: Option<i32> = maybe_word().map(|w| w.len()).map(|l| l*l).map(|x| (x as f64).sqrt() as i32);
	match word_len {
		Some(num) => println!("{:?}", num),
		None => (),
	}
}
