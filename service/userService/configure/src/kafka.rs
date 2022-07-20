use std::sync::Once;
use serde_json::from_str;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug,Clone)]
pub struct KafkaConfigure {
    pub address : String,
    #[serde(rename = "producerTopic")]
    pub producer_topic : String,
    #[serde(rename = "consumerTopic")]
    pub consumer_topic : String,
    #[serde(rename = "consumerGroup")]
    pub consumer_group : String
}

static mut CONFIGURE : Option<KafkaConfigure> = None;
static mut ONCE_CALL: Once = Once::new();


use crate::get_env_raw_data;

pub fn get_kafka_configure() -> KafkaConfigure {
    unsafe {
        ONCE_CALL.call_once(|| {
            let cast = from_str(
                get_env_raw_data!("KAFKA")
            );
            if cast.is_err() {panic!("{}",cast.unwrap_err())}
            CONFIGURE = Some(cast.unwrap());
        });
        return CONFIGURE.clone().unwrap();
    }
}