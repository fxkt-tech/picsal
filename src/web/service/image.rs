use axum::{
    extract::{Query, State},
    http::{HeaderMap, HeaderValue, StatusCode},
    routing::get,
    Json, Router,
};
use image::{imageops::FilterType, ImageFormat};
use serde::{Deserialize, Serialize};
use std::io::Cursor;

use crate::web::server::SharedState;

pub fn register() -> Router<SharedState> {
    Router::new()
        .route("/resize", get(resize))
        .route("/set_name", get(set_name))
        .route("/get_name", get(get_name))
}

fn image_headers() -> HeaderMap {
    let mut headers = HeaderMap::new();
    headers.insert("Content-Type", HeaderValue::from_static("image/jpeg"));

    headers
}

#[derive(Deserialize, Debug)]
struct ResizeRequest {
    jobid: String,
    width: u32,
    height: u32,
}

// 缩放
async fn resize(Query(q): Query<ResizeRequest>) -> Result<(HeaderMap, Vec<u8>), StatusCode> {
    println!("{}", q.jobid);

    let img = image::open("src/i.jpg").unwrap();
    let img = img.resize(q.width, q.height, FilterType::Gaussian);
    let mut bytes = Vec::new();
    img.write_to(&mut Cursor::new(&mut bytes), ImageFormat::Jpeg)
        .unwrap();

    Ok((image_headers(), bytes))
}

#[derive(Deserialize, Debug)]
struct SetNameRequest {
    jobid: String,
    name: String,
}

#[derive(Serialize, Debug)]
struct SetNameResponse {
    jobid: String,
    name: String,
}

async fn set_name(
    Query(q): Query<SetNameRequest>,
    State(state): State<SharedState>,
) -> Json<SetNameResponse> {
    let mut data = state.write().unwrap();
    data.db.insert(q.jobid, q.name);

    let res = SetNameResponse {
        jobid: String::new(),
        name: String::new(),
    };

    Json(res)
}

#[derive(Deserialize, Debug)]
struct GetNameRequest {
    jobid: String,
}

#[derive(Serialize, Debug)]
struct GetNameResponse {
    jobid: String,
    name: String,
}

async fn get_name(
    Query(q): Query<GetNameRequest>,
    State(state): State<SharedState>,
) -> Json<GetNameResponse> {
    let data = state.read().unwrap();
    let default = String::from("");
    let val = data.db.get(&q.jobid).unwrap_or(&default);
    let res = GetNameResponse {
        jobid: q.jobid,
        name: val.clone(),
    };

    Json(res)
}
