fn main() {
	let a: Vec<String> = vec!["hello".to_owned(), "world".to_owned()];
	let b = "hello".to_owned();
	println!("{:?}", a.iter().position(|x| *x == b ));

}
