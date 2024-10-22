enum Discount {
	Percent(i32),
	Flat(i32),
}

struct Ticket {
	event: String,
	price: i32,
}

fn main() {
	let n = 3;
	match n {
		3 => println!("three"),
		other => println!("{:?}", other)
	}
	let flat = Discount::Flat(2);
	match flat {
		Discount::Flat(amount) => println!("{:?}", amount),
		Discount::Flat(2) => println!("flat 2"),
		_ => (),
	}	
	let conc = Ticket {
		event: "concert".to_owned(),
		price: 50,
	};
	match conc {
		Ticket{price: 50, event} => println!("event {}", event),
		Ticket{price, ..} => println!("price {}", price),
	}

}

