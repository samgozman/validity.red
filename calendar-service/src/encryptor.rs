#![allow(dead_code, unused_variables)]

use aes_gcm::{
    aead::{heapless::Vec, AeadInPlace, KeyInit, OsRng},
    Aes256Gcm, Nonce,
};

// TODO: Retrun unfixed size Vec
pub fn encrypt(data: String, key: String) -> Result<Vec<u8, 128>, Box<dyn std::error::Error>> {
    let key = Aes256Gcm::generate_key(&mut OsRng);
    let cipher = Aes256Gcm::new(&key);
    let nonce = Nonce::from_slice(b"unique nonce"); // 96-bits; unique per message

    // TODO: Split data into chunks of 128 bytes
    // TODO: PKCS5Padding and PKCS5UnPadding methods

    let mut buffer: Vec<u8, 128> = Vec::new(); // Note: buffer needs 16-bytes overhead for auth tag tag
    buffer
        .extend_from_slice(data.as_bytes())
        .expect("buffer too small");

    // Encrypt `buffer` in-place, replacing the plaintext contents with ciphertext
    cipher
        .encrypt_in_place(nonce, b"", &mut buffer)
        .expect("encryption failure!");

    // `buffer` now contains the message ciphertext
    assert_ne!(&buffer, data.as_bytes());

    Ok(buffer)
}

pub fn decrypt(data: &[u8], key: &[u8]) {}
