use serde::{Deserialize, Debug};

#[derive(Deserialize, Debug)]
struct CatFact {
    fact: String,
    length: i32,
}

#[tokio::main]
async fn main() {
    let fact = get_cat_fact().await;

    println!("fact = {:#?}", fact);
}

async fn get_cat_fact() -> Result<CatFact, Box<dyn std::error::Error>> {
    let client = reqwest::Client::new();
    let body = client.get("https://catfact.ninja/fact").send()
        .await?
        .json::<CatFact>()
        .await?;

    Ok(body)
}