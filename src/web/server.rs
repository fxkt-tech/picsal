use axum::Router;
use serde_json::de;
use std::{
    collections::HashMap,
    net::SocketAddr,
    sync::{Arc, RwLock},
};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

use crate::web::service;

#[tokio::main]
pub async fn start() {
    tracing_subscriber::registry()
        .with(
            tracing_subscriber::EnvFilter::try_from_default_env()
                .unwrap_or_else(|_| "picsal=debug,tower_http=debug".into()),
        )
        .with(tracing_subscriber::fmt::layer())
        .init();
    let state = Arc::new(RwLock::new(AppState::new()));
    let app = Router::new()
        .nest("/", service::image::register())
        .with_state(state);

    let addr: SocketAddr = "127.0.0.1:4396".parse().unwrap();
    tracing::debug!("listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

pub type SharedState = Arc<RwLock<AppState>>;

pub struct AppState {
    pub db: HashMap<String, String>,
}

impl AppState {
    pub fn new() -> Self {
        Self { db: HashMap::new() }
    }
}
