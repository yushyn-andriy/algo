

struct Person {
	name: String,
	age: i32,
	color: String
}

fn print_string(s: &str) {
	println!("{:?}", s);
}


fn main() {
	let people = vec![
		Person{
			name: "John".to_owned(),
			age: 9,
			color: String::from("green"),
		},
		Person{
			name: "Allen".to_owned(),
			age: 91,
			color: String::from("red"),
		},
		Person{
			name: "Lina".to_owned(),
			age: 87,
			color: String::from("green"),
		},
	];
	
	for item in people {
		if item.age<=10 {
			print_string(&item.name);
			print_string(&item.color);
			println!("=============================");
		}
	}
}
