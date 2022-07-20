use std::sync::Once;
use serde_json::from_str;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug,Clone)]
pub struct DatabaseConfigure {
    pub address : String,
    #[serde(rename = "userName")]
    pub username : String,
    pub password : String,
    #[serde(rename = "dbName")]
    pub dbname : String,
    #[serde(rename = "maxConnection")]
    pub max_connection : i32,
    #[serde(rename = "minConnection")]
    pub min_connection : i32,
    #[serde(rename = "connectTimeoutSec")]
    pub connect_timeout_sec : i32,
    #[serde(rename = "connectionIdleTimeoutSec")]
    pub connection_idle_timeout_sec : i32
}

static mut CONFIGURE : Option<DatabaseConfigure> = None;
static mut ONCE_CALL: Once = Once::new();


use crate::get_env_raw_data;

pub fn get_database_configure() -> DatabaseConfigure {
    unsafe {
        ONCE_CALL.call_once(|| {
            let cast = from_str(
                get_env_raw_data!("DATABASE")
            );
            if cast.is_err() {panic!("{}",cast.unwrap_err())}
            CONFIGURE = Some(cast.unwrap());
        });
        return CONFIGURE.clone().unwrap();
    }
}