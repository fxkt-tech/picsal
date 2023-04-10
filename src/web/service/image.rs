use axum::{
    extract::Query,
    http::{HeaderMap, HeaderValue, StatusCode},
    routing::get,
    Router,
};
use image::{imageops::FilterType, ImageFormat};
use serde::Deserialize;
use std::io::Cursor;

pub fn register() -> Router {
    Router::new().route("/resize", get(resize))
}

#[derive(Deserialize, Debug)]
struct ResizeRequest {
    jobid: String,
    width: u32,
    height: u32,
}

// 缩放
async fn resize(Query(payload): Query<ResizeRequest>) -> Result<(HeaderMap, Vec<u8>), StatusCode> {
    let mut headers = HeaderMap::new();
    headers.insert("Content-Type", HeaderValue::from_static("image/jpeg"));

    println!("{}", payload.jobid);

    let img = image::open("src/i.jpg").unwrap();
    let img = img.resize(payload.width, payload.height, FilterType::Gaussian);
    let mut bytes = Vec::new();
    img.write_to(&mut Cursor::new(&mut bytes), ImageFormat::Jpeg)
        .unwrap();

    Ok((headers, bytes))
}
