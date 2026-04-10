fn main() {
    println!("Nexus Chain Rust-Go Interop Interface");
    let data = call_go_core("NEXUS_DATA".to_string());
    println!("Go Response: {}", data);
}

fn call_go_core(input: String) -> String {
    return format!("RUST_INTEROP_{}", input);
}
