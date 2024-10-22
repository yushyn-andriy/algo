

fn print_it(s: &str) {
	println!("{}", s);
}


fn main() {
	print_it("Hello world");
	let s = "Owned string".to_owned();
	let s1 = String::from("from some string");
	print_it(&s);
	print_it(&s1);
}

