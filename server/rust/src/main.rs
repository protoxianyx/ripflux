use axum::{Router, routing::get};
use std::net::SocketAddr;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    // define routes
    let app = Router::new().route("/", get(root));

    // bind to address
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    let listener = TcpListener::bind(addr).await.unwrap();
    println!("🚀 Server running at http://{}", addr);

    // serve the app (new Axum 0.8 way)
    axum::serve(listener, app).await.unwrap();
}

async fn root() -> &'static str {
    "Ripflux Axum backend (v0.8) is running V4! And this is ruiing mildly nice"
    
}
