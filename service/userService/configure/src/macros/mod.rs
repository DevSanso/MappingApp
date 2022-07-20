

#[macro_export]
#[cfg(not(target_os="windows"))]
macro_rules! base_path {
    () => {
        concat!(env!("CARGO_MANIFEST_DIR"),"/","resources/")
    };
}

#[macro_export]
#[cfg(target_os="windows")]
macro_rules! base_path {
    () => {
        concat!(env!("CARGO_MANIFEST_DIR"),"\\","resources\\")
    };
}



#[macro_export]
macro_rules! get_env_raw_data {
    ($prefix:expr) => {
       {
            let data = include_str!(
                concat!(
                    crate::base_path!(),
                    env!(concat!($prefix,"_CONFIG_PATH"))
                )
            );
            data
       }
    };
}




