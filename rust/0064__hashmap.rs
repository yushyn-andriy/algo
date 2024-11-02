use std::collections::HashMap;

fn main() {
	let mut people = HashMap::new();

	people.insert("A", 1);
	people.insert("B", 2);
	
	people.remove("B");
	match people.get("B") {
		Some(val) => println!("{:?}", val),
		None => println!("Not found")
	}
	// people.keys()
	// people.values()
	// people.iter()

	for (k, v) in people.iter() {
		println!("{}: {}", k, v);
	}
	

}
