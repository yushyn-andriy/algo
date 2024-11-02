use std::collections::HashMap;



struct StockItem {
	name: String,
	total_number: i32,
}


fn main() {
	let items = vec![("Chair", 223), ("Table", 0), ("Lamp", 100), ("Bed", 30), ("Couch", 450)];
	

	let mut store = HashMap::new();
	for item in &items {
		store.insert(
			item.0.to_owned(),
			StockItem{
				name: item.0.to_owned(),
				total_number: item.1,
			},
			
		);
	}
	let mut total_stock = 0;	
	for (key, val) in store.iter() {
		total_stock = total_stock + val.total_number;
		if val.total_number == 0 {
			println!("{:?} out of stock", val.name);
		} else {
			println!("{:?} has {:?} items", val.name, val.total_number);
		}
	}

	println!("total stock number is {:?}", total_stock);
}
